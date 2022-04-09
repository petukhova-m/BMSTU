package main

import (
	"flag"
	"fmt"
	"os"
	"bytes"
	"github.com/skorobogatov/input"
	"golang.org/x/crypto/ssh"
)

// creating command-line flags with default values
var (
	user = flag.String("u", "ivanov", "User name")
	host = flag.String("h", "lab2.posevin.com", "Host")
	port = flag.Int("p", 22, "port")
)

func main() {
	var password string
	fmt.Println("Input password:")
	fmt.Fscan(os.Stdin, &password)
	//creating config for ssh-client
	config := &ssh.ClientConfig{
		//specify username
		User: *user,
		Auth: []ssh.AuthMethod{
			//specify authentification method by password
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//making string type of 'host:port'
	addr := fmt.Sprintf("%s:%d", *host, *port)

	//calling to server
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(err)
	}

	for {
		// creating ssh session
		session, err := client.NewSession()
		if err != nil {
			panic(err)
		}

		// always remember to close session
		defer session.Close()
		fmt.Printf("Input a command: ")
		cmd := input.Gets()
		if cmd == "quit" {
			break
		}
		// copy and print output stream
		var stdoutBuf bytes.Buffer
		session.Stdout = &stdoutBuf
		session.Run(cmd)
		fmt.Println(stdoutBuf.String())
	}
}
