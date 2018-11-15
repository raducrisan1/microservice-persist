package main

import "github.com/streadway/amqp"

func setupQueue() <-chan amqp.Delivery {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "could not open the connection to rabbitmq exchange")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"StockRatingData",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to register a consumer")
	return msgs
}
