package service

import (
	"github.com/la4zen/rostelecom-hackaton/pkg/service/handler"
	"github.com/la4zen/rostelecom-hackaton/pkg/store"
)

type Service struct {
	Store   *store.Store
	Handler *handler.Handler
}

func New(store *store.Store) *Service {
	return &Service{
		Store:   store,
		Handler: handler.New(store),
	}
}
