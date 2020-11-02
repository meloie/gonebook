package webserver

import (
	"github.com/meloie/gonebook/internal/services"
)

type WebServer struct {
	service *services.Service
}

func NewWebServer(service *services.Service) *WebServer {
	return &WebServer{service}
}
