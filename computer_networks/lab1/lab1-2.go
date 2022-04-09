package main

import (
	"fmt"
	"log"
	"os"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

const login string = "petukhova"
const password string = "masha"
const serverPort int = 9008

func main() {
	pwd, err := os.Getwd() //дефолтная папка где и находится ftp сервер
	if err != nil {
		log.Fatal(err)
	}

	path := "serverPath"
	//настройка файлового драйвера (название папки где все будет работать и настройка для коректного создания и изменения)
	factory := &filedriver.FileDriverFactory{
		RootPath: fmt.Sprintf("%s/%s/", pwd, path),
		Perm:     server.NewSimplePerm("user", "group"),
	}

	//настройки сервера
	opts := &server.ServerOpts{
		Factory:  factory,     //настройки с директорией
		Port:     serverPort,  //порт запуска
		Hostname: "localhost", //url или доменное имя можно ip
		Auth: &server.SimpleAuth{
			Name:     login,
			Password: password,
		}, //авторизация
	}

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Login %v, Password %v", login, password)

	//создаем объект сервера
	server := server.NewServer(opts)
	//запускаем сервер
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
