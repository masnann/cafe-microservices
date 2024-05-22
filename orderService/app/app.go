package app

import (
	"database/sql"
	"order-service/handler"
	"order-service/repository"
	"order-service/service"
)

func SetupApp(DB *sql.DB) handler.OrderHandler {
	repo := repository.NewOrderRepository(DB)
	service := service.NewOrderService(repo)
	
	return handler.NewOrderHandler(service)
}