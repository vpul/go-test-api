package middleware

import (
	"strings"

	"go-test-api/cmd/api/helpers/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func IsAuthenticated(c *fiber.Ctx) error {
	headerVal := c.Get("Authorization")

	if headerVal == "" {
		log.Error("no auth header provided")
		return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}

	// Get the token from the bearer
	var jwtToken = strings.Split(headerVal, " ")[1]

	if jwtToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}

	claims, err := token.Parse(jwtToken)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	c.Locals("user", claims)

	return c.Next()
}
