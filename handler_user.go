package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	internal "github.com/Abhiram0106/pravaha-sv/internal/api"
	"github.com/Abhiram0106/pravaha-sv/internal/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (apiCfg *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	loginJson, readErr := io.ReadAll(r.Body)

	if readErr != nil {
		log.Printf("Failed to read create user request %v\n", readErr.Error())
		internal.RespondWithError(w, http.StatusBadRequest, readErr.Error())
		return
	}

	loginRequest := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}
	unMarshalErr := json.Unmarshal(loginJson, &loginRequest)

	if unMarshalErr != nil {
		log.Printf("Failed to unmarshal create user request %v\n", unMarshalErr.Error())
		internal.RespondWithError(w, http.StatusBadRequest, unMarshalErr.Error())
		return
	}

	if loginRequest.Name == "" || loginRequest.Password == "" {
		internal.RespondWithError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	passwordBytes := []byte(loginRequest.Password)
	if len(passwordBytes) > 72 {
		internal.RespondWithError(w, http.StatusBadRequest, "Password can't be greater than 72 bytes")
		return
	}

	hashedPassword, hashingErr := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if hashingErr != nil {
		internal.RespondWithError(w, http.StatusInternalServerError, hashingErr.Error())
		return
	}

	newUser, createUserErr := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      loginRequest.Name,
		Password:  string(hashedPassword),
	})

	if createUserErr != nil {
		log.Printf("Failed to create user %v\n", createUserErr.Error())
		internal.RespondWithError(w, http.StatusInternalServerError, createUserErr.Error())
		return
	}

	internal.RespondWithJson(w, http.StatusCreated, databaseUserToUser(&newUser))
}
