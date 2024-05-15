package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CrosInit(app *fiber.App) {
	app.Use(cors.New())
}
