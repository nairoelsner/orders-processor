package producer

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMessage(ch *amqp.Channel, exchangeName, routingKey string, body []byte) {
	err := ch.Publish(
		exchangeName,
		routingKey,
		false, // Mandatory
		false, // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	log.Printf("Message published: %s", string(body))
}
