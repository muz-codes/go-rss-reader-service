package main

import (
	routes "go-rss-reader-service/router"
	"go.uber.org/zap"
)

func main() {
	logger := zap.L()
	r := routes.SetupRouter()
	// Listens to Server at 0.0.0.0:8086
	if err := r.Run(":8086"); err != nil {
		logger.Error("Failed to run server", zap.Error(err))
	}
}
