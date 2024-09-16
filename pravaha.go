package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	port      string
	jwtSecret string
}

func startServer(apiCfg *apiConfig) {

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: mux,
	}

	log.Printf("Listening on port: %s\n", apiCfg.port)
	log.Fatalln(server.ListenAndServe())
}
