package main

import (
	"database/sql"
	"screening/config"
	db "screening/db/sqlc"
	"screening/internal/app"

	_ "github.com/go-sql-driver/mysql"

	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	// load config
	zap.S().Info("loading config...")
	config, err := config.LoadConfig(".")
	if err != nil {
		zap.S().Fatalf("cannot load config: %v", err)
		return
	}

	// initialize the database
	zap.S().Info("connecting to db...")
	conn, err := sql.Open("mysql", config.DataBaseURL)
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
	server.Start(config.ServerAddress)
}
