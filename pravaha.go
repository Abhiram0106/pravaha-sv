package main

import (
	"log"
	"net/http"

	"github.com/Abhiram0106/pravaha-sv/internal/database"
)

type apiConfig struct {
	port string
	DB   *database.Queries
}

func startServer(apiCfg *apiConfig) {

	muxV1 := http.NewServeMux()

	muxV1.HandleFunc("GET /healthz", handlerReadiness)
	muxV1.HandleFunc("GET /err", handlerError)
	muxV1.HandleFunc("POST /users", apiCfg.handlerUserCreate)

	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", muxV1))
	server := http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: mux,
	}

	log.Printf("Listening on port: %s\n", apiCfg.port)
	log.Fatalln(server.ListenAndServe())
}
