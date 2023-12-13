package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type FailureResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}

func AppInfoHandler(c *fiber.Ctx) error {
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

func UserInfoHandler(c *fiber.Ctx) error {
	user := c.Locals("user")

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Status:  "success",
		Payload: user,
	})
}
