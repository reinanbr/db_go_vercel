package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/tbxark/g4vercel"
)

type ResponseIndex struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}


func Index(w http.ResponseWriter, r *http.Request) {
	Server := New()
	Server.GET("/",func(context *Content){

	response := ResponseIndex{
		Message:   "Estamos online",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	
	context.JSON(200,H{"message":response.Message,"date":response.Timestamp})

	//fmt.Print("server ok\n")
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(response)
})

