package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/", handleRequest) 

	// Create a new HTTP server with CORS enabled
	http.ListenAndServe(":8080", enableCors(router))
}

func enableCors(handler http.Handler) http.Handler {
	corsObj := handlers.AllowedOrigins([]string{"*"}) // Add the allowed origins here
	corsObj = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}) // Add the allowed methods here
	corsObj = handlers.AllowedHeaders([]string{"Content-Type"}) // Add the allowed headers here
	return handlers.CORS(corsObj)(handler)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "This is a CORS-enabled API response"}`))
}
