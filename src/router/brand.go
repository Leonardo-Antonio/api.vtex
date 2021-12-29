package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

var Brand = func(app *echo.Echo) {
	handlerBrand := handler.NewBrand()
	group := app.Group("/api/v1/brands")
	group.POST("", handlerBrand.Create)
}
