package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/vyctor/go-expert-aulas/11-EVENTOS/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "orders")
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
