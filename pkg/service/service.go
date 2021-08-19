package service

import (
	"github.com/la4zen/rostelecom-hackaton/pkg/store"
	"github.com/labstack/echo"
)

type Service struct {
	Store *store.Store
}

func New(store *store.Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) Upload(c echo.Context) error {
	return c.NoContent(200)
}
