package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"inventory_queue",    // queue
		"inventory_consumer", // consumer
		false,                // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	failOnError(err, "Failed to register a consumer")

	func() {
		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			switch d.RoutingKey {
			case "order.created":
				fmt.Println("Order criada -> Diminuir o estoque")
				handleOrderCreated(d)

			case "order.failed":
				fmt.Println("Order falhou -> Repor o estoque")
				handleOrderFailed(d)
			default:
				log.Printf("Unhandled routing key: %s", d.RoutingKey)
			}

			d.Ack(false)
		}
	}()

}

type Order struct {
	Id       uint64
	Status   string
	Customer string
	Product  uint64
}

func handleOrderCreated(d amqp.Delivery) error {
	var order Order
	err := json.Unmarshal(d.Body, &order)
	if err != nil {
		d.Ack(false)
		return fmt.Errorf("error: %s", err)
	}

	fmt.Println(order)

	return nil
}

func handleOrderFailed(delivery amqp.Delivery) {
	var order Order
	json.Unmarshal(delivery.Body, &order)
	fmt.Println(order)
}
