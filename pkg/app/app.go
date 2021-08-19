package app

import (
	"github.com/la4zen/rostelecom-hackaton/pkg/service"
	"github.com/la4zen/rostelecom-hackaton/pkg/store"
	"github.com/labstack/echo"
)

func Run() {
	store := store.New()
	service := service.New(store)
	e := echo.New()
	e.POST("/tables/upload", service.Upload)
	e.Logger.Fatal(e.Start(":8080"))
}
