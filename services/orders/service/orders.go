package service

import (
	"context"

	"grpc-kitchen/services/common/genproto/orders/github.com/Laoooq/common/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// store

}

func NewOrderServce() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}
