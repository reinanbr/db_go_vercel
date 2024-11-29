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
};


func Handler(w http.ResponseWriter, r *http.Request) {
	Server := New();
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))
	Server.GET("/",func(context *Content){

//	response := ResponseIndex{
//		Message:   "Estamos online",
//		Timestamp: time.Now().Format(time.RFC3339),
//	};
	
	context.JSON(200,H{"message":"estamos fudendo","date":time.Now().format.RFC3339});
	
	server.Handle(w,r);
});
