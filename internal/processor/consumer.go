package processor

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
	"github.com/joho/godotenv"
)

func StartConsumer(ctx context.Context) {

	// Load environment variables
	err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: Could not load .env file: %v", err)
    }

	//fetch URL from .env if no url specified use guest as default
	rabbitUrl := os.Getenv("RABBITMQ_URL")
	if rabbitUrl == "" {
		if _, err := os.Stat(".env"); err == nil {
			// .env exists but RABBITMQ_URL missing? That's misconfiguration!
			log.Fatal("RABBITMQ_URL missing in .env file")
		}
		// PRODUCTION: Require explicit variable in real environments
		log.Fatal("RABBITMQ_URL environment variable not set")
	}

	//open connection to rabbitmq message broker with retry logic
	var conn *amqp.Connection
	err = nil
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(rabbitUrl)
		if err == nil {
			break
		}

		log.Printf("Failed to connect to RabbitMQ (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ after %v attempts: %v", maxRetries, err)
	}
	defer conn.Close()

	//open a channel to consume messages through rabbitmq
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	//declare a queue to consume messages
	queue, err := channel.QueueDeclare(
		"logs_queue", // name
		true,         // durable
		false,        // autoDelete
		false,        // exclusive
		false,        // noWait
		nil,          // arguments
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	//consume messages from the queue
	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Consumer started Successfully")

	// Start consumer
	go func() {
		for msg := range msgs {
			// Process message here
			log.Printf("Received message: %s", msg.Body)
			
			// Acknowledge message after successful processing
			if err := msg.Ack(false); err != nil {
				log.Printf("Error acknowledging message: %v", err)
			}
		}
		log.Println("Consumer stopped")
	}()

	// Wait for context cancellation
	<-ctx.Done()
	log.Println("Shutting down consumer...")


}	