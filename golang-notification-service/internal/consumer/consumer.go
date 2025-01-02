package consumer

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessages(ch *amqp.Channel, queueName string, messages chan<- amqp.Delivery) {
	msgs, err := ch.Consume(
		queueName,
		"notification-consumer",
		false, // Auto-acknowledge
		false, // Not exclusive
		false, // No-local
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		messages <- msg
	}
}
