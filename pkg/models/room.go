package models

type Room struct {
	ID            uint `gorm:"primary" json:"id,omitempty"`
	RoomName      string
	Content       string
	Contents      map[string]interface{} `gorm:"-"`
	Clients       map[uint]*Client       `gorm:"-"`
	EventListener chan Event             `gorm:"-"`
}

func (r *Room) Listen() {
	for {
		select {
		case event := <-r.EventListener:
			for _, client := range r.Clients {
				client.WS.WriteJSON(event)
			}
		}
	}
}

func (r *Room) SendEvent(event Event) {
	r.EventListener <- event
}

func (r *Room) AddClient(client *Client) {
	r.Clients[client.User.ID] = client
}

func (r *Room) RemoveClient(client *Client) {
	if r.Clients[client.User.ID] == nil {
		return
	}
	delete(r.Clients, client.User.ID)
	r.EventListener <- map[string]interface{}{
		"event_type": "leave",
		"user_id":    client.User.ID,
	}
}

func (r *Room) Connect(client *Client) {
	r.AddClient(client)
	var data map[string]interface{}
	var err error
	for {
		err = client.WS.ReadJSON(&data)
		if err != nil {
			r.RemoveClient(client)
			return
		}
		r.EventListener <- data
	}
}
