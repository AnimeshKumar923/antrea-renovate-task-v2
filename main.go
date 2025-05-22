package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Demonstrate vulnerable dependency usage.")
	_ = ssh.ClientConfig{}
}
