package handler

import (
	"os"
	"time"

	"acm.nenu.edu.cn/xcpc/global"
	"acm.nenu.edu.cn/xcpc/model"
	"acm.nenu.edu.cn/xcpc/pkg/tasks"
	"github.com/gofiber/fiber/v2"
	jwttoken "github.com/golang-jwt/jwt/v4"
)

func LoginHandler(ctx *fiber.Ctx) error {

	var userReq model.UserReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, ok := global.Database[userReq.TeamName]

	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Team Name Wrong or Password Wrong",
		})
	}

	hour := time.Hour
	ipAddr := ctx.IP()

	claims := jwttoken.MapClaims{
		"TeamName": user.TeamName,
		"Password": user.Password,
		"Location": user.Location,
		"IPAddr":   ipAddr,
		"Exp":      time.Now().Add(hour).Unix(),
	}

	token := jwttoken.NewWithClaims(jwttoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(global.SecretKey))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.UserResp{
		Message: "Success",
		Token:   t,
	})
}

func PrintHandler(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwttoken.Token)
	claims := user.Claims.(jwttoken.MapClaims)
	err := claims.Valid()
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	IPAddr := claims["IPAddr"].(string)
	if IPAddr != ctx.IP() {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var rawcode model.CodeReq
	if err := ctx.BodyParser(&rawcode); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	Teamname := claims["TeamName"].(string)
	Password := claims["Password"].(string)
	Location := claims["Location"].(string)

	filename := global.MyConfig.CodePath + "/teamname-" + Teamname + "-password-" +
		Password + "-location-" + Location + time.Now().Format("20060102150405") + ".code"

	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	f.WriteString(filename + "\n\n")
	f.WriteString(rawcode.Code + "\n\n")
	f.WriteString("---------------------------------------END-------------------------------------------")
	f.Close()

	go tasks.PrintCode(filename)

	return ctx.JSON(fiber.Map{
		"message": "Success",
	})
}
