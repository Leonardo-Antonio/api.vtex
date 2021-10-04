package app

import (
	"fmt"

	"github.com/Leonardo-Antonio/api.pdq-prices/src/router"
	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/enum"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	app *echo.Echo
}

func New() *server {
	return &server{echo.New()}
}

func (s *server) Middlewares() {
	/* s.app.Use(middleware.Logger()) */
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.CORS())
}

func (s *server) Routers() {
	router.Price(s.app)
	router.Stock(s.app)
	router.Category(s.app)
	router.Slug(s.app)
	router.Formating(s.app)
}

func (s *server) Listening() {
	s.app.Start(fmt.Sprintf(":%s", enum.Env.PORT))
}
