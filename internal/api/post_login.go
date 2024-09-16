package internal

import (
	"github.com/Abhiram0106/pravaha-sv/main"
	"net/http"
)

func handlerPostUser(w http.ResponseWriter, r *http.Request) {

	huh := main.ApiConfig
	huh.jwtSecret

}
