package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseIndex struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}


func Today(w http.ResponseWriter, r *http.Request) {
	response := ResponseIndex{
		Message:   "Estamos online",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	fmt.Print("server ok\n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

