package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

type User struct {
	ID       string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Password string `gorm:"not null"`
	Name     string `gorm:"size:60;not null"`
	Email    string `gorm:"size:60;not null"`
}

// oviously other details like address, favioustes and all model will there

func (u *UserModel) AuthenticateUser(userName, password string) (*User, error) {
	var user User

	if err := u.DB.Where("userName= ?", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Invalid Credential")
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userName)); err != nil {
		return nil, errors.New("Invalid Credentail")
	}

	return &user, nil

}

func (u *UserModel) GetUserById(id string) (*User, error) {

	var user User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errMsg := "User Not Found"
			return nil, errors.New(errMsg)
		}

		errMsg := "Internal Server Error"
		return nil, errors.New(errMsg)
	}

	return &user, nil

}
