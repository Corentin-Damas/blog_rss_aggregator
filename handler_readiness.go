package main

import "net/http"


// HandleReadiness Report if a serveur is ready and live
func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, 200, struct{}{})
}