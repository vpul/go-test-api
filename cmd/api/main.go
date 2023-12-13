package main

import (
	"errors"
	"go-test-api/internal/config"
	"log"

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

			ctx.Status(e.Code).JSON(FailureResponse{
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

	v1 := app.Group("/v1")
	RegisterV1Routes(v1)

	// 404 Not found Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "error",
			"message": "Not found",
		})
	})

	app.Listen(":3000")
}
