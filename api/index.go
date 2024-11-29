package handler
import (
	"encoding/json"
	"net/http"
	"time"
        "fmt"
        "log"
        "api_save_data/api/db"
        "api_save_data/api/models"
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






func ReadAccess(w http.ResponseWriter,r*http.Request){
        pool :=psql_vercel.ConnectDB()
        defer pool.Close()
        infoAccess,err := access_site_model.ReadAccessSites(pool)
        if err == nil{
                w.Header().Set("Content-Type", "application/json")
		if errJson := json.NewEncoder(w).Encode(infoAccess); errJson != nil {
                        http.Error(w, "Erro ao gerar o JSON", http.StatusInternalServerError)
                        log.Printf("Erro ao codificar JSON: %v", errJson)
        }
        }else{
                fmt.Fprintf(w,"error: %v\n",err)
        }
}

