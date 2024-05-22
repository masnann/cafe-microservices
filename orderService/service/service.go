package service

import (
	"order-service/repository"
)

type OrderService struct {
	repo repository.OrderRepositoryInterface
}

func NewOrderService(repo repository.OrderRepositoryInterface) OrderServiceInterface {
	return OrderService{
		repo: repo,
	}
}

