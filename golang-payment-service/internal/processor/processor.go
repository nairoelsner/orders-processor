package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"

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

func ProcessPayment(msg amqp.Delivery) (ProcessedMessage, error) {
	defer msg.Ack(false)

	log.Printf("Received a message: %s", msg.Body)
	fmt.Println("Estoque confirmado -> Processar pagamento")

	var order Order
	err := json.Unmarshal(msg.Body, &order)
	if err != nil {
		return ProcessedMessage{"", []byte{}}, fmt.Errorf("error: %s", err)
	}

	n := rand.IntN(2)
	if n == 0 {
		order.Status = "Payment accepted"
	} else {
		order.Status = "Payment denied"
	}

	orderBytes, err := json.Marshal(order)
	if err != nil {
		return ProcessedMessage{"", []byte{}}, fmt.Errorf("error: %s", err)
	}

	if n == 0 {
		return ProcessedMessage{"order.payment.accepted", orderBytes}, nil
	}
	return ProcessedMessage{"order.payment.denied", orderBytes}, nil
}
