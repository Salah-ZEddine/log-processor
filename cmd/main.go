package main

import (
    "log-processor/internal/processor"
	"context"
	"os"
	"os/signal"
	"syscall"

    
)

func main() {
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