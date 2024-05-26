package handler

import (
	"os"
	"time"

	"acm.nenu.edu.cn/xcpc/global"
	"acm.nenu.edu.cn/xcpc/model"
	"acm.nenu.edu.cn/xcpc/pkg/tasks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	hour := time.Hour * 5
	ipAddr := ctx.IP()

	claims := jwt.MapClaims{
		"teamName": user.TeamName,
		"password": user.Password,
		"location": user.Location,
		"ipaddr":   ipAddr,
		"exp":      time.Now().Add(hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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

func MockLoginHandler(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.SecretKey), nil
	})

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	if !token.Valid {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func PrintHandler(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.SecretKey), nil
	})

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	if !token.Valid {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	claims := token.Claims.(jwt.MapClaims)
	IPAddr := claims["ipaddr"].(string)
	Password := claims["password"].(string)

	if IPAddr != ctx.IP() {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if Password != global.Database[claims["teamName"].(string)].Password {
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

	Teamname := claims["teamName"].(string)
	Location := claims["location"].(string)

	filename := global.MyConfig.CodePath + "/teamname--" + Teamname + "--location--" + Location + "-time-" + time.Now().Format("20060102150405") + ".code"

	endline := "---------------------------------------END-------------------------------------------"

	rawFile := filename + "\n\n" + rawcode.Code + "\n\n" + endline
	go tasks.PrintCode(rawFile)

	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	f.WriteString(rawFile)
	f.Close()

	return ctx.JSON(fiber.Map{
		"message": "Success",
	})
}
