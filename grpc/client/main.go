package main

import (
	"context"
	pb "examples/go-learning/grpc/order_proto"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot creater client %v", err)
	}

	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	order := pb.Order{
		Id:      1,
		Food:    "rice",
		Payment: 10,
	}

	rc, err := c.CreateOrder(ctx, &order)
	fmt.Printf("result of creating order is: %v", rc)

	rg, err := c.GetOrder(ctx, rc)
	if err != nil {
		log.Fatal("Error getting order", err)
	}

	done := make(chan bool)
	go func() {
		for {
			rorder, err := rg.Recv()

			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatal("Error reading order: %v", err)
			}
			fmt.Printf("Received order: %v", rorder.String())
		}

	}()

	<-done
}
