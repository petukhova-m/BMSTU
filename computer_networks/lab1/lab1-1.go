package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"bytes"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
)

func main() {
	c, err := ftp.Dial("students.yss.su:21", ftp.DialWithTimeout(5*time.Second)) // addr host:port соединяемся с ftp-сервером по ука
	//занному адресу

	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("ftpiu8", "3Ru7yOTA") // авторизуемся на указанном ftp-сервере

	if err != nil {
		log.Fatal(err)
	}

	// stor(c) // загружает файл на сервер
	// retr(c) // скачивает файл c серверa
	mkdir(c) // создает директорию на сервере
	delete(c) // удаляет файл на сервере
	ls(c) // запрашивает список содержимого в текущей директории на сервере

	if err := c.Quit(); err != nil { // закрываем соединение с сервером
		log.Fatal(err)
	}
}

func stor(c *ftp.ServerConn) {
	var localfile, ftpfile, path string
	fmt.Println("Input filename from local computer:")
	fmt.Fscan(os.Stdin, &localfile)

	fmt.Println("Input path for your file on ftp-server:")
	fmt.Fscan(os.Stdin, &path)

	fmt.Println("Input filename for your local file on ftp-server:")
	fmt.Fscan(os.Stdin, &ftpfile)
	file, err := os.Open(localfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // закрываем файл в конце выполнения функции

	// получаем размер файла
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bs) // преобразуем массив байтов в строку
	data := bytes.NewBufferString(str)
	err = c.Stor(path+ftpfile, data)
	if err != nil {
		log.Fatal(err)
	}
}

func retr(c *ftp.ServerConn) {
	var localfile, ftpfile, path string
	fmt.Println("Input path of your file on ftp-server:")
	fmt.Fscan(os.Stdin, &path)
	fmt.Println("Input filename of your file on ftp-server:")
	fmt.Fscan(os.Stdin, &ftpfile)
	fmt.Println("Input name for saving your file on local computer:")
	fmt.Fscan(os.Stdin, &localfile)

	pwd, err := os.Getwd() // запрашиваем текущую директорию
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pwd)

	file, err := os.Create(pwd + "/" + localfile) // создаем локальный файл
	if err != nil {
		log.Fatal(err)
	}

	r, err := c.Retr(path + ftpfile) // скачиваем файл с ftp-сервера
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close() // закрываем файл по окончании выполнения функции

	bs, err := ioutil.ReadAll(r) // считываем массив байтов
	str := string(bs) // преобразуем массив байтов в строку
	file.WriteString(str) // записываем в файл
}

func mkdir(c *ftp.ServerConn) {
	var ftpdir string
	fmt.Println("Input name for creating directory:")
	fmt.Fscan(os.Stdin, &ftpdir) // вводим с клавиатуры название для новой директории
	pwdftp, err := c.CurrentDir() // запрашиваем текущую директорию на ftp-сервере
	if err != nil {
		log.Fatal(err)
	}
	err = c.MakeDir(pwdftp + ftpdir) // создаем директорию на ftp-сервере
	if err != nil {
		log.Fatal(err)
	}
}

func delete(c *ftp.ServerConn) {
	var path string
	fmt.Println("Input full path for file that you wanna delete from frp-server:")
	fmt.Fscan(os.Stdin, &path)
	err := c.Delete(path) // удаляем файл по заданному пути
	if err != nil {
		log.Fatal(err)
	}
}

func ls(c *ftp.ServerConn) {
	pwdftp, err := c.CurrentDir() // запрашиваем текущую директорию на ftp-сервере
	if err != nil {
		log.Fatal(err)
	}

	entries, err := c.List(pwdftp) // запрашиваем содержимое текущей директории на ftp-сервере
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries { // выводим имя и размер каждого элемента массива вхождений
		fmt.Println(entry.Name)
		fmt.Println(entry.Size)
	}
}