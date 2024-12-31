package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	Id       uint64
	Status   string
	Customer string
	Product  uint64
}

func ProcessMessage(msg amqp.Delivery) []byte {
	log.Printf("Received a message: %s", msg.Body)

	switch msg.RoutingKey {
	case "order.created":
		fmt.Println("Order criada -> Diminuir o estoque")
		handleOrderCreated(msg)

	case "order.failed":
		fmt.Println("Order falhou -> Repor o estoque")
		handleOrderFailed(msg)
	default:
		log.Printf("Unhandled routing key: %s", msg.RoutingKey)
	}

	return nil
}

func handleOrderCreated(msg amqp.Delivery) error {
	// Unmarshal the message body into the Order struct
	// Verify if the unmarshal operation was successful
	// Consult database to update the stock
	// Acknowledge the message
	// Return message order.inventory.check or order.failed

	var order Order
	err := json.Unmarshal(msg.Body, &order)
	if err != nil {
		msg.Ack(false)
		return fmt.Errorf("error: %s", err)
	}

	fmt.Println(order)

	// if database is out of service -> msg.nack
	random := rand.Intn(2)
	if random == 0 {
		msg.Nack(false, true)
		return fmt.Errorf("error: database is out of service")
	}

	msg.Ack(false)

	return nil
}

func handleOrderFailed(delivery amqp.Delivery) {
	// Unmarshal the message body into the Order struct
	// Verify if the unmarshal operation was successful
	// Consult database to update the stock
	// Acknowledge the message

	var order Order
	json.Unmarshal(delivery.Body, &order)
	fmt.Println(order)
}
