package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func createRabbitMqReader() (<-chan amqp.Delivery, *amqp.Connection, *amqp.Channel, error) {
	rabbitMqAddr := os.Getenv("RABBITMQ_ADDR")
	conn, err := amqp.Dial(rabbitMqAddr)
	failOnError(err, "could not open the connection to rabbitmq exchange")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Could not open a RabbitMQ channel for read: %v", err)
		conn.Close()
		panic("Could not open a RabbitMQ channel for read")
	}

	q, err := ch.QueueDeclare(
		"StockRatingData",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalf("Could not declare a RabbitMQ queue to read from.  %v", err)
		conn.Close()
		ch.Close()
		panic("Could not declare a RabbitMQ queue to read from")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to register a consumer")
	return msgs, conn, ch, err
}
