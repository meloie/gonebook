package webserver

import (
	"gonebook/internal/services"
)

type WebServer struct {
	service *services.Service
}

func NewWebServer(service *services.Service) *WebServer {
	return &WebServer{service}
}
