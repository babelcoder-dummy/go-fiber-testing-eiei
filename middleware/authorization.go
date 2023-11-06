package middleware

import (
	"net/http"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/repository"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Authorizer(userRepo repository.User) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("claims").(model.UserClaims)
		user, err := userRepo.FindByID(claims.ID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(err)
		}

		enforcer, err := casbin.NewEnforcer("config/acl_model.conf", "config/policy.csv")
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		ok, err := enforcer.Enforce(string(user.Role), c.Path(), c.Method())
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		if !ok {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "you are not allowed to access this resource"})
		}

		return c.Next()
	}
}
