package models

type Room struct {
	ID            uint `gorm:"primary"`
	RoomName      string
	Content       string
	Clients       []*Client
	EventListener chan Event
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
	r.Clients = append(r.Clients, client)
}

func (r *Room) Connect(client *Client) error {
	var data map[string]interface{}
	var err error
	for {
		err = client.WS.ReadJSON(data)
		if err != nil {
			return err
		}
		r.EventListener <- data
	}
}
