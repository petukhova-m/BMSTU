package main

import (
	"github.com/gliderlabs/ssh"
	"io"
	"log"
	"os/exec"
)

// Handler ssh
func Handler(s ssh.Session) {

	cmd := s.Command()
	log.Println("Command handled: ", cmd)
	if len(cmd) != 0 {
		if out, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput(); err != nil {
			io.WriteString(s, err.Error())
		} else {
			io.WriteString(s, string(out))
		}
	}
}


func main() {
	ssh.Handle(Handler)

	log.Println("starting ssh server on port 2222")
	log.Fatal(ssh.ListenAndServe(":2222", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			return pass == "qwerty"
		}),
	))
}
