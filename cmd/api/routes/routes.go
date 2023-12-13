package routes

import (
	"go-test-api/cmd/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterV1(app *fiber.App) {
	// API v1 group
	v1 := app.Group("/v1")

	v1.Get("/info", handlers.InfoHandler) // /v1/info
}
