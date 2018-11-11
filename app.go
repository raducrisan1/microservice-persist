package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/raducrisan1/microservice-persist/stockinfo"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(msg)
	}
}

//AnnotatedRating type wraps the StockRating and appends an object id so
//that it can be used in mongo db
type AnnotatedRating struct {
	stockinfo.StockRating
	ID objectid.ObjectID
}

func main() {

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

	mng, err := mongo.NewClient("mongodb://localhost:27017")
	failOnError(err, "Could not create a mongodb client")
	defer mng.Disconnect(nil)

	err = mng.Connect(context.Background())
	failOnError(err, "Could not connect do mongodb server")
	coll := mng.Database("trading").Collection("ratings")

	forever := make(chan bool)
	limiter := time.Tick(time.Second / 70)

	go func() {
		for d := range msgs {
			<-limiter
			//here, write the message d into mongodb
			protomsg := stockinfo.StockRating{}
			if err = proto.Unmarshal(d.Body, &protomsg); err != nil {
				log.Printf("Could not deserialize protobuf message %v", d.MessageId)
				continue
			}
			annotatesMsg := AnnotatedRating{StockRating: protomsg, ID: objectid.New()}
			coll.InsertOne(nil, annotatesMsg)
		}
	}()

	<-forever
}
