package app

import (
	"github.com/la4zen/rostelecom-hackaton/pkg/models"
	"github.com/la4zen/rostelecom-hackaton/pkg/service"
	"github.com/la4zen/rostelecom-hackaton/pkg/store"
	"github.com/la4zen/rostelecom-hackaton/pkg/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	store := store.New()
	service := service.New(store)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Static("/", "/static")
	e.PUT("/register", service.Register)
	e.POST("/login", service.Login)
	e.POST("/upload", service.Upload)
	e.GET("/load", service.Load)
	r := e.Group("/api")
	{
		r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			TokenLookup: "query:token",
			SigningKey:  util.Secret,
			Claims:      &models.Claims{},
		}))
		r.GET("/rooms/:id", service.Connect)
	}
	e.Logger.Fatal(e.Start(":8080"))
}
