package main

import (
	"log"
	"taskozon/internal/server"

	"github.com/BurntSushi/toml"
)

func main() {
	config := server.NewConfig()
	_, err := toml.DecodeFile("configs/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(config.MaxSize)
	server.Start(config.Port)
}
