package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topdeoo/codeprint/back/global"
	"github.com/topdeoo/codeprint/back/middleware"
	"github.com/topdeoo/codeprint/back/pkg/tasks"
	"github.com/topdeoo/codeprint/back/routes"
)

func main() {

	go tasks.Start()

	global.Init()

	app := fiber.New()

	jwt := middleware.Init()

	routes.RouteInit(app, &jwt)

	app.Listen(global.MyConfig.GetHttpConfig().Host + ":" + global.MyConfig.GetHttpConfig().Port)

}
