package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Demo using golang.org/x/crypto")
_ = &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.Password("dummy"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}
