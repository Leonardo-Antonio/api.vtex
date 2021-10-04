package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func Category(app *echo.Echo) {
	handler := &handler.Category{}
	group := app.Group("/api/v1/categories")
	group.POST("", handler.Create)
}
