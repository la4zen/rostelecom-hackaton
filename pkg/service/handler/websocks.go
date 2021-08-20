package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/la4zen/rostelecom-hackaton/pkg/models"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func (h *Handler) WSHandler(c echo.Context, roomid uint, user *models.User) {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return
	}
	defer ws.Close()
	client := &models.Client{
		WS:   ws,
		User: user,
	}
	h.Rooms[roomid].Connect(client)
}
