package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topdeoo/codeprint/back/global"
	"github.com/topdeoo/codeprint/back/routes"
)

func main() {

	global.Init()

	app := fiber.New()

	routes.RouteInit(app)

	app.Listen(global.MyConfig.GetHttpConfig().Host + ":" + global.MyConfig.GetHttpConfig().Port)

}
