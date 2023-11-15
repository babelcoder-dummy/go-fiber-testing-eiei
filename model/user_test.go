package model_test

import (
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUser_GenerateEncryptedPassword(t *testing.T) {
	password := utils.UUIDv4()
	user := model.User{
		Password: password,
	}

	user.Password = user.GenerateEncryptedPassword()
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	assert.NoError(t, err)
}
