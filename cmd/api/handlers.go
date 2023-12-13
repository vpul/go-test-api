package main

import (
	"github.com/gofiber/fiber/v2"
)

func InfoHandler(c *fiber.Ctx) error {
	type Info struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	var info = Info{
		Name:    "go-test-api",
		Version: "1.0.0",
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Status:  "success",
		Payload: info,
	})
}
