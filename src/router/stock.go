package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func Stock(app *echo.Echo) {
	handler := &handler.Stock{}
	group := app.Group("/api/v1/stocks")
	group.PUT("", handler.Set)
}
