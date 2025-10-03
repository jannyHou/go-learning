package main

import (
	"context"
	pb "examples/go-learning/grpc/order_proto"
	"fmt"
	"io"
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

func (s *server) Communicate(srv grpc.BidiStreamingServer[pb.Orders, pb.Orders]) error {
	var result []*pb.Orders
	var count int

	for {
		orders, err := srv.Recv()
		if err == io.EOF {
			fmt.Printf("Communicate result is: %v", result)
			break
		}

		if err != nil {
			log.Fatal("Err receiving orders from client: %v", err)
		}

		fmt.Printf("Communicate received orders: %v", orders)
		fmt.Println()
		result = append(result, orders)

		count++
		if count%3 == 0 {
			err := srv.Send(orders)
			if err != nil {
				log.Fatal("Err sending orders to client: %v", err)
			}
		}
	}
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
