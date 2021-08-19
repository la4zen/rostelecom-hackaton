package app

import (
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
	e.Use(middleware.CORS())
	e.PUT("/register", service.Register)
	e.POST("/login", service.Login)
	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup:   "cookie:token, header:Authorization, query:token",
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    util.Secret,
	}))
	r.GET("/tables/:id", service.Connect)
	e.Logger.Fatal(e.Start(":8080"))
}
