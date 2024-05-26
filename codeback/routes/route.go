package routes

import (
	"acm.nenu.edu.cn/xcpc/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	g := app.Group("/api/v1/client")
	g.Post("/login", handler.LoginHandler)
	g.Get("/mock", handler.MockLoginHandler)
	g.Post("/print", handler.PrintHandler)
}
