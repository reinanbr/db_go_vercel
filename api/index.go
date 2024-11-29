package handler

import (
	"fmt"
	"net/http"
	"time"
	. "github.com/tbxark/g4vercel"
)

type ResponseIndex struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
};


func Index(w http.ResponseWriter, r *http.Request) {
	server := New();
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
	server.GET("/",func(context *Context) {

//	response := ResponseIndex{
//		Message:   "Estamos online",
//		Timestamp: time.Now().Format(time.RFC3339),
//	};
	
	context.JSON(200,
	H{"message":"estamos fudendo","date":time.Now().Format(time.RFC3339)});
})	
	server.Handle(w,r);
}
