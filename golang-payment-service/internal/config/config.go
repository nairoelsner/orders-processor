package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareQueue(ch *amqp.Channel, name string) {
	_, err := ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
}

func BindQueue(ch *amqp.Channel, queueName, routingKey, exchange string) {
	err := ch.QueueBind(
		queueName,  // queue
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}
}

func ConnectRabbitMQ(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return conn, ch
}
