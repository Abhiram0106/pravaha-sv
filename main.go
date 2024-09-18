package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Abhiram0106/pravaha-sv/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if secretsErr := godotenv.Load(); secretsErr != nil {
		log.Fatal(secretsErr)
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	db, dbConnectionErr := sql.Open("postgres", dbURL)

	if dbConnectionErr != nil {
		log.Fatalf("database failed to connect : %s\n", dbConnectionErr.Error())
	}
	if port == "" {
		log.Fatalln("port is empty")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		port: port,
		DB:   dbQueries,
	}
	startServer(&apiCfg)
}
