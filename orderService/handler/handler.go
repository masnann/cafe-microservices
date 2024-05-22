package handler

import (
	"net/http"
	"order-service/service"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	service service.OrderServiceInterface
}

func NewOrderHandler(service service.OrderServiceInterface) OrderHandler {
	return OrderHandler{
		service: service,
	}
}

func (h OrderHandler) GetListOrder(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Order")
}
