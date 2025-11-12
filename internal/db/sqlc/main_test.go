package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/YosaZiege/banksystem/internal/util"
)

var (
	testQueries *Queries
	testDB      *pgxpool.Pool
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Failed to Connect to the Database  ", err)
	}

	defer testDB.Close()

	testQueries = New(testDB)
	os.Exit(m.Run())
}
