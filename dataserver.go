package main

import (
	"context"
	"log"
	"net"

	"github.com/raducrisan1/microservice-persist/stockreport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct{}

func grpcdataserver() {
	lis, err := net.Listen("tcp", ":3040")
	failOnError(err, "Could not start gRPC server")
	s := grpc.NewServer()
	stockreport.RegisterStockReportDataServiceServer(s, &grpcServer{})
	//this is used to allow API inspection via grpc_cli tool
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *grpcServer) GetStockReportData(ctx context.Context, req *stockreport.StockReportRequest) (*stockreport.StockReportResponse, error) {
	sd := make([]*stockreport.StockReportItem, 1)
	sd[0] = new(stockreport.StockReportItem)
	sd[0].Rating = 2
	sd[0].Stockname = "NVDA"
	rsp := stockreport.StockReportResponse{
		StockData: sd}
	return &rsp, nil
}
