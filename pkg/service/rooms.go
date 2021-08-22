package service

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/la4zen/rostelecom-hackaton/pkg/models"
	"github.com/labstack/echo"
)

func (s *Service) Connect(c echo.Context) error {
	user := &models.User{}
	user.ID = c.Get("user").(*jwt.Token).Claims.(*models.Claims).ID
	_id := c.Param("id")
	if _id == "" {
		return c.String(400, "id required")
	}
	id, err := strconv.Atoi(_id)
	if err != nil {
		return c.String(500, err.Error())
	}
	if s.Handler.Rooms[uint(id)] == nil {
		return c.String(404, "room not found")
	}
	if len(s.Handler.Rooms[uint(id)].Clients) >= 4 {
		return c.String(403, "room is full")
	}
	go s.Handler.WSHandler(c, uint(id), user)
	return nil
}

func (s *Service) GetRooms(c echo.Context) error {
	rooms := []map[string]interface{}{}
	for _, room := range s.Handler.Rooms {
		rooms = append(rooms, map[string]interface{}{
			"room_id":   room.ID,
			"room_name": room.RoomName,
		})
	}
	return c.JSON(200, echo.Map{
		"rooms": rooms,
	})
}
