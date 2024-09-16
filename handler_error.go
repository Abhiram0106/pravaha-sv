package main

import (
	"net/http"

	internal "github.com/Abhiram0106/pravaha-sv/internal/api"
)

func handlerError(w http.ResponseWriter, r *http.Request) {
	internal.RespondWithError(
		w,
		http.StatusInternalServerError,
		"Internal Server Error",
	)
}
