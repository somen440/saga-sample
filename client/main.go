package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/somen440/saga-sample/proto"
	"google.golang.org/grpc"
)

const (
	orderAddr = ":5002"
)

func main() {
	itemID := flag.Int64("itemid", 0, "Item ID")
	orderID := flag.Int64("orderid", 0, "Order ID")
	method := flag.String("method", "", "method create/get")
	flag.Parse()

	if !(*method == "create" || *method == "get") {
		log.Fatal("not allowed method ", *method)
	}

	conn, err := grpc.Dial(orderAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewOrderClient(conn)

	switch *method {
	case "create":
		err = create(client, *itemID)
	case "get":
		err = get(client, *orderID)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func create(client proto.OrderClient, itemID int64) error {
	ctx := context.Background()
	in := &proto.CreateOrderRequest{
		ConsumerId: 1,
		ItemId:     itemID,
	}

	reply, err := client.CreateOrder(ctx, in)
	if err != nil {
		return err
	}

	log.Print(reply)

	return nil
}

func get(client proto.OrderClient, orderID int64) error {
	ctx := context.Background()
	in := &proto.GetOrderRequest{
		OrderId: orderID,
	}

	reply, err := client.GetOrder(ctx, in)
	if err != nil {
		return err
	}

	log.Print(reply)

	return nil
}
