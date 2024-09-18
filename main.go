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
	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	db, dbConnectionErr := sql.Open("postgres", dbURL)

	if dbConnectionErr != nil {
		log.Fatalf("database failed to connect : %s\n", dbConnectionErr.Error())
	}

	if jwtSecret == "" {
		log.Fatalln("jwtSecret is empty")
	}
	if port == "" {
		log.Fatalln("port is empty")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		jwtSecret: jwtSecret,
		port:      port,
		DB:        dbQueries,
	}
	startServer(&apiCfg)
}
