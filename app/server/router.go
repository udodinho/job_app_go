package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/udodinho/job-app/app/controllers"
	"github.com/udodinho/job-app/pkg/middleware"
	"gorm.io/gorm"
)

type Server struct {
	Db	*gorm.DB
}

// SetupRouter set up the router for the server
func (s *Server) SetupRouter(app *fiber.App) {
	api := app.Group("/api/v1")
	// Health check route
	api.Get("/healthcheck", controllers.HealthCheck)

	// Unauthenticated routes
	api.Post("/auth/register", controllers.Register)
	api.Post("/auth/login", controllers.Login)
	
	// Authenticated routes
	api.Post("/job", middleware.JWTProtected(), controllers.CreateJob)
	api.Get("/job", middleware.JWTProtected(), controllers.GetAllJob)
	api.Get("/job/:id", middleware.JWTProtected(), controllers.GetJob)
	api.Patch("/job/:id", middleware.JWTProtected(), controllers.UpdateJob)
	api.Delete("/job/:id", middleware.JWTProtected(), controllers.DeleteJob)
}

// NotFoundRoute func for describe 404 Error route.
func (s *Server) NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "Sorry, this route does not exist",
			})
		},
	)
}
