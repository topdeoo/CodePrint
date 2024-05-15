package middleware

import (
	"acm.nenu.edu.cn/xcpc/global"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Init() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(global.SecretKey),
	})
}
