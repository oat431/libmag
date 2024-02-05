package configs

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}
}
