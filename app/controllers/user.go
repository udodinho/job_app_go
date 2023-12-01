package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/udodinho/job-app/domain/entity"

	// "github.com/udodinho/job-app/infrastructure/repository"
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
fmt.Println("Here")
fmt.Println(user.Name)
fmt.Println(user.Email)
fmt.Println(user.Password)
fmt.Println("Here")
	// validate := utils.NewValidator()

	// fmt.Println("Here0")
	// if err := validate.Struct(user); err != nil {
	// 	c.Status(http.StatusBadRequest).JSON(fiber.Map{

	// 		"error": true,
	// 		"msg":   utils.ValidatorErrors(err),
	// 	})
	// 	return err
	// }
	// fmt.Println("Here1")

	fmt.Println("start")
	// exist, err := app.Db.GetUserByEmail(user.Email)
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


