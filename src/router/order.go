package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func Order(app *echo.Echo) {
	handler := &handler.Order{}
	group := app.Group("/api/v1/orders")
	group.DELETE("", handler.Cancel)
}
