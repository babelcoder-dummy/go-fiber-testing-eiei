package dto

import (
	"mime/multipart"
)

type AuthForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateProfileForm struct {
	Email   string               `form:"email"`
	Name    string               `form:"name"`
	Address string               `form:"address"`
	Tel     string               `form:"tel"`
	Avatar  multipart.FileHeader `form:"avatar"`
}

type AuthProfileWithTokenResponse struct {
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
	User         UserResponse `json:"user"`
}

type RefreshTokenForm struct {
	RefreshToken string `json:"refreshToken"`
}
