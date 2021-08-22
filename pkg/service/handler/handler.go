package handler

import (
	"encoding/json"

	. "github.com/la4zen/rostelecom-hackaton/pkg/models"
	"github.com/la4zen/rostelecom-hackaton/pkg/store"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	Rooms map[uint]*Room
	Store *store.Store
}

func New(store *store.Store) *Handler {
	_rooms := []Room{}
	rooms := map[uint]*Room{}
	res := store.DB.Find(&_rooms)
	if res.Error != nil {
		log.Error(res.Error)
	}
	for _, _room := range _rooms {
		room := _room
		room.Clients = map[uint]*Client{}
		room.Contents = map[string]interface{}{}
		_ = json.Unmarshal([]byte(_room.Content), &room.Contents)
		room.EventListener = make(chan Event)
		room.DB = store.DB
		go room.Listen()
		rooms[room.ID] = &room
	}
	return &Handler{
		Rooms: rooms,
		Store: store,
	}
}
