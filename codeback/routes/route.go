package routes

import (
	"acm.nenu.edu.cn/xcpc/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App, jwt *fiber.Handler) {
	g := app.Group("/api/v1/client")
	g.Post("/login", handler.LoginHandler)
	g.Post("/print", *jwt, handler.PrintHandler)
}
