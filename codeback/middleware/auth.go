package middleware

import (
	"acm.nenu.edu.cn/xcpc/global"
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func AuthInit(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(global.SecretKey)},
		Filter:     func(c *fiber.Ctx) bool { return c.Path() == "/api/v1/client/login" },
	}))
}
