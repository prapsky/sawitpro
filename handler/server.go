package handler

import (
	"github.com/prapsky/sawitpro/service"
)

type Server struct {
	service service.Service
}

func NewServer(service service.Service) *Server {
	return &Server{
		service: service,
	}
}
