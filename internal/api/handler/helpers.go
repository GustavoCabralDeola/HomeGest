package handler

import (
	"encoding/json"
	"net/http"
)

// respondJSON escreve a resposta HTTP com status e corpo JSON.
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// respondError escreve uma resposta de erro em JSON.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// readJSON decodifica o corpo da requisição JSON.
func readJSON(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(dst)
}
