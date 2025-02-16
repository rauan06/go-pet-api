package main

import (
	"assignment2/internal/handlers"
	"assignment2/pkg/lib/logger"
	"log/slog"
	"os"
)

func main() {
	logger := logger.SetupPrettySlog(os.Stdout)
	slog.SetDefault(logger)

	srv := handlers.NewAPIServer(":8080")

	srv.Run()
}
