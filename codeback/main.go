package main

import (
	"acm.nenu.edu.cn/xcpc/global"
	"acm.nenu.edu.cn/xcpc/middleware"
	"acm.nenu.edu.cn/xcpc/pkg/tasks"
	"acm.nenu.edu.cn/xcpc/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	global.Init()

	app := fiber.New()

	jwt := middleware.Init()

	routes.RouteInit(app, &jwt)

	go tasks.Start()

	app.Listen(global.MyConfig.GetHttpConfig().Host + ":" + global.MyConfig.GetHttpConfig().Port)

}
