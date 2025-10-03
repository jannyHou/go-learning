package main

import (
	"context"
	pb "examples/go-learning/grpc/order_proto"
	"fmt"
	"io"
	"log"
	"sync"
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

	// createOrder
	order := pb.Order{
		Id:      1,
		Food:    "rice",
		Payment: 10,
	}

	rc, err := c.CreateOrder(ctx, &order)
	fmt.Printf("Create created order: %v", rc)
	fmt.Println()

	var wg sync.WaitGroup
	// getOrder
	rg, err := c.GetOrder(ctx, rc)
	if err != nil {
		log.Fatal("Error getting order", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			rorder, err := rg.Recv()

			if err == io.EOF {
				return
			}

			if err != nil {
				log.Fatal("Error reading order: %v", err)
			}
			fmt.Printf("Get received order (server streaming): %v", rorder.String())
			fmt.Println()
		}

	}()

	// communicate
	rcm, err := c.Communicate(ctx)
	if err != nil {
		log.Fatal("Error communicating with server", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		var orders = []*pb.Order{
			{Id: 1, Food: "rice", Payment: 10},
			{Id: 2, Food: "rice", Payment: 10},
			{Id: 3, Food: "rice", Payment: 10},
			{Id: 4, Food: "rice", Payment: 10},
			{Id: 5, Food: "rice", Payment: 10},
			{Id: 6, Food: "rice", Payment: 10},
		}

		for i := 0; i < len(orders); i++ {
			err := rcm.Send(&pb.Orders{Orders: []*pb.Order{orders[i]}})
			if err != nil {
				log.Fatal("Cannot send ordres to server %v", err)
			}
		}

		if err := rcm.CloseSend(); err != nil {
			log.Fatal("Cannot close send method %v", err)
		}

		for {
			rcmr, err := rcm.Recv()

			if err == io.EOF {
				return
			}

			if err != nil {
				log.Fatal("Error receiving result from communication server: %v", err)
			}

			fmt.Printf("Received result from communication server: %v", rcmr)
			fmt.Println()
		}

	}()

	wg.Wait()

}
