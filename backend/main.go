package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Message struct to parse user input
type Message struct {
	Content string `json:"content"`
}

// Response struct for Valeera's output
type Response struct {
	Content string `json:"content"`
	Reply   string `json:"reply"`
}

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %s", time.Since(start))
	})
}

// Middleware to handle CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Allow requests from your frontend
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Handler for Valeera's interaction
func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Simulate Valeera's reply
	reply := fmt.Sprintf("Valeera says: %s", generateReply(msg.Content))

	// Prepare response
	response := Response{
		Content: msg.Content,
		Reply:   reply,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Simple function to generate replies (replace with more complex logic later)
func generateReply(input string) string {
	switch input {
	case "hi":
		return "Hey, what can i help with?"
	case "how are you":
		return "I don't do, 'How', but thanks for asking!"
	default:
		return "I'm not sure how to respond to that, but I'm here to help!"
	}
}

func main() {
	// Setting up the HTTP server
	mux := http.NewServeMux()

	// API endpoint for Valeera's interaction
	mux.HandleFunc("/api/message", messageHandler)

	// Wrapping the server with middleware for logging and CORS
	loggedAndCORS := loggingMiddleware(enableCORS(mux))

	fmt.Println("Valeera is ready at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", loggedAndCORS))
}
