package routes

import (
	"go-test-api/cmd/api/handlers"
	middleware "go-test-api/cmd/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterV1(app *fiber.App) {
	// API v1 group
	v1 := app.Group("/v1")

	v1.Get("/info", handlers.AppInfoHandler)                            // /v1/info
	v1.Get("/my", middleware.IsAuthenticated, handlers.UserInfoHandler) // /v1/my
}
