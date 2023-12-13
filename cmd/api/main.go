package main

import (
	"errors"
	"log"

	"go-test-api/cmd/api/handlers"
	"go-test-api/cmd/api/routes"
	"go-test-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	config.LoadEnv()
}

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var e *fiber.Error

			ok := errors.As(err, &e)

			if !ok {
				log.Println(err.Error())
				e = fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
			}

			ctx.Status(e.Code).JSON(handlers.FailureResponse{
				Status:  "error",
				Message: e.Message,
			})
			return nil
		},
	})

	// Middlewares
	app.Use(cors.New())
	app.Use(recover.New()) // Recover from panics anywhere in the stack chain
	app.Use(logger.New())

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API running successfully!")
	})

	routes.RegisterV1(app)

	// 404 Not found Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(handlers.FailureResponse{
			Status:  "error",
			Message: "Not found",
		})
	})

	app.Listen(":3000")
}
