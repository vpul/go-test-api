package main

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterV1Routes(v1 fiber.Router) {
	v1.Get("/info", InfoHandler) // /v1/info
}
