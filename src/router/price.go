package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func Price(app *echo.Echo) {
	handler := &handler.Price{}
	group := app.Group("/api/v1/prices")
	group.PUT("", handler.Set)
}
