package main
/*
import (
	"fmt"
	"github.com/go-gl/gl/v4.3-compatibility/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"golang.org/x/image/bmp"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Color struct {
	r uint8
	g uint8
	b uint8
}

type Point struct {
	coords []float64
	n      []float32
	tCoords []float32
	aCoords []float64
}

type Poly struct {
	col   Color
	verts []Point
}

type Cylinder struct {
	a, b, h        float64
	e_step, h_step int
	tris  []Poly
}

type Projector struct {
	pos      []float32
	spot_dir []float32
	das      []float32
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
	p_mod = uint32(gl.FILL)
	anim = false
	isText = true
	flag = false

	//Projection
	theta = Pi / 6
	phi   = -Pi / 18
	r     = -1.0
	projection3 = [16]float64{
		math.Cos(theta), math.Sin(phi) * math.Sin(theta), math.Sin(theta) * math.Cos(phi), -r * math.Sin(theta) * math.Cos(phi),
		0.0, math.Cos(phi), -math.Sin(phi), r * math.Sin(phi),
		math.Sin(theta), -math.Cos(theta) * math.Sin(phi), -math.Cos(theta) * math.Cos(phi), r * math.Cos(theta) * math.Cos(phi),
		0.0, 0.0, 1, 2,
	}

	//Cylinder params
	cyl Cylinder
	sc = 1.0
	pitch, roll, yaw = 0.0, 0.0, 0.0
	oX = []float64{1, 0, 0}
	oY = []float64{0, 1, 0}
	oZ = []float64{0, 0, 1}
	cX, cY, cZ = 0.0, 0.0, 0.0
	lX, lY, lZ = 0.0, 0.0, 0.0
	t = 0.0

	//Light
	prjctr = Projector{pos: []float32{0.0, 0.0, 0.0, 1.0},
		spot_dir: *normalize(1, -1, -1),
		das:      []float32{0.97, 0.91, 0.64},
		exp: 25.0,
		angle: 30.0}
	matAmb = []float32{0.1, 0.1, 0.1, 1.0}
	matDif = []float32{1.0, 0.8, 1.0, 1.0}

	//Texture
	texture uint32
	texData []uint8
	texSize = 512

	//Cube params
	GREEN  = Color{0.0, 255, 0.0}
	RED    = Color{255, 0.0, 0.0}
	BLUE   = Color{0.0, 0.0, 255}
	PURPLE = Color{255, 0.0, 255}
	WHITE  = Color{255, 255, 255}
	YELLOW = Color{255, 255, 255}
	colors = [6]Color{
		GREEN, RED, BLUE, PURPLE, WHITE, YELLOW,
	}

	cube = *calcCube()

	verticies = [8][]float64{
		{-0.5, 0.5, -0.5},
		{0.5, 0.5, -0.5},
		{0.5, -0.5, -0.5},
		{-0.5, -0.5, -0.5},
		{-0.5, 0.5, 0.5},
		{0.5, 0.5, 0.5},
		{0.5, -0.5, 0.5},
		{-0.5, -0.5, 0.5},
	}
	inds_f = [24]int{
		0, 1, 2, 3,
		0, 3, 7, 4,
		4, 7, 6, 5,
		5, 6, 2, 1,
		0, 1, 5, 4,
		3, 2, 6, 7,
	}

	norms = [6]*[]float32{
		normalize(0, 0, -1),
		normalize(-1, 0, 0),
		normalize(0, 0, 1),
		normalize(1, 0, 0),
		normalize(0, 1, 0),
		normalize(0, -1, 0),
	}
)

//Other Functions

func normalize(x float64, y float64, z float64) *[]float32 {
	vecLen := float32(math.Sqrt(x*x + y*y + z*z))
	return &[]float32{float32(x) / vecLen, float32(y) / vecLen, float32(z) / vecLen}
}

func randColor() Color {
	return Color{r: 200 + uint8(rand.Intn(50)), g: uint8(rand.Intn(100)), b: uint8(rand.Intn(40))}
}

func calcCylinder(rad_a float64, rad_b float64, c_height float64, step1 int, step2 int) *Cylinder {
	var tris []Poly
	e_step, h_step := Pi/float64(step1), c_height/float64(step2)
	for angle := 0.0; angle < 2*Pi; angle += e_step {
		x1 := rad_a * math.Cos(angle)
		z1 := rad_b * math.Sin(angle)
		x2 := rad_a * math.Cos(angle+e_step)
		z2 := rad_b * math.Sin(angle+e_step)

		tx00 := []float32{float32(x2/2), float32(z2/2)}
		tx10 := []float32{float32(x1/2), float32(z1/2)}
		tx11 := []float32{float32(x1/4), float32(z1/4)}
		tx01 := []float32{float32(x2/4), float32(z2/4)}

		t1t := []float64{0, c_height / 2, 0}
		t2t := []float64{x2 / 2, c_height / 2, z2 / 2}
		t3t := []float64{x1 / 2, c_height / 2, z1 / 2}
		t4t := []float64{x2, c_height / 2, z2}
		t5t := []float64{x1, c_height / 2, z1}
		t1b := []float64{0, -c_height / 2, 0}
		t2b := []float64{x2 / 2, -c_height / 2, z2 / 2}
		t3b := []float64{x1 / 2, -c_height / 2, z1 / 2}
		t4b := []float64{x2, -c_height / 2, z2}
		t5b := []float64{x1, -c_height / 2, z1}

		t1Anim := []float64{x1, 0.0, z1}
		t2Anim := []float64{x2, 0.0, z2}
		t3Anim := []float64{x1/2, c_height/4, z1/2}
		t4Anim := []float64{x2/2, c_height/4, z2/2}
		t5Anim := []float64{x1/2, -c_height/4, z1/2}
		t6Anim := []float64{x2/2, -c_height/4, z2/2}

		p1t := Poly{verts: []Point{{coords: t2t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx01, aCoords: t4Anim},
			{coords: t3t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx11, aCoords: t3Anim},
			{coords: t1t, n: []float32{0.0, 1.0, 0.0}, tCoords: []float32{0.0, 0.0}, aCoords: t1t}},
			col: randColor()}
		p2t := Poly{verts: []Point{{coords: t3t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx11, aCoords: t3Anim},
			{coords: t2t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx01, aCoords: t4Anim},
			{coords: t4t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx00, aCoords: t2Anim}},
			col: randColor()}
		p3t := Poly{verts: []Point{{coords: t3t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx11, aCoords: t3Anim},
			{coords: t4t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx00, aCoords: t2Anim},
			{coords: t5t, n: []float32{0.0, 1.0, 0.0}, tCoords: tx10, aCoords: t1Anim}},
			col: randColor()}
		p1b := Poly{verts: []Point{{coords: t3b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx11, aCoords: t5Anim},
			{coords: t2b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx01, aCoords: t6Anim},
			{coords: t1b, n: []float32{0.0, -1.0, 0.0}, tCoords: []float32{0.0, 0.0}, aCoords: t1b}},
			col: randColor()}
		p2b := Poly{verts: []Point{{coords: t2b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx01, aCoords: t6Anim},
			{coords: t3b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx11, aCoords: t5Anim},
			{coords: t4b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx00, aCoords: t2Anim}},
			col: randColor()}
		p3b := Poly{verts: []Point{{coords: t3b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx11, aCoords: t5Anim},
			{coords: t5b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx10, aCoords: t1Anim},
			{coords: t4b, n: []float32{0.0, -1.0, 0.0}, tCoords: tx00, aCoords: t2Anim}},
			col: randColor()}
		tris = append(tris, []Poly{p1t, p2t, p3t, p1b, p2b, p3b}...)
		for i := 0; i < step2; i++ {
			y := -c_height/2 + h_step*float64(i)
			t00 := []float32{float32((angle+e_step)/e_step)/float32(step1), float32(i)/float32(step2)/2}
			t10 := []float32{float32((angle)/e_step)/float32(step1), float32(i)/float32(step2)/2}
			t11 := []float32{float32((angle)/e_step)/float32(step1), float32(i+1)/float32(step2)/2}
			t01 := []float32{float32((angle+e_step)/e_step)/float32(step1), float32(i+1)/float32(step2)/2}

			p1Side := Poly{verts: []Point{{coords: []float64{x2, y, z2}, n: *normalize(x2, 0, z2), tCoords: t00, aCoords: t2Anim},
				{coords: []float64{x1, y, z1}, n: *normalize(x1, 0, z1), tCoords: t10, aCoords: t1Anim},
				{coords: []float64{x1, y + h_step, z1}, n: *normalize(x1, 0, z1), tCoords: t11, aCoords: t1Anim}},
				col: randColor(),}
			p2Side := Poly{verts: []Point{{coords: []float64{x1, y + h_step, z1}, n: *normalize(x1, 0, z1), tCoords: t11, aCoords: t1Anim},
				{coords: []float64{x2, y + h_step, z2}, n: *normalize(x2, 0, z2), tCoords: t01, aCoords: t2Anim},
				{coords: []float64{x2, y, z2}, n: *normalize(x2, 0, z2), tCoords: t00, aCoords: t2Anim}},
				col: randColor()}
			tris = append(tris, []Poly{p1Side, p2Side}...)
		}
	}
	fmt.Println(4*step1*(1 + step2))
	return &Cylinder{a: rad_a, b: rad_b, h: c_height, e_step: step1, h_step: step2, tris: tris}
}

func calcCube() *[]Poly {
	var cube []Poly
	for i := 0; i < 6; i++ {
		pl := Poly{verts: make([]Point, 0)}
		for j := 4 * i; j < 4*i+4; j++ {
			p := Point{coords: verticies[inds_f[j]], n: *norms[i]}
			pl.verts = append(pl.verts, p)
		}
		pl.col = colors[i]
		cube = append(cube, pl)
	}
	return &cube
}

func getPixel(r, g, b, a uint32) []uint8 {
	return []uint8{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func RGBA() *[]uint8 {
	textureFile, _ := os.Open("plaster.bmp")
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
				cX, _ = strconv.ParseFloat(configLine[1], 64)
			case "cY":
				cY, _ = strconv.ParseFloat(configLine[1], 64)
			case "cZ":
				cZ, _ = strconv.ParseFloat(configLine[1], 64)
			case "lX":
				lX, _ = strconv.ParseFloat(configLine[1], 64)
			case "lY":
				lX, _ = strconv.ParseFloat(configLine[1], 64)
			case "lZ":
				lX, _ = strconv.ParseFloat(configLine[1], 64)
			case "exp":
				prjctr.exp, _ = strconv.ParseFloat(configLine[1], 64)
			case "angle":
				prjctr.angle, _ = strconv.ParseFloat(configLine[1], 64)
			case "pitch":
				pitch, _ = strconv.ParseFloat(configLine[1], 64)
			case "roll":
				roll, _ = strconv.ParseFloat(configLine[1], 64)
			case "yaw":
				yaw, _ = strconv.ParseFloat(configLine[1], 64)
			case "sc":
				sc, _ = strconv.ParseFloat(configLine[1], 64)
			case "t":
				t, _ = strconv.ParseFloat(configLine[1], 64)
			case "anim":
				anim, _ = strconv.ParseBool(configLine[1])
			case "isText":
				isText, _ = strconv.ParseBool(configLine[1])
			case "cyl":
				a, _ := strconv.ParseFloat(configLine[1], 64)
				b, _ := strconv.ParseFloat(configLine[2], 64)
				h, _ := strconv.ParseFloat(configLine[3], 64)
				eStep, _ := strconv.ParseInt(configLine[4],10,  64)
				hStep, _ := strconv.ParseInt(configLine[5], 10, 64)
				cyl = *calcCylinder(a, b, h, int(eStep), int(hStep))
			}
		}
	}
	log.Println("Config loaded")
}

func makeConfig()  {
	configFile, err := os.Create("config")
	if err != nil {
		panic(err)
	}
	configInfo := "cX " + strconv.FormatFloat(cX, 'E', -1, 64) + "\n"
	configInfo += "cY " + strconv.FormatFloat(cY, 'E', -1, 64) + "\n"
	configInfo +=  "cZ " + strconv.FormatFloat(cZ, 'E', -1, 64) + "\n"
	configInfo += "lX " + strconv.FormatFloat(lX, 'E', -1, 64) + "\n"
	configInfo += "lY " + strconv.FormatFloat(lY, 'E', -1, 64) + "\n"
	configInfo += "lZ " + strconv.FormatFloat(lZ, 'E', -1, 64) + "\n"
	configInfo += "exp " + strconv.FormatFloat(prjctr.exp, 'E', -1, 64) + "\n"
	configInfo += "angle " + strconv.FormatFloat(prjctr.angle, 'E', -1, 64) + "\n"
	configInfo += "pitch " + strconv.FormatFloat(pitch, 'E', -1, 64) + "\n"
	configInfo += "roll " + strconv.FormatFloat(roll, 'E', -1, 64) + "\n"
	configInfo += "yaw " + strconv.FormatFloat(yaw, 'E', -1, 64) + "\n"
	configInfo += "sc " + strconv.FormatFloat(sc, 'E', -1, 64) + "\n"
	configInfo += "t " + strconv.FormatFloat(t, 'E', -1, 64) + "\n"
	configInfo += "anim " + strconv.FormatBool(anim) + "\n"
	configInfo += "isText " + strconv.FormatBool(isText) + "\n"
	configInfo += "cyl " + strconv.FormatFloat(cyl.a, 'E', -1, 64) + " "
	configInfo += strconv.FormatFloat(cyl.b, 'E', -1, 64) + " "
	configInfo += strconv.FormatFloat(cyl.h, 'E', -1, 64) + " "
	configInfo += strconv.FormatInt(int64(cyl.e_step), 10) + " "
	configInfo += strconv.FormatInt(int64(cyl.h_step), 10) + "\n"
	_, err = configFile.WriteAt([]byte(configInfo), 0)
	log.Println("Config saved")
	err = configFile.Close()
	if err != nil {
		panic(err)
	}
}

//Initializations

func initWindow() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(width, height, "Cylinder", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval( 0);

	window.SetKeyCallback(key_callback)

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println(version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.ALPHA_TEST)
	gl.AlphaFunc(gl.NOTEQUAL, 0.0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(0.3, 0.3, 0.3, 0.0)
	gl.Enable(gl.LIGHTING)
	gl.LightModelf(gl.LIGHT_MODEL_TWO_SIDE, gl.TRUE)
	gl.Enable(gl.NORMALIZE)

	return prog
}

func initTexture()  {
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexEnvf(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	texData = *RGBA()
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(texSize), int32(texSize), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&texData[0]))
	//gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

//Draw_Function

func drawCube() {
	gl.Translated(-0.7, -0.4, 0.4)
	gl.Scaled(0.25, 0.25, 0.25)
	gl.Begin(gl.QUADS)
	for _, pl := range cube {
		gl.Color3ub(pl.col.r, pl.col.g, pl.col.b)
		for _, p := range pl.verts {
			gl.Normal3fv(&p.n[0])
			gl.Vertex3dv(&p.coords[0])
		}
	}
	gl.End()
}

func drawCylinder()  {
	gl.Scaled(sc, sc, sc)
	gl.Translated(cX, cY, cZ)
	gl.Rotated(-yaw, 1.0, 0.0, 0.0)
	gl.Rotated(-pitch, 0.0, 0.0, 1.0)
	gl.Rotated(-roll, 0.0, 1.0, 0.0)

	gl.CallList(1)
}

func setLight() {
	gl.Translated(-2.0, 2.0, 2.0)
	gl.Rotated(lX, 1.0, 0.0, 0.0)
	gl.Rotated(lY, 0.0, 1.0, 0.0)
	gl.Materialfv(gl.FRONT_AND_BACK, gl.AMBIENT, &matAmb[0])
	gl.Materialfv(gl.FRONT_AND_BACK, gl.DIFFUSE, &matDif[0])
	gl.Materialfv(gl.FRONT_AND_BACK, gl.SPECULAR, &matAmb[0])
	gl.Materialf(gl.FRONT_AND_BACK, gl.SHININESS, 64)
	gl.Enable(gl.LIGHT0)
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &prjctr.pos[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &prjctr.das[0])
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &prjctr.das[0])
	gl.Lightfv(gl.LIGHT0, gl.SPECULAR, &prjctr.das[0])
	gl.Lightfv(gl.LIGHT0, gl.SPOT_DIRECTION, &prjctr.spot_dir[0])
	gl.Lightf(gl.LIGHT0, gl.SPOT_CUTOFF, float32(prjctr.angle))
	gl.Lightf(gl.LIGHT0, gl.SPOT_EXPONENT, float32(prjctr.exp))
}

func draw(window glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.PolygonMode(gl.FRONT_AND_BACK, p_mod)

	w, h := window.GetSize()
	wd, ht := int32(w), int32(h)
	if wd > width || ht > height {
		if h > w {
			gl.Viewport((wd-ht)/2, 0, ht, ht)
		} else {
			gl.Viewport(0, (ht-wd)/2, wd, wd)
		}
	} else {
		gl.Viewport(0, 0, width, height)
	}

	gl.PointSize(2.0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadMatrixd(&projection3[0])

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.PushMatrix()
	setLight()
	gl.PopMatrix()

	gl.PushMatrix()
	gl.BindTexture(gl.TEXTURE_2D, texture)
	drawCylinder()
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.PopMatrix()

	gl.PushMatrix()
	drawCube()
	gl.PopMatrix()
}

//CallBacks

func key_callback(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		win.SetShouldClose(true)
	}
	//Rotation
	if key == glfw.KeyE && action == glfw.Repeat {
		roll += 1
	}
	if key == glfw.KeyQ && action == glfw.Repeat {
		roll -= 1
	}
	if key == glfw.KeyS && action == glfw.Repeat {
		pitch += 1
	}
	if key == glfw.KeyW && action == glfw.Repeat {
		pitch -= 1
	}
	if key == glfw.KeyA && action == glfw.Repeat {
		yaw += 1
	}
	if key == glfw.KeyD && action == glfw.Repeat {
		yaw -= 1
	}
	//Movement
	if key == glfw.KeyI && action == glfw.Repeat {
		cZ -= 0.05 * math.Cos(theta)
		cX += 0.05 * math.Sin(theta)
	}
	if key == glfw.KeyK && action == glfw.Repeat {
		cZ += 0.05 * math.Cos(theta)
		cX -= 0.05 * math.Sin(theta)
	}
	if key == glfw.KeyJ && action == glfw.Repeat {
		cX -= 0.05 * math.Cos(theta)
		cZ -= 0.05 * math.Sin(theta)
	}
	if key == glfw.KeyL && action == glfw.Repeat {
		cX += 0.05 * math.Cos(theta)
		cZ += 0.05 * math.Sin(theta)
	}
	if key == glfw.KeyY && action == glfw.Repeat {
		cY += 0.05 * math.Cos(phi)
	}
	if key == glfw.KeyH && action == glfw.Repeat {
		cY -= 0.05 * math.Cos(phi)
	}
	//Scale
	if key == glfw.Key0 && action == glfw.Repeat {
		sc += 0.05
	}
	if key == glfw.Key9 && action == glfw.Repeat {
		sc -= 0.05
	}
	//Quality
	if key == glfw.KeyZ && action == glfw.Press {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step + 1, cyl.h_step)
	}
	if key == glfw.KeyX && action == glfw.Press && cyl.e_step > 2 {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step - 1, cyl.h_step)
	}
	if key == glfw.KeyC && action == glfw.Press {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step, cyl.h_step + 1)
	}
	if key == glfw.KeyV && action == glfw.Press && cyl.h_step > 1 {
		cyl = *calcCylinder(cyl.a, cyl.b, cyl.h, cyl.e_step, cyl.h_step - 1)
	}
	//Light
	if key == glfw.KeyUp && action == glfw.Repeat && cyl.b < 1{
		lX += 1
		lY -= 1
	}
	if key == glfw.KeyDown && action == glfw.Repeat && cyl.b > 0.1{
		lX -= 1
		lY += 1
	}
	if key == glfw.KeyRight && action == glfw.Repeat && cyl.a < 1{
		lY -= 1
	}
	if key == glfw.KeyLeft && action == glfw.Repeat && cyl.a > 0.1{
		lY += 1
	}
	if key == glfw.KeyN && action == glfw.Press && cyl.h_step > 1 {
		prjctr.exp += 3.0
	}
	if key == glfw.KeyM && action == glfw.Press && cyl.h_step > 1 {
		prjctr.exp -= 3.0
	}
	if key == glfw.KeyO && action == glfw.Press && cyl.h_step > 1 {
		prjctr.angle += 3.0
	}
	if key == glfw.KeyP && action == glfw.Press && cyl.h_step > 1 {
		prjctr.angle -= 3.0
	}
	//Texture
	if key == glfw.Key0 && action == glfw.Press {
		if !isText {
			gl.Enable(gl.TEXTURE_2D)
		} else {
			gl.Disable(gl.TEXTURE_2D)
		}
		isText = !isText
	}
	//Mod
	if key == glfw.Key3 && action == glfw.Press {
		p_mod = gl.FILL
	}
	if key == glfw.Key4 && action == glfw.Press {
		p_mod = gl.LINE
	}
	if key == glfw.Key5 && action == glfw.Press {
		t = 0
		anim = !anim
	}
	//Reset
	if key == glfw.KeyR && action == glfw.Press{
		//cyl = *calcCylinder(0.8, 0.8, 0.8, 15, 8)
		cyl = *calcCylinder(0.8, 0.8, 0.8, 50, 60)
		//cyl = *calcCylinder(0.8, 0.8, 0.8, 75, 80)
		gl.NewList(1, gl.COMPILE)
		gl.Begin(gl.TRIANGLES)
		for _, p := range cyl.tris {
			for _, v := range p.verts {
				if isText {
					gl.TexCoord2fv(&v.tCoords[0])
				}
				gl.Normal3fv(&v.n[0])
				gl.Vertex3dv(&v.coords[0])
			}
		}
		gl.End()
		gl.EndList()


		sc = 1.0
		pitch, roll, yaw = 0, 0, 0
		cX, cY, cZ = 0, 0, 0
		lX, lY, lZ = 0, 0, 0
		prjctr.exp = 25
		p_mod, anim = gl.FILL, false
		t = 0
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

	program := initOpenGL()

	initTexture()

	rand.Seed(time.Now().UnixNano())
	c, calcTime := 0, 0.0
	readConfig()

	for !window.ShouldClose() {
		gl.UseProgram(program)
		t0 := time.Now()
		draw(*window)
		calcTime += time.Since(t0).Seconds()
		c++
		if c == 200 {
			fmt.Println(calcTime/float64(c))
			c, calcTime = 0, 0.0
		}
			glfw.PollEvents()
		window.SwapBuffers()
	}
}
*/