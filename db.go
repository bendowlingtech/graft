package graft

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*sqlx.DB
}

func New(dsn string, dbType string) {
	switch dbType {
	case "postgres":
		fmt.Println("postgres selected")
	case "mysql":

	case "sqlite":
	case "mongo":
	}

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Couldnt connect to DB")
	}

	return
}
