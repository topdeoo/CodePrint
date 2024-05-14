package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topdeoo/codeprint/back/handler"
)

func RouteInit(app *fiber.App) {
	g := app.Group("/api/v1/client")
	g.Post("/login", handler.LoginHandler)
	g.Post("/print", handler.PrintHandler)
}