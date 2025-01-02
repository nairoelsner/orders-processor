package main

import (
	"inventory-service/internal/config"
	"inventory-service/internal/consumer"
	"inventory-service/internal/processor"
	"inventory-service/internal/producer"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := "amqp://guest:guest@localhost:5672/"

	conn, ch := config.ConnectRabbitMQ(rabbitURL)
	defer conn.Close()
	defer ch.Close()

	queueName := "inventory_queue"
	exchangeName := "amq.topic"
	routingKeys := []string{"order.created", "order.payment.denied"}

	config.DeclareQueue(ch, queueName)
	for _, routingKey := range routingKeys {
		config.BindQueue(ch, queueName, routingKey, exchangeName)
	}

	messages := make(chan amqp.Delivery)

	go consumer.ConsumeMessages(ch, "inventory_queue", messages)

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	for msg := range messages {
		processed, err := processor.ProcessMessage(msg)
		if err != nil {
			log.Printf("Error processing message: %s", err)
			continue
		}
		producer.PublishMessage(ch, exchangeName, processed.RoutingKey, processed.Body)
	}
}
