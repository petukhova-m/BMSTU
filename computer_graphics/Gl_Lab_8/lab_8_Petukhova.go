package main

import (
	"encoding/binary"
	"fmt"
	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/image/bmp"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Color struct {
	r uint8
	g uint8
	b uint8
}

type Point struct {
	tCoords []float32
	n      []float32
	coords []float32
}

type Texture struct {
	id uint32
	data []uint8
}

type Cylinder struct {
	a, b, h        float32
	e_step, h_step int
	points  []float32
	count int32
	color mgl32.Vec3
}


type Projector struct {
	pos mgl32.Vec3
	col mgl32.Vec3
	dir mgl32.Vec3
	exp float64
	angle float64
}

const (
	width  = 640
	height = 640
	Pi     = math.Pi
)

var (
	err error

	//Mods
	r_mod = true
	isText = false

	//Projection
	theta = Pi / 6
	phi   = -Pi / 18
	r     = -1.0
	col1 = mgl32.Vec4{float32(math.Cos(theta)), float32(math.Sin(phi) * math.Sin(theta)), float32(math.Sin(theta) * math.Cos(phi)), float32(-r * math.Sin(theta) * math.Cos(phi))}
	col2 = mgl32.Vec4{0.0, float32(math.Cos(phi)), float32(-math.Sin(phi)), float32(r * math.Sin(phi))}
	col3 = mgl32.Vec4{float32(math.Sin(theta)), float32(-math.Cos(theta) * math.Sin(phi)), float32(-math.Cos(theta) * math.Cos(phi)), float32(r * math.Cos(theta) * math.Cos(phi))}
	col4 = mgl32.Vec4{0.0, 0.0, 1.0, 2.0}
	projection = mgl32.Mat4FromCols(col1, col2, col3, col4)
	view = mgl32.Ident4()
	model = mgl32.Ident4()

	//Cylinder params
	cyl Cylinder
	sc = float32(1.0)
	pitch, roll, yaw = float32(0.0), float32(0.0), float32(0.0)
	cX, cY, cZ = float32(0.0), float32(0.0), float32(0.0)
	lX, lY, lZ = float32(0.0), float32(0.0), float32(0.0)
	t = 0.0

	//Light
	prjctr = Projector{pos: mgl32.Vec3{-1, 1, 1},
		dir: mgl32.Vec3{1, 0, -2},
		col: mgl32.Vec3{1.0, 1.0, 1.0},
		exp: 25.0,
		angle: 90}

	viewPos = mgl32.Vec3{-1.0, 0.0, 2.0}
	lightRotate = mgl32.Ident4()

	//Texture
	texture, spec_map, amb_map Texture
	texSize = 1024

	//Shaders
	VBO, VAO uint32
	vShader, fShader uint32
)

//Other Functions

func normalize(x float32, y float32, z float32) *[]float32 {
	vecLen := float32(math.Sqrt(float64(x*x + y*y + z*z)))
	return &[]float32{float32(x) / vecLen, float32(y) / vecLen, float32(z) / vecLen}
}

func deconsPt(ps *[]Point) []float32 {
	var pars []float32
	for _, p := range *ps {
		pars = append(pars, p.coords...)
		pars = append(pars, p.n...)
		pars = append(pars, p.tCoords...)
	}
	return pars
}

func calcCylinder(rad_a float32, rad_b float32, c_height float32, step1 int, step2 int) *Cylinder {
	var points []float32
	e_step, h_step := Pi/float64(step1), c_height/float32(step2)
	tx00 := []float32{0.0, 1.0}
	tx10 := []float32{1.0/9.0, 1.0}
	tx11 := []float32{1.0/9.0, 0.7}
	tx01 := []float32{0.0, 0.7}
	for angle := 0.0; angle < 2*Pi; angle += e_step {
		x1 := rad_a * float32(math.Cos(angle))
		z1 := rad_b * float32(math.Sin(angle))
		x2 := rad_a * float32(math.Cos(angle+e_step))
		z2 := rad_b * float32(math.Sin(angle+e_step))

		t1t := []float32{0, c_height / 2, 0}
		t2t := []float32{x2 / 2, c_height / 2, z2 / 2}
		t3t := []float32{x1 / 2, c_height / 2, z1 / 2}
		t4t := []float32{x2, c_height / 2, z2}
		t5t := []float32{x1, c_height / 2, z1}
		t1b := []float32{0, -c_height / 2, 0}
		t2b := []float32{x2 / 2, -c_height / 2, z2 / 2}
		t3b := []float32{x1 / 2, -c_height / 2, z1 / 2}
		t4b := []float32{x2, -c_height / 2, z2}
		t5b := []float32{x1, -c_height / 2, z1}

		p1t := []Point{{coords: t2t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx01},
			{coords: t3t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx11},
			{coords: t1t, n: []float32{0.0, 1.0, 0.0}, tCoords: []float32{0.0, 0.4}}}
		p2t := []Point{{coords: t3t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx11},
			{coords: t2t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx01},
			{coords: t4t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx00}}
		p3t := []Point{{coords: t3t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx11},
			{coords: t4t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx00},
			{coords: t5t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx10}}
		p1b := []Point{{coords: t3b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx11},
			{coords: t2b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx01},
			{coords: t1b, n: []float32{0.0, -1.0, 0.0}, tCoords: []float32{0.0, 0.4}}}
		p2b := []Point{{coords: t2b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx01},
			{coords: t3b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx11},
			{coords: t4b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx00}}
		p3b := []Point{{coords: t3b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx11},
			{coords: t5b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx10},
			{coords: t4b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx00}}
		points = append(points, deconsPt(&p1t)...)
		points = append(points, deconsPt(&p2t)...)
		points = append(points, deconsPt(&p3t)...)
		points = append(points, deconsPt(&p1b)...)
		points = append(points, deconsPt(&p2b)...)
		points = append(points, deconsPt(&p3b)...)

		for i := 0; i < step2; i++ {
			y := -c_height/2 + h_step*float32(i)
			t00 := []float32{float32((angle+e_step)/e_step)/float32(step1)*2, float32(i)/float32(step2)}
			t10 := []float32{float32((angle)/e_step)/float32(step1)*2, float32(i)/float32(step2)}
			t11 := []float32{float32((angle)/e_step)/float32(step1)*2, float32(i+1)/float32(step2)}
			t01 := []float32{float32((angle+e_step)/e_step)/float32(step1)*2, float32(i+1)/float32(step2)}

			p1Side := []Point{{coords: []float32{x2, y, z2}, n: *normalize(x2, 0, z2), tCoords: t00},
				{coords: []float32{x1, y, z1}, n: *normalize(x1, 0, z1), tCoords: t10},
				{coords: []float32{x1, y + h_step, z1}, n: *normalize(x1, 0, z1), tCoords: t11}}
			p2Side := []Point{{coords: []float32{x1, y + h_step, z1}, n: *normalize(x1, 0, z1), tCoords: t11},
				{coords: []float32{x2, y + h_step, z2}, n: *normalize(x2, 0, z2), tCoords: t01},
				{coords: []float32{x2, y, z2}, n: *normalize(x2, 0, z2), tCoords: t00}}
			points = append(points, deconsPt(&p1Side)...)
			points = append(points, deconsPt(&p2Side)...)
		}
	}
	count := int32(len(points)/8)
	fmt.Println(4*step1*(1 + step2))
	return &Cylinder{a: rad_a, b: rad_b, h: c_height, e_step: step1, h_step: step2, points: points, count: count, color: mgl32.Vec3{0.97, 0.91, 0.64}}
}

func readConfig() {
	configFile, err := ioutil.ReadFile("config")
	if err != nil {
		panic(err)
	}

	configLines := strings.Split(string(configFile), "\n")
	for i := 0; i < len(configLines); i++ {
		if configLines[i] != "" {
			configLine := strings.Split(string(configLines[i]), " ")
			switch configLine[0] {
			case "cX":
				tcX, _ := strconv.ParseFloat(configLine[1], 64)
				cX = float32(tcX)
			case "cY":
				tcY, _ := strconv.ParseFloat(configLine[1], 64)
				cY = float32(tcY)
			case "cZ":
				tcZ, _ := strconv.ParseFloat(configLine[1], 64)
				cZ = float32(tcZ)
			case "lX":
				tlX, _ := strconv.ParseFloat(configLine[1], 64)
				lX = float32(tlX)
			case "lY":
				tlY, _ := strconv.ParseFloat(configLine[1], 64)
				lY = float32(tlY)
			case "lZ":
				tlZ, _ := strconv.ParseFloat(configLine[1], 64)
				lZ = float32(tlZ)
			case "exp":
				prjctr.exp, _ = strconv.ParseFloat(configLine[1], 64)
			case "pitch":
				tpitch, _ := strconv.ParseFloat(configLine[1], 64)
				pitch = float32(tpitch)
			case "roll":
				troll, _ := strconv.ParseFloat(configLine[1], 64)
				roll = float32(troll)
			case "yaw":
				tyaw, _ := strconv.ParseFloat(configLine[1], 64)
				yaw = float32(tyaw)
			case "sc":
				tsc, _ := strconv.ParseFloat(configLine[1], 64)
				sc = float32(tsc)
			case "t":
				t, _ = strconv.ParseFloat(configLine[1], 64)
			case "isText":
				isText, _ = strconv.ParseBool(configLine[1])
			case "cyl":
				a, _ := strconv.ParseFloat(configLine[1], 64)
				b, _ := strconv.ParseFloat(configLine[2], 64)
				h, _ := strconv.ParseFloat(configLine[3], 64)
				eStep, _ := strconv.ParseInt(configLine[4],10,  64)
				hStep, _ := strconv.ParseInt(configLine[5], 10, 64)
				cyl = *calcCylinder(float32(a), float32(b), float32(h), int(eStep), int(hStep))
			}
		}
	}
	initArray()
	changeModel()
	rotateLight()
	log.Println("Config loaded")
}

func makeConfig()  {
	configFile, err := os.Create("config")
	if err != nil {
		panic(err)
	}
	configInfo := "cX " + strconv.FormatFloat(float64(cX), 'E', -1, 64) + "\n"
	configInfo += "cY " + strconv.FormatFloat(float64(cY), 'E', -1, 64) + "\n"
	configInfo += "cZ " + strconv.FormatFloat(float64(cZ), 'E', -1, 64) + "\n"
	configInfo += "lX " + strconv.FormatFloat(float64(lX), 'E', -1, 64) + "\n"
	configInfo += "lY " + strconv.FormatFloat(float64(lY), 'E', -1, 64) + "\n"
	configInfo += "lZ " + strconv.FormatFloat(float64(lZ), 'E', -1, 64) + "\n"
	configInfo += "exp " + strconv.FormatFloat(prjctr.exp, 'E', -1, 64) + "\n"
	configInfo += "pitch " + strconv.FormatFloat(float64(pitch), 'E', -1, 64) + "\n"
	configInfo += "roll " + strconv.FormatFloat(float64(roll), 'E', -1, 64) + "\n"
	configInfo += "yaw " + strconv.FormatFloat(float64(yaw), 'E', -1, 64) + "\n"
	configInfo += "sc " + strconv.FormatFloat(float64(sc), 'E', -1, 64) + "\n"
	configInfo += "t " + strconv.FormatFloat(t, 'E', -1, 64) + "\n"
	configInfo += "isText " + strconv.FormatBool(isText) + "\n"
	configInfo += "cyl " + strconv.FormatFloat(float64(cyl.a), 'E', -1, 64) + " "
	configInfo += strconv.FormatFloat(float64(cyl.b), 'E', -1, 64) + " "
	configInfo += strconv.FormatFloat(float64(cyl.h), 'E', -1, 64) + " "
	configInfo += strconv.FormatInt(int64(cyl.e_step), 10) + " "
	configInfo += strconv.FormatInt(int64(cyl.h_step), 10) + "\n"
	_, err = configFile.WriteAt([]byte(configInfo), 0)
	log.Println("Config saved")
	err = configFile.Close()
	if err != nil {
		panic(err)
	}
}

func getPixel(r, g, b, a uint32) []uint8 {
	return []uint8{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func RGBA(path string) *[]uint8 {
	textureFile, _ := os.Open(path)
	defer textureFile.Close()

	loadedTexture, _ := bmp.Decode(textureFile)

	var data []uint8
	for y := 0; y < texSize; y++ {
		for x := 0; x < texSize; x++ {
				pixel := getPixel(loadedTexture.At(x,y).RGBA())
			for i := 0; i < 4; i++ {
				data = append(data, pixel[i])
			}
		}
	}
	return &data
}

func changeModel()  {
	model = mgl32.Ident4()
	model = mgl32.Translate3D(cX, cY, cZ)
	model = model.Mul4(mgl32.HomogRotate3D(-yaw, mgl32.Vec3{1, 0, 0}))
	model = model.Mul4(mgl32.HomogRotate3D(-pitch, mgl32.Vec3{0, 0, 1}))
	model = model.Mul4(mgl32.HomogRotate3D(-roll, mgl32.Vec3{0, 1, 0}))
	model = model.Mul4(mgl32.Scale3D(sc, sc, sc))
}

func rotateLight()  {
	lightRotate = mgl32.Ident4()
	lightRotate = lightRotate.Mul4(mgl32.HomogRotate3D(-lX, mgl32.Vec3{1, 0, 0}))
	lightRotate = lightRotate.Mul4(mgl32.HomogRotate3D(-lZ, mgl32.Vec3{0, 0, 1}))
	lightRotate = lightRotate.Mul4(mgl32.HomogRotate3D(-lY, mgl32.Vec3{0, 1, 0}))
}

//Initializations

func initWindow() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	if err := gl.Init(); err != nil {
		panic(err)
	}
	window, err := glfw.CreateWindow(width, height, "Cylinder", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval( 1);

	window.SetKeyCallback(keyCallback)
	window.SetFramebufferSizeCallback(Resize)

	return window
}

func initOpenGL() uint32 {
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println(version)

	vShader = gl.CreateShader(gl.VERTEX_SHADER)
	initShader(vShader, "vertexShader.glsl")
	defer gl.DeleteShader(vShader)

	fShader = gl.CreateShader(gl.FRAGMENT_SHADER)
	initShader(fShader, "fragmentShader.glsl")
	defer gl.DeleteShader(fShader)

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vShader)
	gl.AttachShader(prog, fShader)
	gl.LinkProgram(prog)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(0.3, 0.3, 0.3, 0.0)

	return prog
}

func initShader(sh uint32, source string) {
	var success int32

	shString, err := ioutil.ReadFile(source)
	if err != nil {
		panic(err)
	}
	ss, freeFnv := gl.Strs(string(shString) + "\x00")
	gl.ShaderSource(sh, 1, ss, nil)
	gl.CompileShader(sh)
	gl.GetShaderiv(sh, gl.COMPILE_STATUS, &success)
	if success != gl.TRUE {
		var logLen int32
		gl.GetShaderiv(sh, gl.INFO_LOG_LENGTH, &logLen)
		inflog := strings.Repeat("\x00", int(logLen+1))
		gl.GetShaderInfoLog(sh, logLen, nil, gl.Str(inflog))
		log.Fatal("shader error \n" + source + ":\n" + inflog)
	}
	freeFnv()
}

func initTexture(tex *uint32, texConst uint32, texFile string, data *[]uint8)  {
	gl.GenTextures(1, tex)
	gl.ActiveTexture(texConst)
	gl.BindTexture(gl.TEXTURE_2D, *tex)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	*data = *RGBA(texFile)
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(texSize), int32(texSize), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr((*data)))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func initArray() {
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	defer gl.BindVertexArray(0)

	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(cyl.points), gl.Ptr(cyl.points), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)
}

//CallBacks

func Resize(win *glfw.Window, w int, h int) {
	var size, dx, dy int32
	if h > w {
		size = int32(h)
		dx, dy = int32((w-h)/2), 0
	} else {
		size = int32(w)
		dx, dy = 0, int32((h-w)/2)
	}
	gl.Viewport(dx, dy, size, size)
}

func keyCallback(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		win.SetShouldClose(true)
	}
	//Rotation
	if key == glfw.KeyLeftShift && action == glfw.Press {
		r_mod = !r_mod
	}
	if key == glfw.KeyE && action == glfw.Repeat {
		roll += 0.01
		changeModel()
	}
	if key == glfw.KeyQ && action == glfw.Repeat {
		roll -= 0.01
		changeModel()
	}
	if key == glfw.KeyD && action == glfw.Repeat {
		if r_mod {
			lY -= 0.05
			rotateLight()
		} else {
			pitch += 0.01
			changeModel()
		}
	}
	if key == glfw.KeyA && action == glfw.Repeat {
		if r_mod {
			lY += 0.05
			rotateLight()
		} else {
			pitch -= 0.01
			changeModel()
		}
	}
	if key == glfw.KeyW && action == glfw.Repeat {
		if r_mod {
			lX += 0.05
			lY -= 0.05
			rotateLight()
		} else {
			yaw += 0.01
			changeModel()
		}
	}
	if key == glfw.KeyS && action == glfw.Repeat {
		if r_mod {
			lX -= 0.05
			lY += 0.05
			rotateLight()
		} else {
			yaw -= 0.01
			changeModel()
		}
	}
	//Movement
	if key == glfw.KeyI && action == glfw.Repeat {
		cZ -= 0.05 * float32(math.Cos(theta))
		cX += 0.05 * float32(math.Sin(theta))
		changeModel()
	}
	if key == glfw.KeyK && action == glfw.Repeat {
		cZ += 0.05 * float32(math.Cos(theta))
		cX -= 0.05 * float32(math.Sin(theta))
		changeModel()
	}
	if key == glfw.KeyJ && action == glfw.Repeat {
		cX -= 0.05 * float32(math.Cos(theta))
		cZ -= 0.05 * float32(math.Sin(theta))
		changeModel()
	}
	if key == glfw.KeyL && action == glfw.Repeat {
		cX += 0.05 * float32(math.Cos(theta))
		cZ += 0.05 * float32(math.Sin(theta))
		changeModel()
	}
	if key == glfw.KeyY && action == glfw.Repeat {
		cY += 0.05 * float32(math.Cos(phi))
		changeModel()
	}
	if key == glfw.KeyH && action == glfw.Repeat {
		cY -= 0.05 * float32(math.Cos(phi))
		changeModel()
	}
	//Scale
	if key == glfw.Key0 && action == glfw.Repeat {
		sc += 0.01
		changeModel()
	}
	if key == glfw.Key9 && action == glfw.Repeat {
		sc -= 0.01
		changeModel()
	}
	//Quality
	if key == glfw.KeyZ && action == glfw.Press {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step + 1, cyl.h_step)
		initArray()
	}
	if key == glfw.KeyX && action == glfw.Press && cyl.e_step > 2 {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step - 1, cyl.h_step)
		initArray()
	}
	if key == glfw.KeyC && action == glfw.Press {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step, cyl.h_step + 1)
		initArray()
	}
	if key == glfw.KeyV && action == glfw.Press && cyl.h_step > 1 {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step, cyl.h_step - 1)
		initArray()
	}
	//Light
	if key == glfw.KeyN && action == glfw.Press {
		prjctr.exp += 3.0
	}
	if key == glfw.KeyM && action == glfw.Press && prjctr.exp > 10 {
		prjctr.exp -= 3.0
	}
	if key == glfw.KeyO && action == glfw.Press && prjctr.angle < 90.0 {
		prjctr.angle += 5.0
	}
	if key == glfw.KeyP && action == glfw.Press && prjctr.angle > 5.0 {
		prjctr.angle -= 5.0
	}
	//Texture
	if key == glfw.Key8 && action == glfw.Press {
		if !isText {
			gl.Enable(gl.TEXTURE_2D)
		} else {
			gl.Disable(gl.TEXTURE_2D)
		}
		isText = !isText
	}
	//Reset
	if key == glfw.KeyR && action == glfw.Press{
		cyl = *calcCylinder(0.6, 0.8, 0.8, 15, 8)
		sc = 1.0
		pitch, roll, yaw = 0, 0, 0
		cX, cY, cZ = 0, 0, 0
		lX, lY, lZ = 0, 0, 0
		prjctr.exp = 25
		prjctr.angle = 90
		isText = false
		t = 0
		initArray()
		changeModel()
		rotateLight()
	}
	if key == glfw.Key1 && action == glfw.Press{
		readConfig()
	}
	if key == glfw.Key2 && action == glfw.Press{
		makeConfig()
	}
}

//Main

func main() {
	runtime.LockOSThread()
	window := initWindow()
	defer glfw.Terminate()
	cyl = *calcCylinder(0.6, 0.8, 0.8, 15, 8)
	program := initOpenGL()
	initTexture(&texture.id, gl.TEXTURE0, "barrel.bmp", &texture.data)
	initTexture(&spec_map.id, gl.TEXTURE1, "spec_map.bmp", &spec_map.data)
	initTexture(&amb_map.id, gl.TEXTURE2, "amb_map.bmp", &amb_map.data)
	initArray()

	readConfig()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)

		gl.UniformMatrix4fv(gl.GetUniformLocation(program, gl.Str("projection\x00")), 1, false, &projection[0])
		gl.UniformMatrix4fv(gl.GetUniformLocation(program, gl.Str("view\x00")), 1, false, &view[0])
		gl.UniformMatrix4fv(gl.GetUniformLocation(program, gl.Str("model\x00")), 1, false, &model[0])

		gl.Uniform3fv(gl.GetUniformLocation(program, gl.Str("cylCol\x00")), 1, &cyl.color[0])

		gl.UniformMatrix4fv(gl.GetUniformLocation(program, gl.Str("projRot\x00")), 1, false, &lightRotate[0])
		gl.Uniform3fv(gl.GetUniformLocation(program, gl.Str("lightCol\x00")), 1, &prjctr.col[0])
		gl.Uniform3fv(gl.GetUniformLocation(program, gl.Str("lightPos\x00")), 1, &prjctr.pos[0])
		gl.Uniform3fv(gl.GetUniformLocation(program, gl.Str("projDir\x00")), 1, &prjctr.dir[0])
		gl.Uniform1f(gl.GetUniformLocation(program, gl.Str("exp\x00")), float32(prjctr.exp))
		gl.Uniform1f(gl.GetUniformLocation(program, gl.Str("angle\x00")), float32(prjctr.angle))

		vPosUniform := gl.GetUniformLocation(program, gl.Str("viewPos\x00"))
		gl.Uniform3fv(vPosUniform, 1, &viewPos[0])

		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("texture1\x00")), 0)
		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("spec_map\x00")), 1)
		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("amb_map\x00")), 2)
		isTextUniform := gl.GetUniformLocation(program, gl.Str("isText\x00"))
		if isText {
			gl.Uniform1i(isTextUniform, 0)
		} else {
			gl.Uniform1i(isTextUniform, 1)
		}

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture.id)
		gl.ActiveTexture(gl.TEXTURE1)
		gl.BindTexture(gl.TEXTURE_2D, spec_map.id)
		gl.ActiveTexture(gl.TEXTURE2)
		gl.BindTexture(gl.TEXTURE_2D, amb_map.id)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, cyl.count)
		gl.BindVertexArray(0)
		gl.UseProgram(0)
		glfw.PollEvents()
		window.SwapBuffers()
	}
}