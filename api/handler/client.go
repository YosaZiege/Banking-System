package handler

import (
	"github.com/YosaZiege/banksystem/config"
	db "github.com/YosaZiege/banksystem/internal/db/sqlc"
	"github.com/YosaZiege/banksystem/internal/util"
)

type ClientHandler struct {
	queries *db.Queries
	config  util.Config
}
