package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

///////// First Code Snippet Commented out is the an Old way of Handling Transactions

// Will provide all functions to execute Db queries And Transactions
// type Store struct {
// 	connPool *pgx.Conn
// 	*Queries
// }

// thats an Old methode now we work with an Interface

// type Store interface {

// }

// Since Queries is Limited for executing a Transaction we do a Compostion aka *Queries
// func NewStore(db *pgx.Conn) *Store {
// 	return &Store{
// 		connPool: db,
// 		Queries:  New(db), // New is generated from sqlc it returns queries
// 	}
// }

// execTx executes a function Within a Database Transaction
// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.Begin(ctx)
// 	if err != nil {
// 		return err
// 	}

// q := New(tx)
// err
// }

type Store interface{}

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx) // use the transaction to execute queries
	err = fn(q)  // execute the function that uses q
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
