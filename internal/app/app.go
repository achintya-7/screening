package app

import (
	db "screening/db/sqlc"
	v1 "screening/internal/controllers/v1"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.Store
}

func NewServer(store db.Store) *Server {
	server := &Server{
		router: gin.Default(),
	}

	server.registerRoutes()

	return server
}

func (s *Server) registerRoutes() {
	baseRouter := s.router.Group("service-name")

	v1Router := v1.NewRouter(s.store)
	v1Router.RegisterRoutes(baseRouter)
}

func (s *Server) Start() error {
	return s.router.Run()
}
