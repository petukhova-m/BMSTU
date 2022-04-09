package main
 
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
 
	"github.com/jlaffaye/ftp"
)
 
func main() {
	connect("localhost:9015", "admin", "111111")
	fmt.Println("===============")
	connect("students.yss.su:21", "ftpiu8", "3Ru7yOTA")
}
 
func connect(host, username, pass string) {
	c, err := ftp.Connect(host)
	if err != nil {
		log.Fatal(err)
	}
 
	err = c.Login(username, pass)
	if err != nil {
		log.Fatal(err)
	}
 
	log.Println("Succefly connected to: ", c)
 
	list := listOfData("/", c)
	//Вывод всех данных
	for i := range list {
		fmt.Println(list[i] + " \n")
	}
 
	//Удаление папки
	res := rmDir("petukhova", "/", c)
	fmt.Println(res)
 
	//Создание папки
	result := mkDir("petukhova", "/", c)
	fmt.Println(result)
 
	//Переход в папку
	error := c.ChangeDir("/petukhova")
	if error != nil {
		log.Fatal(error)
	}
 
	//Cоздаем файл, который потом загружаем в мою папку
	data := bytes.NewBufferString("It works")
	err = c.Stor("Test.txt", data)
	if err != nil {
		log.Fatal(err)
	}
 
	list = listOfData("/petukhova", c)
	//Вывод всех данных
	for i := range list {
		fmt.Println(list[i] + " \n")
	}
 
	//ниже код отвечаюший за скачивание файлы с FTP
	mkFile("Test.txt", c)
	fmt.Println("file copied")
 
	//удаление файла
	/* err = c.Delete("/petukhova/Test1.txt")
	if err != nil {
		log.Fatal(err)
	} */
 
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
 
func rmDir(dir, path string, c *ftp.ServerConn) string {
	exist := false
	list := listOfData(path, c)
 
	for _, value := range list {
		if value == path+dir {
			exist = true
		}
	}
 
	if exist {
		err := c.RemoveDir(path + dir)
		return "directory deleted"
		if err != nil {
			log.Fatal(err)
		}
	}
	return "such directory does not  exist"
}
 
func mkFile(file string, c *ftp.ServerConn) {
	res, err := c.Retr(file)
	if err != nil {
		log.Fatal(err)
	}
 
	defer res.Close()
 
	outFile, err := os.Create("zaitsev.txt")
	if err != nil {
		log.Fatal(err)
	}
 
	defer outFile.Close()
 
	_, err = io.Copy(outFile, res)
	if err != nil {
		log.Fatal(err)
	}
}
 
func listOfData(str string, c *ftp.ServerConn) []string {
 
	array, err := c.NameList(str) //получение данных о том что есть на сервере
	if err != nil {
		log.Fatal(err)
	}
	return array
}
 
func mkDir(dir string, path string, c *ftp.ServerConn) string {
	exist := false
	list := listOfData(path, c)
 
	for _, value := range list {
		if value == path+dir {
			exist = true
		}
	}
 
	if !exist {
		error := c.MakeDir(dir)
		return "directory was created"
		if error != nil {
			log.Fatal(error)
		}
	}
	return "such directory is already exist"
}