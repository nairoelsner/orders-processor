package main

import (
	"log"
	"payment-service/internal/config"
	"payment-service/internal/consumer"
	"payment-service/internal/processor"
	"payment-service/internal/producer"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := "amqp://guest:guest@localhost:5672/"

	conn, ch := config.ConnectRabbitMQ(rabbitURL)
	defer conn.Close()
	defer ch.Close()

	queueName := "payment_queue"
	exchangeName := "amq.topic"
	routingKeys := []string{"order.inventory.available"}

	config.DeclareQueue(ch, queueName)
	for _, routingKey := range routingKeys {
		config.BindQueue(ch, queueName, routingKey, exchangeName)
	}

	messages := make(chan amqp.Delivery)

	go consumer.ConsumeMessages(ch, "payment_queue", messages)

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	for msg := range messages {
		processed, err := processor.ProcessPayment(msg)
		if err != nil {
			log.Printf("Error processing message: %s", err)
			continue
		}
		producer.PublishMessage(ch, exchangeName, processed.RoutingKey, processed.Body)
	}
}
