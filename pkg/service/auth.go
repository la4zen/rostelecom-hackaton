package service

import (
	"github.com/la4zen/rostelecom-hackaton/pkg/models"
	"github.com/la4zen/rostelecom-hackaton/pkg/util"
	"github.com/labstack/echo"
)

func (s *Service) Register(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.Validate(); err != nil {
		return c.String(400, err.Error())
	}
	if err := s.Store.CreateUser(user); err != nil {
		return c.String(401, err.Error())
	}
	token, err := util.GenerateToken(user)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, echo.Map{
		"token": token,
	})
}

func (s *Service) Login(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.Validate(); err != nil {
		return c.String(400, err.Error())
	}
	if err := s.Store.GetUser(user); err != nil {
		return c.String(404, err.Error())
	}
	token, err := util.GenerateToken(user)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, echo.Map{
		"token": token,
	})
}
