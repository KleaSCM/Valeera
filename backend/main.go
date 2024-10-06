package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Snippet struct to hold the code snippets
type Snippet struct {
	Description string `json:"description"`
	Code        string `json:"code"`
}

// Global variable to hold the snippets
var Snippets map[string]Snippet

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
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin for testing
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent) // Use NoContent for OPTIONS
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Function to get a response from OpenAI
func getOpenAIResponse(question string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key is not set")
	}

	client := &http.Client{}
	reqBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": question},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("Error marshaling request body:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating new request:", err)
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request to OpenAI:", err)
		return "", err
	}
	defer resp.Body.Close()

	log.Println("OpenAI API Response Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OpenAI API returned status: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding response body:", err)
		return "", err
	}

	// Extracting the response text
	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if message, ok := choices[0].(map[string]interface{}); ok {
			if content, ok := message["message"].(map[string]interface{})["content"].(string); ok {
				return content, nil
			}
		}
	}

	return "", fmt.Errorf("failed to get a valid response from OpenAI")
}

// Simple function to determine if the message is a question
func isQuestion(input string) bool {
	return len(input) > 0 && input[len(input)-1] == '?'
}

// Handler for Valeera's interaction
func messageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println("Error decoding message:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	log.Printf("Received message: %s", msg.Content)

	var reply string
	if isQuestion(msg.Content) {
		// Check if the question exists in the snippets
		if snippet, exists := Snippets[msg.Content]; exists {
			reply = snippet.Code // Use the snippet if it exists
		} else {
			// Only call OpenAI if the question is not found in snippets
			reply, err = getOpenAIResponse(msg.Content)
			if err != nil {
				log.Println("Failed to get a response from OpenAI:", err)
				http.Error(w, "Failed to get a response from OpenAI", http.StatusInternalServerError)
				return
			}
		}
	} else {
		reply = generateReply(msg.Content) // Generate reply using snippets for non-questions
	}

	response := Response{
		Content: msg.Content,
		Reply:   reply,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to generate replies using snippets
func generateReply(input string) string {
	if snippet, exists := Snippets[input]; exists { // Use Snippets
		return snippet.Code // Return the code snippet directly
	}
	return "I'm not sure how to respond to that, but I'm here to help!"
}

// Handler for updating snippets
func updateSnippetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming request body
	var newSnippet struct {
		Question string `json:"question"`
		Response string `json:"response"`
	}

	err := json.NewDecoder(r.Body).Decode(&newSnippet)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Load existing snippets
	if err := loadSnippets(); err != nil {
		http.Error(w, "Could not load snippets", http.StatusInternalServerError)
		return
	}

	// Add or update the question with the new response
	Snippets[newSnippet.Question] = Snippet{ // Use Snippets
		Description: newSnippet.Question,
		Code:        newSnippet.Response,
	}

	// Save the updated snippets back to file
	fileToSave, err := os.Create("snippets.json")
	if err != nil {
		log.Println("Error creating snippets.json:", err)
		http.Error(w, "Could not save snippets", http.StatusInternalServerError)
		return
	}
	defer fileToSave.Close()

	encoder := json.NewEncoder(fileToSave)
	encoder.SetIndent("", "  ")                       // Make the JSON pretty
	if err := encoder.Encode(&Snippets); err != nil { // Use Snippets
		log.Println("Error encoding updated snippets.json:", err)
		http.Error(w, "Failed to save the snippets", http.StatusInternalServerError)
		return
	}

	// Successfully added or updated the snippet
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Snippet successfully added or updated"))
}

// Load snippets from a JSON file
func loadSnippets() error {
	file, err := os.Open("snippets.json")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&Snippets); err != nil {
		return err
	}

	return nil
}

// Load environment variables from a .env file
func loadEnv() error {
	// Implement your environment loading logic here if needed
	return nil
}

func main() {
	// Load environment variables from .env file
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Load snippets from JSON file
	Snippets = make(map[string]Snippet) // Initialize Snippets
	if err := loadSnippets(); err != nil {
		log.Fatalf("Failed to load snippets: %v", err)
	}

	// Set up HTTP routes and middleware
	http.Handle("/api/message", enableCORS(loggingMiddleware(http.HandlerFunc(messageHandler))))
	http.Handle("/api/update", enableCORS(loggingMiddleware(http.HandlerFunc(updateSnippetsHandler))))

	// Enable CORS for all routes
	http.Handle("/", enableCORS(http.DefaultServeMux))

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}
	log.Printf("Server is running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
