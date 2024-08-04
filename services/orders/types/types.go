package types	 

import (
	"context"

	"github.com/Laoooq/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, *orders.Order) error
} 