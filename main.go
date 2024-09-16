package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if secretsErr := godotenv.Load(); secretsErr != nil {
		log.Fatal(secretsErr)
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")

	if jwtSecret == "" {
		log.Fatalln("jwtSecret is empty")
	}
	if port == "" {
		log.Fatalln("port is empty")
	}

	apiCfg := apiConfig{
		jwtSecret: jwtSecret,
		port:      port,
	}
	startServer(&apiCfg)
}
