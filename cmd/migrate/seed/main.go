package main

import (
	"log"

	"github.com/ivanpaghubasan/go-social/internal/db"
	"github.com/ivanpaghubasan/go-social/internal/env"
	"github.com/ivanpaghubasan/go-social/internal/store"
)

func main() {
	addr := env.GetString("DB_URL", "postgresql://admin:adminpassword@localhost/social?sslmode=disable")

	conn, err := db.New(addr, 30, 30, "15m")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
