package repository

import (
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func (u *User) Create(user *model.User) error {
	return u.DB.Create(user).Error
}

func (u *User) Save(user *model.User) error {
	return u.DB.Save(user).Error
}

func (u *User) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (u *User) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := u.DB.First(&user, id).Error

	return &user, err
}

func (u *User) FindByRefreshToken(refreshToken string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("refresh_token = ?", refreshToken).First(&user).Error

	return &user, err
}

func (u *User) Update(user *model.User) error {
	return u.DB.Save(user).Error
}

func (u *User) RemoveRefreshToken(id uint) error {
	return u.DB.Where("id = ?", id).Update("refresh_token", nil).Error
}
