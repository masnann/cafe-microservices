package app

import (
	"database/sql"
	"kitchen-service/handler"
	"kitchen-service/repository"
	"kitchen-service/service"
)

func SetupApp(DB *sql.DB) handler.KitchenHandler {
	repo := repository.NewKitchenRepository(DB)
	service := service.NewKitchenService(repo)
	
	return handler.NewKitchenHandler(service)
}