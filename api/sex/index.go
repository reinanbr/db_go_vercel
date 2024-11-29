package handler
import (
	"encoding/json"
	"net/http"
	"time"
//        "fmt"
//        "log"
//        "api_save_data/api/db"
//        "api_save_data/api/models"
  )

type Response struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// Handler é a função exportada para o Vercel.
func Sex(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message:   "Estamos online",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
