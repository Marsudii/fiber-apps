package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	// Middleware untuk mencatat log setiap request
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} - ${method} ${path} ${status} ${latency} - ${ua}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// Route utama
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Backend service is running (use docker prune)",
		})
	})

	// Route lain untuk testing
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "This is a test route"})
	})

	// Port dari environment variable atau default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
