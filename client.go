package main

import "github.com/gorilla/websocket"

type Client struct {
	Connection *websocket.Conn
	ClinetID   string
}
