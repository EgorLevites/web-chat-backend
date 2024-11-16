// main.go is the entry point for the web chat application.
package main

import (
	"log"
	"net/http"

	"web-chat-backend/handlers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешить запросы с любого источника (для тестов, в продакшене лучше ограничить конкретным доменом)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	// Serve static files
	fs := http.FileServer(http.Dir("../web-chat-frontend"))

	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("/ws", handlers.HandleWebSocket)
	mux.HandleFunc("/api/clients", handlers.HandleClientCount)
	mux.HandleFunc("/api/welcome", handlers.HandleWelcomeMessage)

	// Wrap the mux with CORS middleware
	http.Handle("/", enableCORS(mux))

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error:", err)
	}

}
