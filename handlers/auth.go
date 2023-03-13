package handlers

import (
	"auth-jwt/database"
	"auth-jwt/models/entity"
	"auth-jwt/models/request"
	"auth-jwt/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.Login)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "Bad Request",
		})
	}

	// validasi login
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "Wrong creadential",
			"error":   errValidate,
		})
	}

	// pengecekan email
	var users entity.User
	err := database.DB.First(&users, "email = ?", loginRequest.Email).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// validasi password
	isValid := utils.CheckPassword(loginRequest.Password, users.Password)

	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Password salah",
		})
	}

	// generate jwt
	claims := jwt.MapClaims{}
	claims["name"] = users.Name
	claims["email"] = users.Email
	claims["address"] = users.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if users.Email == "nabila@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)

	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}
	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
