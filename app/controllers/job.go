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

func CreateJob(c *fiber.Ctx) error {
	job := &entity.Job{}

	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": true,
			"msg":   "Token is invalid or expired",
		})
	}

	// Set expiration time from JWT data of current job.
	expires := claims.Expires

	// Return status 401 and unauthorized error message.
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "You are unauthorized, please login",
		})
	}

	// Check, if received JSON data is valid.
	err = c.BodyParser(&job)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	// Create a new validator for a Job model.
	validate := utils.NewValidator()

	job.ID = uuid.New()
	job.CreatedAt = time.Now()
	job.CreatedBy = claims.UserID

	// Validate job fields.
	if err := validate.Struct(job); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": true,
			"msg":   "Request failed",
		})
	}

	jb, err := job.CreateJob()

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"error":   true,
			"message": "Could not create job",
		})

	}

	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"error": false,
		"msg":   "Successfully created a job",
		"data":  jb,
	})

}

func GetAllJob(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": true,
			"msg":   "Token is invalid or expired",
		})
	}

	// Set expiration time from JWT data of current job.
	expires := claims.Expires

	// Return status 401 and unauthorized error message.
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "You are unauthorized, please login",
		})
	}

	createdBy := claims.UserID

	jobs, err := entity.GetAllJob(createdBy)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"error":   true,
			"message": "Could not get jobs",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"error": false,
		"msg":   "Successfully retrieved all jobs",
		"data":  jobs,
		"count": len(jobs),
	})

}

func GetJob(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": true,
			"msg":   "Token is invalid or expired",
		})
	}

	// Set expiration time from JWT data of current job.
	expires := claims.Expires
	fmt.Println("now", now)
	fmt.Println("expire", expires)

	// Return status 401 and unauthorized error message.
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "You are unauthorized, please login",
		})
	}

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": true,
			"msg":   "Error parsing uuid",
		})
	}

	createdBy := claims.UserID

	job, _, err := entity.GetSingleJob(id, createdBy)

	if id != job.ID {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "No job with the ID",
			"data":    id})
	}

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"error": true,
			"msg":   "Could not fetch job",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"error": false,
		"msg":   "Successfully retrieved job",
		"data":  job,
	})
}
