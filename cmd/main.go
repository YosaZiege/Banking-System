package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/YosaZiege/banksystem/api"
	"github.com/YosaZiege/banksystem/config"
	db "github.com/YosaZiege/banksystem/internal/db/sqlc"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := pgxpool.New(context.Background(), cfg.DBSource)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer conn.Close()

	queries := db.New(conn)

	server, err := api.NewServer(cfg, queries)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	if err := server.Start(cfg.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
