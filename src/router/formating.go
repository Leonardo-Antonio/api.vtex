package router

import (
	"github.com/Leonardo-Antonio/api.pdq-prices/src/handler"
	"github.com/labstack/echo/v4"
)

func Formating(app *echo.Echo) {
	handler := &handler.Formating{}
	group := app.Group("/api/v1/formating")
	group.POST("", handler.SearchSkuEanAndRemplaceCoincidences)
	group.POST("/specifications", handler.SearchAndRemplaceTypeProductDigitalBySpecification)
}
