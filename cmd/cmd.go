package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())

	app.Get("/health_check", func(c *fiber.Ctx) error {
		return nil
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/lambda", Lambda)

	app.Listen(":4000")
}

type Result struct {
	Status   bool   `json:"status"`
	Hostname string `json:"hostname"`
}

func Lambda(c *fiber.Ctx) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	result := Result{
		Status:   true,
		Hostname: hostname,
	}

	return c.JSON(result)
}
