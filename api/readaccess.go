package handler



import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"api_save_data/api/db"
	"api_save_data/api/models"
)


func Read(w http.ResponseWriter,r*http.Request){
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

