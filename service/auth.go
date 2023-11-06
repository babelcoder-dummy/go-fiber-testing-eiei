package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/config"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/repository"
	"github.com/jinzhu/copier"
)

type Auth struct {
	Token          Token
	Storage        Storage
	UserRepository repository.User
}

func (a *Auth) GetProfile(id uint) (*model.User, error) {
	return a.UserRepository.FindByID(id)
}

func (a *Auth) SignUp(form *dto.AuthForm) (*model.User, error) {
	user := new(model.User)

	copier.Copy(user, form)
	user.Password = user.GenerateEncryptedPassword()

	return user, a.UserRepository.Create(user)
}

func (a *Auth) SignIn(form *dto.AuthForm) (*model.UserWithTokens, error) {
	user, err := a.UserRepository.FindByEmail(form.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if !user.IsPasswordValid(form.Password) {
		return nil, errors.New("the password is incorrect")
	}

	accessToken, refreshToken, err := a.Token.GenerateAccessAndRefreshTokens(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	user.RefreshToken = refreshToken
	a.UserRepository.Save(user)

	return &model.UserWithTokens{User: user, AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (a *Auth) SignOut(id uint) error {
	return a.UserRepository.RemoveRefreshToken(id)
}

func (a *Auth) RefreshToken(token string) (*model.UserWithTokens, error) {
	secretKey := config.Env.Secret.RefreshToken
	claims, err := a.Token.VerifyToken(token, secretKey)
	if err != nil {
		return nil, err
	}

	userID := uint(claims["sub"].(float64))
	role := claims["role"].(string)
	accessToken, refreshToken, err := a.Token.GenerateAccessAndRefreshTokens(uint(userID), model.Role(role))
	if err != nil {
		return nil, err
	}

	user, err := a.UserRepository.FindByRefreshToken(token)
	if err != nil {
		return nil, err
	}

	user.RefreshToken = refreshToken
	err = a.UserRepository.Save(user)

	return &model.UserWithTokens{User: user, AccessToken: accessToken, RefreshToken: refreshToken}, err
}

func (a *Auth) Update(id uint, avatar *multipart.FileHeader, form *dto.UpdateProfileForm) (*model.User, error) {
	user, err := a.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if avatar == nil {
		copier.CopyWithOption(user, form, copier.Option{IgnoreEmpty: true})
		err = a.UserRepository.Update(user)
		if err != nil {
			return nil, err
		}

		return user, err
	}

	oldAvatarPath := user.Avatar
	copier.CopyWithOption(user, form, copier.Option{IgnoreEmpty: true})

	avatarPath, err := a.Storage.Save(avatar)
	if err != nil {
		return nil, err
	}

	user.Avatar = avatarPath

	err = a.UserRepository.Update(user)
	if err != nil {
		os.Remove(avatarPath)

		return nil, err
	}

	a.Storage.Remove(oldAvatarPath)
	return user, nil
}
