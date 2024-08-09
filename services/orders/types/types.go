package types

import (
	"context"

	"grpc-kitchen/services/common/genproto/orders/github.com/Laoooq/common/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
