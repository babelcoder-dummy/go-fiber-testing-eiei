package dto

import (
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
)

type UserResponse struct {
	ID      uint       `json:"id"`
	Email   string     `json:"email"`
	Avatar  string     `json:"avatar"`
	Name    string     `json:"name"`
	Address string     `json:"address"`
	Tel     string     `json:"tel"`
	Role    model.Role `json:"role"`
}
