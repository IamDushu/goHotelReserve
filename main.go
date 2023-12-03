package main

import (
	"flag"

	"github.com/IamDushu/goHotelReserve/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listen", ":5000", "The listen address of the API Server")
	flag.Parse()
	//TRY: make build and ./bin/api --listen :7000

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
