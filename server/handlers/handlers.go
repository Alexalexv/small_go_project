package handlers

import "small_go_project/logger"

type Handlers struct {
	log *logger.Logger
}

func NewHandlers(log *logger.Logger) *Handlers {
	return &Handlers{
		log: log,
	}
}
