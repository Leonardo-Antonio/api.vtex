package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func Slug(app *echo.Echo) {
	handler := &handler.Slug{}
	group := app.Group("/api/v1/slugs")
	group.POST("", handler.SetByNameProduct)
}
