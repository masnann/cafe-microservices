package routes

import (
	"order-service/handler"

	"github.com/labstack/echo/v4"
)

func RoutesApi(e *echo.Echo, handler handler.OrderHandler) {

	public := e.Group("/public")
	public.GET("/", handler.GetListOrder)
}