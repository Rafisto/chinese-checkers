package main

import (
	"chinese-checkers/server"
)

func main() {
	s := server.NewServer()

	s.RunServer(8080)
}
