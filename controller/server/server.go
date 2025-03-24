package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// UserHandler contains user-related route handlers
type UserHandler struct{}

// GetUser handles retrieving user information
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	user := fiber.Map{
		"id":   1,
		"name": "inky-7.3",
		"ip":   "192.168.88.202",
	}

	return c.JSON(user)
}

func New() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("PiDisplay Controller API")
	})

	// API
	v1 := app.Group("/api/v1")

	displayHandler := &UserHandler{}
	users := v1.Group("/displays")
	users.Get("/", displayHandler.GetUser)

	return app
}
