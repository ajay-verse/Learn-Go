package orders

import (
	// Go Internal Packages
	"context"
	"learn-go/utils/helpers"
	"time"

	// Local Packages
	omodels "learn-go/models/orders"
)

type OrdersRepository interface {
	Insert(ctx context.Context, order omodels.Order) error
	Get(ctx context.Context, orderID string) (omodels.Order, error)
	Update(ctx context.Context, order omodels.Order) error
	Delete(ctx context.Context, orderID string) error
}

type OrdersService struct {
	ordersRepository OrdersRepository
}

func NewService(ordersRepository OrdersRepository) *OrdersService {
	return &OrdersService{ordersRepository: ordersRepository}
}

func (s *OrdersService) Insert(ctx context.Context, order omodels.Order) error {
	order.OrderID = helpers.GenerateOrderID()
	currTime := time.Now()
	order.CreatedAt = currTime
	order.UpdatedAt = currTime
	return s.ordersRepository.Insert(ctx, order)
}

func (s *OrdersService) Get(ctx context.Context, orderID string) (omodels.Order, error) {
	return s.ordersRepository.Get(ctx, orderID)
}

func (s *OrdersService) Update(ctx context.Context, order omodels.Order) error {
	order.UpdatedAt = time.Now()
	return s.ordersRepository.Update(ctx, order)
}

func (s *OrdersService) Delete(ctx context.Context, orderID string) error {
	return s.ordersRepository.Delete(ctx, orderID)
}
