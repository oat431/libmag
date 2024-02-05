package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/oat431/libmag/configs"
	"github.com/oat431/libmag/controllers"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(configs.SetupCors()))

	healthCheck := app.Group("/api/v1")
	{
		healthCheck.Get("/health-check", controllers.HealthCheck)
	}

	return app
}
