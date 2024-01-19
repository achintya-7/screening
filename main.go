package main

import (
	"database/sql"
	db "screening/db/sqlc"
	"screening/internal/app"

	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	// initialize the database
	zap.S().Info("connecting to db...")
	conn, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		zap.S().Fatalf("cannot connect to db: %v", err)
		return
	}

	// initialize the store
	zap.S().Info("initializing store...")
	store := db.NewStore(conn)

	// initialize the server
	zap.S().Info("initializing server...")
	server := app.NewServer(store)
	server.Start()
}
