package server

import (
	"net/http"

	"github.com/Scr3amz/EffectiveMobile/config"
	"github.com/Scr3amz/EffectiveMobile/internal/database/postgresql"
	"github.com/Scr3amz/EffectiveMobile/internal/transport/rest/handlers"
	"github.com/Scr3amz/EffectiveMobile/logger"
)

func RunServer(config config.Config, logger logger.Logger) {
	mux := http.NewServeMux()
	store := postgresql.NewStore(config, logger)
	handlers := handlers.NewFioHandler(*store, config, logger)

	mux.HandleFunc("/fios/", handlers.FioHandler)

	logger.InfoLog.Printf("Starting server at %s...\n", config.ServerPort)
	logger.ErrorLog.Fatal(http.ListenAndServe(config.ServerPort, mux))

}
