package server

import (
	"log"
	"net/http"

	"github.com/Scr3amz/EffectiveMobile/config"
	"github.com/Scr3amz/EffectiveMobile/internal/database/postgresql"
	"github.com/Scr3amz/EffectiveMobile/internal/transport/rest/handlers"
)

func RunServer(config config.Config) {
	mux := http.NewServeMux()
	store := postgresql.NewStore(config)
	handlers := handlers.NewFioHandler(*store, config)

	mux.HandleFunc("/fios/", handlers.FioHandler)
	
	log.Printf("Starting server at %s...\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, mux))

}
