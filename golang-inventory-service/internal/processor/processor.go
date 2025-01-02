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

type ProcessedMessage struct {
	RoutingKey string
	Body       []byte
}

func ProcessMessage(msg amqp.Delivery) (ProcessedMessage, error) {
	defer msg.Ack(false)

	log.Printf("Received a message: %s", msg.Body)

	switch msg.RoutingKey {
	case "order.created":
		fmt.Println("Order criada -> Diminuir o estoque")
		body, err := handleOrderCreated(msg.Body)
		if err != nil {
			log.Printf("Error handling order created: %s", err)
			return ProcessedMessage{"", []byte{}}, err
		}

		return ProcessedMessage{"order.inventory.available", body}, nil

	case "order.payment.denied":
		fmt.Println("Order falhou -> Repor o estoque")
		err := handleOrderFailed(msg)
		if err != nil {
			log.Printf("Error handling order failed: %s", err)
			return ProcessedMessage{}, err
		}
		return ProcessedMessage{"", []byte{}}, nil

	default:
		log.Printf("Unhandled routing key: %s", msg.RoutingKey)
		return ProcessedMessage{"", []byte{}}, nil
	}
}

func handleOrderCreated(body []byte) ([]byte, error) {
	var order Order
	err := json.Unmarshal(body, &order)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	// to do
	// check if product is available in inventory

	order.Status = "Inventory Available"

	orderBytes, err := json.Marshal(order)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return orderBytes, nil
}

func handleOrderFailed(delivery amqp.Delivery) error {
	var order Order
	json.Unmarshal(delivery.Body, &order)
	fmt.Println("Repondo o estoque da order:", order.Id)
	return nil
}
