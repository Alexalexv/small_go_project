package server

import (
	"small_go_project/config"
	"small_go_project/server/handlers"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e    *echo.Echo
	port string
}

func NewServer(cfg *config.Configuration, handlers *handlers.Handlers) *Server {
	e := echo.New()

	e.POST("/user", handlers.CreateUser)
	e.GET("/user/:id", handlers.GetUser)
	e.DELETE("/user/:id", handlers.DeleteUser)
	e.PUT("/user/:id", handlers.UpdateUser)

	return &Server{
		e:    e,
		port: cfg.Port,
	}
}

func (s Server) Start() error {
	return s.e.Start(s.port)
}
