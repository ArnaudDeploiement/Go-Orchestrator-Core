package server

import (
	"encoding/json"
	"log"
	"net/http"
	"orchestrator/LLM"
	"orchestrator/tools"
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
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	if ct := r.Header.Get("Content-Type"); ct != "" && !strings.Contains(ct, "application/json") {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var message LLM.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			http.Error(w, "payload too large", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	client := LLM.NewOllamaClient()
	resp, err := client.Chat(LLM.Prompt(message.Content))
	if err != nil {
		log.Printf("LLM call failed: %v", err)
		http.Error(w, "upstream LLM error", http.StatusBadGateway)
		return
	}

	content := strings.TrimSpace(resp.Message.Content)

	//Change ici pour le MAS. Récupérer le contenu de la réponse de l'API pour l'envoyer dans un canal.
	// Les agents écouteront le canal pour récupérer le contenu de la réponse de l'API.

	var tc tools.Tool
	var apiResp LLM.APIResponse

	router([]byte(content), &tc, &apiResp, w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResp)

	log.Println("🧪 LLM RAW:", content)
	log.Printf("🧪 PARSED TOOL: %+v\n", tc)

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
