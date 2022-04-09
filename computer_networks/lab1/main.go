
package main

import (
	"log"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

func main() {
	var (
		root = "/"
		user = "admin"
		pass = "111111"
		port = 9015
		host = "localhost"
	)
	if root == "" {
		log.Fatalf("Please set a root to serve with -root")
	}

	factory := &filedriver.FileDriverFactory{
		RootPath: root,
		Perm:     server.NewSimplePerm("admin", "group"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     port,
		Hostname: host,
		Auth:     &server.SimpleAuth{Name: user, Password: pass},
	}

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Username %v, Password %v", user, pass)
	server := server.NewServer(opts)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
