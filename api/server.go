package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	db "github.com/YosaZiege/banksystem/internal/db/sqlc"
	"github.com/YosaZiege/banksystem/internal/util"
)

type Server struct {
	queries *db.Queries
	config  util.Config
	router  *gin.Engine
}

func NewServer(config util.Config, queries *db.Queries) (*Server, error) {
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

