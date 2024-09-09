package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  ReadBufferSize,
	WriteBufferSize: WriteBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Manager struct {
	clients []Client
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) greeter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("user_id")
	cliendId := ps.ByName("client_id")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
	}
	m.clients = append(m.clients, Client{Connection: conn, ClinetID: cliendId})
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		for _, client := range m.clients {
			if client.ClinetID == userId {
				client.Connection.WriteMessage(messageType, p)
			}
		}
	}
}
