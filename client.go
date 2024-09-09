package main

import "github.com/gorilla/websocket"

type ClientList map[*Client]bool

type Client struct {
	Connection *websocket.Conn
	ClinetID   string
	Manager    *Manager
}
