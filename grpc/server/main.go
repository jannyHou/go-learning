package main

import (
	"context"
	pb "examples/go-learning/grpc/order_proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(c context.Context, order *pb.Order) (*pb.Order, error) {
	return order, nil
}

func (s *server) GetOrder(order *pb.Order, srv grpc.ServerStreamingServer[pb.Order]) error {
	srv.Send(order)
	return nil
}
func main() {
	lis, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatal("Cannot start server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &server{})

	log.Fatal(grpcServer.Serve(lis))
}
