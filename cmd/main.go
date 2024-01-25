package main

import (
	"log"

	"github.com/Scr3amz/EffectiveMobile/config"
	"github.com/Scr3amz/EffectiveMobile/internal/server"
)

func main() {
	config, err := config.LoadMainConfig("", "env")
	if err != nil {
		log.Fatalf("Error occured while reading the config file\nError: %v\n", err)
	}
	server.RunServer(config)

}
