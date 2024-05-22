package routes

import (
	"kitchen-service/handler"

	"github.com/labstack/echo/v4"
)

func RoutesApi(e *echo.Echo, handler handler.KitchenHandler) {

	public := e.Group("/public")
	public.GET("/", handler.GetListOrder)
}