package main

import (
	"go-rss-reader-service/db"
	routes "go-rss-reader-service/router"
	"go.uber.org/zap"
)

func main() {
	logger := zap.L()
	var err error
	db.DbConnection, err = db.OpenDbConnection()
	if err != nil {
		logger.Panic("error while creating db connection")
		return
	}
	r := routes.SetupRouter()
	// Listens to Server at 0.0.0.0:8086
	if err := r.Run(":8086"); err != nil {
		logger.Error("Failed to run server", zap.Error(err))
	}
}
