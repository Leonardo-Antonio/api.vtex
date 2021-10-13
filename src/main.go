package main

import (
	"log"

	"github.com/Leonardo-Antonio/api.pdq-prices/src/app"
	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/enum"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	enum.GetEnv()

	app := app.New()
	app.Middlewares()
	app.Routers()
	app.Listening()
}
