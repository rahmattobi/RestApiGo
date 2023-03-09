package handlers

import (
	"auth-jwt/database"
	"auth-jwt/models/entity"
	"auth-jwt/models/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HandlerUser(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func HandlerUserGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var users entity.User
	err := database.DB.First(&users, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}

func HandlerUserInput(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "Failed",
			"Error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateuser := database.DB.Create(&newUser).Error

	if errCreateuser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Message": "Failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func HandlerUserUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "Bad Request",
		})
	}
	var users entity.User
	userId := ctx.Params("id")
	err := database.DB.First(&users, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Update Data User
	if userRequest.Name != "" {
		users.Name = userRequest.Name
	}

	users.Address = userRequest.Address
	users.Phone = userRequest.Phone
	users.Email = userRequest.Email

	errUpdate := database.DB.Save(&users).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}

func HandlerUserDelete(ctx *fiber.Ctx) error {
	var users entity.User
	userId := ctx.Params("id")
	err := database.DB.First(&users, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Delete(&users).Error

	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})

}
