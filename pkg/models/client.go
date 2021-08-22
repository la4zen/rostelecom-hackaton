package models

import "github.com/gorilla/websocket"

type Client struct {
	WS   *websocket.Conn `json:"-"`
	User *User           `json:"user,omitempty"`
}
