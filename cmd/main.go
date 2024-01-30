package main

import (
	"github.com/Scr3amz/EffectiveMobile/config"
	"github.com/Scr3amz/EffectiveMobile/internal/server"
	"github.com/Scr3amz/EffectiveMobile/logger"
)

func main() {
	logger := logger.NewLogger("log.txt")
	config, err := config.LoadMainConfig("", "env")
	if err != nil {
		logger.ErrorLog.Fatalf("Error occured while reading the config file\nError: %v\n", err)
	}
	server.RunServer(config, logger)
}
