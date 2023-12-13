package controllers

import (
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

func UpdateJob(c *fiber.Ctx) error {
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

	updateJob := &entity.UpdateJob{}

	// Check, if received JSON data is valid.
	err = c.BodyParser(&updateJob)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"error": true,
			"msg":   "Request failed",
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

	updatedJob, db, err := entity.GetSingleJob(id, createdBy)

	if id != updatedJob.ID {
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

	updateJob.ID = id
	updateJob.CreatedBy = createdBy
	updateJob.UpdatedAt = time.Now()

	// Create a new validator for a Job model.
	validate := utils.NewValidator()

	// Validate job fields.
	if err := validate.Struct(updateJob); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if updateJob.Company != "" {
		updatedJob.Company = updateJob.Company
	}

	if updateJob.Position != "" {
		updatedJob.Position = updateJob.Position
	}

	if updateJob.Status != "" {
		updatedJob.Status = updateJob.Status
	}

	db.Save(&updatedJob)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"error": true,
		"msg":   "Successfully updated job",
		"data":  updatedJob,
	})
}

func DeleteJob(c *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": true,
			"msg":   "Token is invalid or expired",
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

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
			"msg":   err.Error(),
		})
	}

	createdBy := claims.UserID

	uJob, _, err := entity.GetSingleJob(id, createdBy)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"error": true,
			"msg":   "Could not fetch job",
		})
	}

	if id == uJob.ID {
		_, err := entity.DeleteJob(id, createdBy)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"error": true,
				"msg":   "Could not delete job",
			})
		} else {
			return c.Status(http.StatusOK).JSON(&fiber.Map{
				"error": false,
				"msg":   "Successfully deleted job",
				"data":  id,
			})
		}

	} else {
		if id != uJob.ID {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"error":   true,
				"message": "No job with the ID",
				"data":    id,
			})
		}
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"msg": "OK",
	})
}
