package main

import (
	"log"
	"small_go_project/config"
	"small_go_project/logger"
	"small_go_project/server"
	"small_go_project/server/handlers"
)

func main() {
	cfg, err := config.NewConfiguration()
	if err != nil {
		log.Fatal("failed to create configuration")
	}

	l := logger.NewLogger(cfg)
	h := handlers.NewHandlers(l)

	s := server.NewServer(cfg, h)
	err = s.Start()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
