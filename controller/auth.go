package controller

import (
	"errors"
	"time"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Auth struct {
	Service service.Auth
}

func (a *Auth) GetProfile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(model.UserClaims)
	profile, err := a.Service.GetProfile(claims.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.SendStatus(fiber.StatusNotFound)
	}

	var serializedProfile dto.UserResponse
	copier.Copy(&serializedProfile, &profile)

	return c.JSON(serializedProfile)
}

func (a *Auth) SignUp(c *fiber.Ctx) error {
	form := new(dto.AuthForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := validate.Struct(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := a.Service.SignUp(form)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var serializedUser dto.UserResponse
	copier.Copy(&serializedUser, &user)

	return c.Status(fiber.StatusCreated).JSON(serializedUser)
}

func (a *Auth) SignIn(c *fiber.Ctx) error {
	form := new(dto.AuthForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	userWithTokens, err := a.Service.SignIn(form)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	a.setRefreshTokenCookie(c, userWithTokens.RefreshToken)

	var serializedProfile dto.AuthProfileWithTokenResponse
	copier.Copy(&serializedProfile, userWithTokens)

	return c.Status(fiber.StatusCreated).JSON(serializedProfile)
}

func (a *Auth) SignOut(c *fiber.Ctx) error {
	claims := c.Locals("claims").(model.UserClaims)
	a.Service.SignOut(claims.ID)
	// c.ClearCookie("refreshToken")
	a.clearRefreshTokenCookie(c)

	return c.SendStatus(fiber.StatusNoContent)
}

func (a *Auth) RemoveCookie(c *fiber.Ctx) error {
	// c.ClearCookie("refreshToken")
	a.clearRefreshTokenCookie(c)

	return c.SendStatus(fiber.StatusNoContent)
}

func (a *Auth) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")
	if refreshToken == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userWithTokens, err := a.Service.RefreshToken(refreshToken)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	a.setRefreshTokenCookie(c, userWithTokens.RefreshToken)

	var serializedProfile dto.AuthProfileWithTokenResponse
	copier.Copy(&serializedProfile, userWithTokens)

	return c.Status(fiber.StatusCreated).JSON(serializedProfile)
}

func (a *Auth) UpdateProfile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(model.UserClaims)
	form := new(dto.UpdateProfileForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := validate.Struct(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userForm := new(dto.UpdateProfileForm)
	copier.Copy(userForm, form)
	avatar, _ := c.FormFile("avatar")
	user, err := a.Service.Update(claims.ID, avatar, userForm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var serializedProfile dto.UserResponse
	copier.Copy(&serializedProfile, &user)

	return c.JSON(serializedProfile)
}

func (a *Auth) setRefreshTokenCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "lax",
	})
}

func (a *Auth) clearRefreshTokenCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    "",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "lax",
		Expires:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	})
}
