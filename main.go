package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	m := NewManager()
	router := httprouter.New()

	// client_id is for the one who is joining
	// user_id is to send message to perticular user
	router.GET("/chat/:client_id/:user_id", m.greeter)

	http.ListenAndServe(":8080", router)
}
