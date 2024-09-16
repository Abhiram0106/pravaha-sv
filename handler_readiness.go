package main

import (
	"net/http"

	internal "github.com/Abhiram0106/pravaha-sv/internal/api"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	internal.RespondWithJson(
		w,
		http.StatusOK,
		struct {
			Status string `json:"status"`
		}{Status: "ok"},
	)
}
