package api

import (
	"github.com/gin-gonic/gin"

	"github.com/YosaZiege/banksystem/api/handler"
	"github.com/YosaZiege/banksystem/config"
	db "github.com/YosaZiege/banksystem/internal/db/sqlc"
)

type Server struct {
	queries *db.Queries
	config  config.Config
	router  *gin.Engine
}

func NewServer(config config.Config, queries *db.Queries) (*Server, error) {
	server := &Server{
		config:  config,
		queries: queries,
	}
	router := gin.Default()

	clientHandler := handler.NewClientHandler(queries, config)

	router.POST("/clients", clientHandler.CreateClient)

	server.router = router
	return server, nil
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
