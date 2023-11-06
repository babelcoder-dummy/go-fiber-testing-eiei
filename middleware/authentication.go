package middleware

import (
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticator(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	c.Locals("claims", model.UserClaims{
		ID:   uint(claims["sub"].(float64)),
		Role: model.Role(claims["role"].(string)),
	})

	return c.Next()
}
