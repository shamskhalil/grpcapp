package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shamskhalil/grpcApp/orderpb"
	"google.golang.org/grpc"
)

var db []*orderpb.PlaceOrderRequest

type server struct {
	orderpb.UnimplementedEcommerceServiceServer
}

func (*server) PlaceOrder(ctx context.Context, req *orderpb.PlaceOrderRequest) (*orderpb.PlaceOrderResponse, error) {
	db = append(db, req)
	msg := fmt.Sprintf("Thank you for ordering %d, %s! Order has been accepted!", req.Qty, req.Item)
	return &orderpb.PlaceOrderResponse{Message: msg}, nil
}

func (*server) GetOrderItems(in *orderpb.GetOrderItemsRequest, stream orderpb.EcommerceService_GetOrderItemsServer) error {
	for _, item := range db {

		obj := &orderpb.GetOrderItemResponse{
			Item:  item.Item,
			Qty:   item.Qty,
			Price: item.Price,
		}
		stream.Send(obj)
		//time.Sleep(5 * time.Second)
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", "0.0.0.0:30000")
	s := grpc.NewServer()
	log.Println("Server listening on port 30000")
	orderpb.RegisterEcommerceServiceServer(s, &server{})
	s.Serve(lis)
}
