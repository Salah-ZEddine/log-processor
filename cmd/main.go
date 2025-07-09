package main

import (
    "log-processor/internal/processor"
	"log-processor/internal/db"
	"context"
	"os"
	"os/signal"
	"syscall"
	"log"

    "github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: Could not load .env file: %v", err)
    }

    autoMigrate := true
    if os.Getenv("DB_AUTO_MIGRATE") == "false" {
        autoMigrate = false
    }

	//Init DB
	db.Init(autoMigrate)
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Handle OS signals for graceful shutdown
    go func() {
        sigchan := make(chan os.Signal, 1)
        signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
        <-sigchan
        cancel()
    }()

    processor.StartConsumer(ctx)
}