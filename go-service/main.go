// main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    
    "github.com/gorilla/mux"
)

// Define a basic response structure
type HealthResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    // Create a response
    response := HealthResponse{
        Status:  "healthy",
        Message: "API is running",
    }
    
    // Set content type header
    w.Header().Set("Content-Type", "application/json")
    
    // Encode and send the response
    json.NewEncoder(w).Encode(response)
}

func main() {
    // Create a new router
    router := mux.NewRouter()
    
    // Register routes
    router.HandleFunc("/health", healthCheck).Methods("GET")
    
    // Start the server
    log.Printf("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}