package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topdeoo/codeprint/back/routes"
)

func main() {

	app := fiber.New()

	routes.RouteInit(app)

}
