package main

import (
	"github.com/saddam-satria/go-boilerplate/cmd/student"
	server "github.com/saddam-satria/go-boilerplate/configs"
)

func main() {
	server.Init(student.Init)
}
