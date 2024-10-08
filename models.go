package main

import (
	"time"

	"github.com/Abhiram0106/pravaha-sv/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser *database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseUsersToUsers(dbUsers []database.User) []User {
	users := make([]User, len(dbUsers))

	for i, dbUser := range dbUsers {
		users[i] = databaseUserToUser(&dbUser)
	}

	return users
}
