package server

import (
	"encoding/json"
	"log"
	"net/http"
	"orchestrator/LLM"
	"strings"
	"time"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{Port: port}
}

func Handler(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MiB max
	defer r.Body.Close()                           // ensure body is closed

	if ct := r.Header.Get("Content-Type"); ct != "" && !strings.Contains(ct, "application/json") {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var message LLM.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		// Distinguish size errors vs bad JSON
		if _, ok := err.(*http.MaxBytesError); ok {
			http.Error(w, "payload too large", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	prompt := LLM.ChatRequest{
		Model: "gpt-oss:120b-cloud",
		Messages: []LLM.Message{
			{Role: "user", Content: message.Content},
		},
		Options: LLM.Options{Temperature: 0.4},
		Format: LLM.Format{
			Type:       "json",
			Properties: map[string]any{"Tool": "string", "Params": "string"},
		},
	}

	client := LLM.NewOllamaClient()
	resp, err := client.Chat(prompt)
	if err != nil {
		log.Printf("LLM call failed: %v", err)
		http.Error(w, "upstream LLM error", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (s *Server) Start() {
	http.HandleFunc("/ia", Handler)
	srv := &http.Server{
		Addr:              s.Port,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
