package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func ImgSku(e *echo.Echo) {
	handler := handler.NewImgSku()
	group := e.Group("/api/v1/images/sku")
	group.POST("", handler.Set)
	group.DELETE("", handler.DeleteAllImgBySku)
}
