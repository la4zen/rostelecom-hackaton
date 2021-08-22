package models

import (
	"encoding/json"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Room struct {
	ID            uint `gorm:"primary" json:"id,omitempty"`
	RoomName      string
	Content       string
	DB            *gorm.DB               `gorm:"-"`
	Contents      map[string]interface{} `gorm:"-"`
	Clients       map[uint]*Client       `gorm:"-"`
	EventListener chan Event             `gorm:"-"`
}

func (r *Room) WriteContent() {
	content, err := json.Marshal(r.Contents)
	if err != nil {
		log.Error(err)
		return
	}
	r.DB.Model(r).Update("content", string(content))
}

func (r *Room) Listen() {
	for {
		select {
		case event := <-r.EventListener:
			for _, client := range r.Clients {
				if client.User.ID != event.User.ID {
					client.WS.WriteJSON(event.Event)
				}
			}
		}
	}
}

func (r *Room) AddClient(client *Client) {
	go func() {
		r.EventListener <- Event{
			*client.User,
			map[string]interface{}{
				"event_type": "user_joined",
			},
		}
	}()
	r.Clients[client.User.ID] = client
}

func (r *Room) RemoveClient(client *Client) {
	if r.Clients[client.User.ID] == nil {
		return
	}
	delete(r.Clients, client.User.ID)
	r.EventListener <- Event{
		*client.User,
		map[string]interface{}{
			"event_type": "user_leaved",
			"user_id":    client.User.ID,
		},
	}
}

func (r *Room) Connect(client *Client) {
	var data map[string]interface{}
	var err error
	err = client.WS.WriteJSON(r.Contents)
	if err != nil {
		return
	}
	r.AddClient(client)
	for {
		err = client.WS.ReadJSON(&data)
		if err != nil {
			r.RemoveClient(client)
			return
		}
		if data["event_type"].(*string) != nil {
			switch data["event_type"].(string) {
			case "add_component":

			}
		}
		r.EventListener <- Event{
			*client.User,
			data,
		}
	}
}
