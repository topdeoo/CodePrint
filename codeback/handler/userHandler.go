package handler

import "github.com/gofiber/fiber/v2"

type UserReq struct {
	TeamName string `json:"teamname"`
	Password string `json:"password"`
	Location string `json:"location"`
}

type User struct {
	TeamName string
	Password string
	Location string
	IPAddr   string
}

func LoginHandler(ctx *fiber.Ctx) error {

	var userReq UserReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
	})
}

func PrintHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Success",
	})
}
