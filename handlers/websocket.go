// Package handlers defines WebSocket handlers for the web chat application.
package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"web-chat-backend/models"
)

var (
	clients   = make(map[*websocket.Conn]bool) // Connected clients.
	broadcast = make(chan models.Message)      // Channel for broadcasting messages to clients.

	// Upgrader is used to upgrade HTTP requests to WebSocket connections.
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow connections from any origin for development purposes.
		// For production, this should be restricted to trusted origins.
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// HandleWebSocket upgrades an HTTP connection to a WebSocket and manages communication with the client.
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil) // Upgrade to WebSocket connection.
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true // Register new client.
	log.Printf("New connection. Total clients: %d", len(clients))

	for {
		var msg models.Message
		// Read incoming message as JSON.
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("JSON read error: %v", err)
			delete(clients, ws) // Remove client on read error.
			log.Printf("Client disconnected. Total clients: %d", len(clients))
			break
		}
		broadcast <- msg // Send message to broadcast channel.
	}
}

// init starts a goroutine to handle message broadcasting.
func init() {
	go HandleMessages()
}

// HandleMessages listens to the broadcast channel and sends messages to all connected clients.
func HandleMessages() {
	for {
		// Retrieve a message from the broadcast channel.
		msg := <-broadcast
		// Send message to all connected clients.
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("JSON write error: %v", err)
				client.Close()
				delete(clients, client) // Remove client on write error.
				log.Printf("Client disconnected. Total clients: %d", len(clients))
			}
		}
	}
}
