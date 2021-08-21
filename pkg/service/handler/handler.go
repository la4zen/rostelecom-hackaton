package handler

import (
	"encoding/json"
	"fmt"

	"github.com/la4zen/rostelecom-hackaton/pkg/models"
	"github.com/la4zen/rostelecom-hackaton/pkg/store"
)

type Handler struct {
	Rooms map[uint]*models.Room
	Store *store.Store
}

func New(store *store.Store) *Handler {
	_rooms := []models.Room{}
	rooms := map[uint]*models.Room{}
	res := store.DB.Find(&_rooms)
	fmt.Print(res)
	for _, room := range _rooms {
		room.Clients = map[uint]*models.Client{}
		room.Contents = map[string]interface{}{}
		_ = json.Unmarshal([]byte(room.Content), &room.Contents)
		room.EventListener = make(chan models.Event)
		go room.Listen()
		rooms[room.ID] = &room
	}
	return &Handler{
		Rooms: rooms,
		Store: store,
	}
}
