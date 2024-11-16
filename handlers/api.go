// Package handlers defines HTTP handlers for the web chat application.
package handlers

import (
	"encoding/json"
	"net/http"
)

// ClientCount represents the JSON structure for returning the number of connected clients.
type ClientCount struct {
	Count int `json:"count"`
}

// WelcomeMessage represents the JSON structure for returning a welcome message.
type WelcomeMessage struct {
	Message string `json:"message"`
}

// HandleClientCount handles requests for the client count API, returning the current number of connected clients.
func HandleClientCount(w http.ResponseWriter, r *http.Request) {
	count := len(clients) // Calculate the number of connected clients.
	response := ClientCount{Count: count}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // Send JSON response with client count.
}

// HandleWelcomeMessage handles requests for the welcome message API, returning a welcome message.
func HandleWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	welcome := WelcomeMessage{Message: "Welcome to the enhanced Go web chat!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(welcome) // Send JSON response with welcome message.
}
