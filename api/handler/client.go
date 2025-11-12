package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/YosaZiege/banksystem/config"
	db "github.com/YosaZiege/banksystem/internal/db/sqlc"
)

type ClientHandler struct {
	queries *db.Queries
	config  config.Config
}

type createClientRequest struct {
	Username     string      `json:"username"`
	Balance      int64       `json:"balance"`
	Currency     interface{} `json:"currency"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password_hash"`
	Provider     string      `json:"provider"`
}

func NewClientHandler(queries *db.Queries, config config.Config) *ClientHandler {
	return &ClientHandler{queries: queries, config: config}
}

func (h *ClientHandler) CreateClient(ctx *gin.Context) {
	var req createClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		arg := db.CreateClientParams{
			Username:     req.Username,
			Balance:      req.Balance,
			Currency:     req.Currency,
			Email:        req.Email,
			PasswordHash: req.PasswordHash,
			Provider:     req.PasswordHash,
		}

		client, err := h.queries.CreateClient(context.Background(), arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, client)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "client Created"})
}

func (h *ClientHandler) ListClients(ctx *gin.Context) {
	var req db.ListClientsParams
	type ListClientsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

		arg := db.CreateClientParams{
			Username:     req.Username,
			Balance:      req.Balance,
			Currency:     req.Currency,
			Email:        req.Email,
			PasswordHash: req.PasswordHash,
			Provider:     req.PasswordHash,
		}

		client, err := h.queries.CreateClient(context.Background(), arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, client)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "client Created"})
}
