package main

import (
	"encoding/json"
	"net/http"
)

type SnippetResponse struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func codingSnippetHandler(w http.ResponseWriter, r *http.Request) {
	// Hardcoded code snippets
	snippets := map[string]SnippetResponse{
		"typescript": {"TypeScript", "const add = (a: number, b: number): number => a + b;"},
		"python":     {"Python", "def add(a, b): return a + b"},
		"go":         {"Go", "func add(a, b int) int { return a + b }"},
	}

	lang := r.URL.Query().Get("lang")
	if snippet, exists := snippets[lang]; exists {
		json.NewEncoder(w).Encode(snippet)
	} else {
		http.Error(w, "Snippet not found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/snippet", codingSnippetHandler)
	http.ListenAndServe(":8080", nil)
}
