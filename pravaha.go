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

	muxV1 := http.NewServeMux()

	muxV1.HandleFunc("/healthz", handlerReadiness)
	muxV1.HandleFunc("/err", handlerError)

	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", muxV1))
	server := http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: mux,
	}

	log.Printf("Listening on port: %s\n", apiCfg.port)
	log.Fatalln(server.ListenAndServe())
}
