package handler

import (
	"github.com/la4zen/rostelecom-hackaton/pkg/models"
	"github.com/la4zen/rostelecom-hackaton/pkg/store"
)

type Handler struct {
	Rooms map[uint]*models.Room
	Store *store.Store
}

func New(store *store.Store) *Handler {
	_rooms := []*models.Room{}
	rooms := map[uint]*models.Room{}
	store.DB.Find(&_rooms)
	for _, room := range rooms {
		rooms[room.ID] = room
	}
	return &Handler{
		Rooms: rooms,
		Store: store,
	}
}
