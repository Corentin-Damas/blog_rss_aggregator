package main

import (
	"net/http"
	"fmt"

	"github.com/Corentin-Damas/rssagg/internal/database"
	"github.com/Corentin-Damas/rssagg/internal/auth"
	
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error : %s", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user : %s", err))
			return
		}

		handler(w, r, user)
	}
}