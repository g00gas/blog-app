package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func InitDB() *pgxpool.Pool {
	//TODO: pass url as env.
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/blog_db")
	if err != nil {
		fmt.Println("Couldn't connect to postgres.")
		log.Fatal(err)
	}
	return conn
}
