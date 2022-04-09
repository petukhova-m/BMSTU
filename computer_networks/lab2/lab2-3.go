package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

var (
	command,
	commandArgs string
	)

func main() {

	config_1 := &ssh.ClientConfig{
		User: "maxim" ,
		Auth: []ssh.AuthMethod{
			ssh.Password("qwerty"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	config_2 := &ssh.ClientConfig{
		User: "ivanov" ,
		Auth: []ssh.AuthMethod{
			ssh.Password("monetochka"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client_1, err := ssh.Dial("tcp", "localhost:2222", config_1)
	if err != nil {
		panic(err)
	}

	client_2, err := ssh.Dial("tcp", "lab2.posevin.com:22", config_2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Вводите команды <сервер> <команда>" )

	var work = true


	for work {

		session_1, err := client_1.NewSession()
		if err != nil {
			panic(err)
		}

		session_2, err := client_2.NewSession()
		if err != nil {
			panic(err)
		}

		var serv_num int

		fmt.Scan(&serv_num)

		fmt.Scan(&command)

		if serv_num == 1{
			switch command {

			case "mkdir":
				fmt.Scan(&commandArgs)
				session_1.Run(command + " " + commandArgs)

			case "rmdir":
				fmt.Scan(&commandArgs)
				session_1.Run(command + " " + commandArgs)
			case "touch":
				fmt.Scan(&commandArgs)
				session_1.Run(command + " " + commandArgs)

			case "rm":
				fmt.Scan(&commandArgs)
				session_1.Run(command + " " + commandArgs)

			case "show":
				fmt.Scan(&commandArgs)
				var b, err = session_1.Output("ls" + commandArgs)
				if err != nil {
					panic(err)
				}
				fmt.Print(string(b))

			case "exit":
				work = false

			default:
				var b, err = session_1.Output(command)
				if err != nil {
					panic(err)
				}
				fmt.Print(string(b))

				defer session_1.Close()
				defer session_2.Close()
			}
		}

		if serv_num == 2{
			switch command {

			case "mkdir":
				fmt.Scan(&commandArgs)
				session_2.Run(command + " " + commandArgs)

			case "rmdir":
				fmt.Scan(&commandArgs)
				session_2.Run(command + " " + commandArgs)
			case "touch":
				fmt.Scan(&commandArgs)
				session_2.Run(command + " " + commandArgs)

			case "rm":
				fmt.Scan(&commandArgs)
				session_2.Run(command + " " + commandArgs)

			case "show":
				fmt.Scan(&commandArgs)
				var b, err = session_1.Output("ls" + commandArgs)
				if err != nil {
					panic(err)
				}
				fmt.Print(string(b))

			case "exit":
				work = false

			default:
				var b, err = session_2.Output(command)
				if err != nil {
					panic(err)
				}
				fmt.Print(string(b))

				defer session_1.Close()
				defer session_2.Close()
			}
		}
	}
}