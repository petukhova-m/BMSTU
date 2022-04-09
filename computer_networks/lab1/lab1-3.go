package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func connectLocal() (c *ftp.ServerConn) {
	c, err := ftp.Dial("localhost:9008", ftp.DialWithTimeout(5*time.Second)) // addr host:port соединяемся с ftp-сервером по ука
	//занному адресу

	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("petukhova", "masha") // авторизуемся на указанном ftp-сервере

	if err != nil {
		log.Fatal(err)
	}
	return c
}

func connectStudents() (c *ftp.ServerConn) {
	c, err := ftp.Dial("students.yss.su:21", ftp.DialWithTimeout(5*time.Second))

	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("ftpiu8", "3Ru7yOTA") // авторизуемся на указанном ftp-сервере

	if err != nil {
		log.Fatal(err)
	}
	return c
}

func quitConn(c *ftp.ServerConn) {
	if err := c.Quit(); err != nil { // закрываем соединение с сервером
		log.Fatal(err)
	}
}

func main() {
	rmdir1()
}

func localToStudents() {
	var filename, path string
	fmt.Println("Input name of your file on localhost")
	fmt.Fscan(os.Stdin, &filename)

	c := connectLocal() // connect to localhost
	pwdftp, err := c.CurrentDir()
	r, err := c.Retr(pwdftp + filename) // download from localhost
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close() // postpone closing file
	quitConn(c)     // quit connection from localhost

	c = connectStudents() // connect to students

	fmt.Println("Input path for your file on students:")
	fmt.Fscan(os.Stdin, &path)

	err = c.Stor(path+filename, r) // upload file on students
	if err != nil {
		log.Fatal(err)
	}
	quitConn(c) // quit connection from students
}

func studentsToLocal() {
	var filename, path string
	fmt.Println("Input path of your file on students")
	fmt.Fscan(os.Stdin, &path)

	fmt.Println("Input name of your file on students")
	fmt.Fscan(os.Stdin, &filename)

	c := connectStudents()            // connect to students
	r, err := c.Retr(path + filename) // download from students
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close() // postpone closing file
	quitConn(c)     // quit connection from students

	c = connectLocal() // connect to localhost

	pwdftp, err := c.CurrentDir()

	err = c.Stor(pwdftp+filename, r) // upload file on localhost
	if err != nil {
		log.Fatal(err)
	}
	quitConn(c) // quit connection from localhost
}

func mkdir1() {
	var ftpdir string
	c := connectLocal() // connect to local
	cpwd, err := c.CurrentDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input name for creating directory:")
	fmt.Fscan(os.Stdin, &ftpdir) // вводим с клавиатуры название для новой директории
	cpwd = cpwd + ftpdir         // full path for new directory in students
	c = connectStudents()        // connect to students
	err = c.MakeDir(cpwd)        // create dir
	if err != nil {
		log.Fatal(err)
	}
}

func rmdir1() {
	var ftpdir string
	c := connectLocal() // connect to local
	fmt.Println("Input path for deleting directory:")
	fmt.Fscan(os.Stdin, &ftpdir) // вводим с клавиатуры название для удаляемой директории
	c = connectStudents()        // connect to students
	err := c.RemoveDir(ftpdir)   // remove dir from students
	if err != nil {
		log.Fatal(err)
	}
}
