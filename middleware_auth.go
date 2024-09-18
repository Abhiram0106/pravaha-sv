package main

import (
	"net/http"

	internal "github.com/Abhiram0106/pravaha-sv/internal/api"
	"github.com/Abhiram0106/pravaha-sv/internal/auth"
	"github.com/Abhiram0106/pravaha-sv/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.GetUserByApiKeyRow)

func (apiCfg *apiConfig) middlewareAuth(authedHandler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, getApiKeyErr := auth.GetApiKey(r.Header)

		if getApiKeyErr != nil {
			internal.RespondWithError(w, http.StatusUnauthorized, getApiKeyErr.Error())
			return
		}

		user, getUserErr := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if getUserErr != nil {
			internal.RespondWithError(w, http.StatusNotFound, getUserErr.Error())
			return
		}

		authedHandler(w, r, user)
	}
}
