package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func ImageSku(app *echo.Echo) {
	image := &handler.ImageSku{}
	group := app.Group("/api/v1/images")
	group.DELETE("", image.DeleteBySkuId)
}
