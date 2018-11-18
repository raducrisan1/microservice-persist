package main

import (
	"context"
	"net"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/raducrisan1/microservice-persist/stockreport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	MongoDB *mongo.Client
}

func newGrpcDataServer(mng *mongo.Client) {
	lis, err := net.Listen("tcp", ":3050")
	failOnError(err, "Could not start gRPC server")
	s := grpc.NewServer()
	stockreport.RegisterStockReportDataServiceServer(
		s,
		&grpcServer{
			MongoDB: mng})

	//this is used to allow API inspection via grpc_cli tool
	reflection.Register(s)
	err = s.Serve(lis)
	failOnError(err, "Could not connect do mongodb server")
}

func (s *grpcServer) GetStockReportData(ctx context.Context, req *stockreport.StockReportRequest) (*stockreport.StockReportResponse, error) {
	coll := s.MongoDB.Database("trading").Collection("ratings")

	cursor, err := coll.Find(ctx,
		bson.D{
			{Key: "stockrating.stockname", Value: "NVDA"},
		})

	if err != nil {
		return nil, err
	}

	sd := make([]*stockreport.StockReportItem, 0)

	for cursor.Next(ctx) {
		doc := bsonx.Doc{}

		err = cursor.Decode(&doc)
		if err != nil {
			continue
		}

		item := new(stockreport.StockReportItem)

		elem := doc.Lookup("stockrating").Document()
		item.Stockname = elem.Lookup("stockname").StringValue()
		item.Rating = elem.Lookup("rating").Int32()
		sd = append(sd, item)
	}

	rsp := stockreport.StockReportResponse{
		StockData: sd}

	return &rsp, nil
}
