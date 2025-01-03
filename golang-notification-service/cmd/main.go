package main

import (
	"log"
	"notification-service/internal/config"
	"notification-service/internal/consumer"
	"notification-service/internal/processor"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/")

	conn, ch := config.ConnectRabbitMQ(rabbitURL)
	defer conn.Close()
	defer ch.Close()

	queueName := "notification_queue"
	exchangeName := "amq.topic"
	routingKeys := []string{"order.inventory.unavailable", "order.payment.*"}

	config.DeclareQueue(ch, queueName)
	for _, routingKey := range routingKeys {
		config.BindQueue(ch, queueName, routingKey, exchangeName)
	}

	messages := make(chan amqp.Delivery)

	go consumer.ConsumeMessages(ch, "notification_queue", messages)

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	for msg := range messages {
		err := processor.ProcessMessage(msg)
		if err != nil {
			log.Printf("Error processing message: %s", err)
			continue
		}
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
