package main

import (
	"taskozon/internal/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
