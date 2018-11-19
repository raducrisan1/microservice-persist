package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/raducrisan1/microservice-persist/stockinfo"
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
	msgs, rabbitConn, rabbitChannel, err := createRabbitMqReader()
	failOnError(err, "Could not initialize rabbitmq reader")
	defer rabbitConn.Close()
	defer rabbitChannel.Close()
	defer rabbitChannel.Close()

	mongoAddr := os.Getenv("MONGO_ADDR")
	mng, err := mongo.NewClient(mongoAddr)
	failOnError(err, "Could not create a mongodb client")
	defer mng.Disconnect(nil)

	err = mng.Connect(context.Background())
	failOnError(err, "Could not connect do mongodb server")
	coll := mng.Database("trading").Collection("ratings")

	fmt.Println("microservice-persist has started")

	stop := make(chan bool)
	limiter := time.Tick(time.Second / 70)

	//this goroutine reads messages from rabbitmq and writes them to mongodb
	go func() {
		for msg := range msgs {
			<-limiter
			//here, annotate the message with ObjectID and write it into mongodb
			protomsg := stockinfo.StockRating{}
			if err = proto.Unmarshal(msg.Body, &protomsg); err != nil {
				log.Printf("Could not deserialize protobuf message %v", msg.MessageId)
				continue
			}
			annotatedMsg := AnnotatedRating{StockRating: protomsg, ID: objectid.New()}
			coll.InsertOne(nil, annotatedMsg)
			msg.Ack(false)
		}
	}()

	//this goroutine listens to gRPC calls, in our case from reports microservice
	go func() {
		newGrpcDataServer(mng)
	}()

	<-stop
}
