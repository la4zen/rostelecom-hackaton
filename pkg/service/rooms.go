package service

import (
	"strconv"

	"github.com/golang-jwt/jwt"
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
	return s.Handler.WSHandler(c, uint(id), user)
}
