package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/udodinho/job-app/domain/entity"
	"github.com/udodinho/job-app/pkg/utils"
)

func Register(c *fiber.Ctx) error {
	user := &entity.User{}

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"msg": "Request failed",
		})
	}

	exist, err := entity.GetUserByEmail(user.Email)

	fmt.Println("ex", exist)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"msg": "Failed to get user",
		})
	}

	if exist.Email == user.Email {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Email already exist",
		})
	}

	validate := utils.NewValidator()
	
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Password = utils.GeneratePassword(user.Password)

	fmt.Println("Here0")
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": true,
			"msg":   "Validation error",
		})
	}
	fmt.Println("Here1")
	newUser, err := user.CreateUser()
	// newUser := app.Db.CreateUser(user)
	fmt.Println("Here2")

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"msg": "Failed to create user",
		})

		return err
	}

	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"msg": "Successfully created user",
		"data": newUser,
	})

}

func Login(c *fiber.Ctx) error {
	user := &entity.User{}

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"msg": "Request failed",
		})
	}

	exist, err := entity.GetUserByEmail(user.Email)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"error": true,
			"msg": "User does mot exist",
		})
	}

	comparePassword := utils.ComparePasswords(exist.Password, user.Password)

	if !comparePassword {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": true,
			"msg": "Incorrect email or password",
		})
	}

	// Generate a new pair of access and refresh tokens.
	token, err := utils.GenerateNewTokens(exist.ID.String())

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"error": false,
		"msg": "Successfully logged in",
		"tokens": &fiber.Map{
			"accessToken": token.Access,
			"refreshToken": token.Refresh,
		},
	})
}
