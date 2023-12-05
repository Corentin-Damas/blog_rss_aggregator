package main

import "net/http"

// HandleErr Everytime we need to report an error we can use this output
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Somthing went Wrong")
}
