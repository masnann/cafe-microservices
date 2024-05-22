package handler

import (
	"kitchen-service/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KitchenHandler struct {
	service service.KitchenServiceInterface
}

func NewKitchenHandler(service service.KitchenServiceInterface) KitchenHandler {
	return KitchenHandler{
		service: service,
	}
}

func (h KitchenHandler) GetListOrder(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
