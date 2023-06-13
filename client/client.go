package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/shamskhalil/grpcApp/orderpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, _ := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer cc.Close()
	client := orderpb.NewEcommerceServiceClient(cc)
	fmt.Println("---------------------UNARY------------------------")
	orderPhone(client)
	fmt.Println("--------------------------------------------------")
	fmt.Println("----------------SERVER STREAMIG-------------------")
	getOrders(client)
	fmt.Println("--------------------------------------------------")

}

func orderPhone(client orderpb.EcommerceServiceClient) {
	now := time.Now()
	req := &orderpb.PlaceOrderRequest{Item: "Infinix", Qty: 1, Price: 95000.00}
	res, _ := client.PlaceOrder(context.Background(), req)
	elapsed := time.Since(now)
	fmt.Printf("Response: %v\n took:%s\n", res, elapsed)
}

func getOrders(client orderpb.EcommerceServiceClient) {
	now := time.Now()
	stream, _ := client.GetOrderItems(context.Background(), &orderpb.GetOrderItemsRequest{})

	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("RECIEVED ORDER ITEM ROME SERVER: \n%v\n", item)
	}
	elapsed := time.Since(now)
	fmt.Printf("End of stream\ntook:%s\n", elapsed)
}
