package main

import (
	"fmt"
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
	clients ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{clients: make(ClientList)}
}

func (m *Manager) chat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("user_id")
	cliendId := ps.ByName("client_id")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	// m.clients = append(m.clients, Client{Connection: conn, ClinetID: cliendId})
	m.AddClient(&Client{Connection: conn, ClinetID: cliendId, Manager: m})
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		for client := range m.clients {
			if client.ClinetID == userId {
				client.Connection.WriteMessage(messageType, p)
			}
		}
	}
}

func (m *Manager) AddClient(c *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[c] = true
}

func (m *Manager) RemoveClient(c *Client) error {
	m.Lock()
	defer m.Unlock()
	ok := m.clients[c]
	if !ok {
		return fmt.Errorf("Client Not Found")
	}
	c.Connection.Close()
	delete(m.clients, c)
	return nil
}
