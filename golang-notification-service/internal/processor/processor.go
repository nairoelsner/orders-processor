package processor

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	Id       uint64
	Status   string
	Customer string
	Product  uint64
}

func ProcessMessage(msg amqp.Delivery) error {
	defer msg.Ack(false)

	log.Printf("Received a message: %s", msg.Body)

	var order Order
	err := json.Unmarshal(msg.Body, &order)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	switch msg.RoutingKey {
	case "order.inventory.unavailable":
		log.Printf("Sending Order inventory unavailable email to: %s", order.Customer)
	case "order.payment.accepted":
		log.Printf("Order payment accepted to: %s", order.Customer)
	case "order.payment.denied":
		log.Printf("Order payment denied to: %s", order.Customer)
	default:
		log.Printf("Unhandled routing key: %s", msg.RoutingKey)
	}

	return nil
}
