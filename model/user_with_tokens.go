package model

type UserWithTokens struct {
	User         *User
	AccessToken  string
	RefreshToken string
}
