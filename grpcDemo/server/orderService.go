package main

import (
	"context"
	"fmt"
	. "grpcDemo/proto/package/hello"
)

type OrderService struct {

}

func (t *OrderService) NewOrder(ctx context.Context, in *OrderModel) (*OrderResponse, error) {
	fmt.Println(in)
	return &OrderResponse{Status: 1, Message: "success"}, nil
}
