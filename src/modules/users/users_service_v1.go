package users

import (
	"go-photopost/src/entities"
	"go-photopost/src/helpers"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersServiceV1Interface interface {
	CreateUser(body CreateUserDto) entities.User
	GetUserList() []gin.H
	GetUser() entities.User
}

type UsersServiceV1 struct {
	DB *gorm.DB
}

func NewUsersServiceV1(
	db *gorm.DB,
) *UsersServiceV1 {
	return &UsersServiceV1{
		DB: db,
	}
}

func (us UsersServiceV1) CreateUser(body CreateUserDto) entities.User {
	hashedPassword := helpers.HashPassword([]byte(body.Password))

	user := entities.User{
		Email:     &body.Email,
		Username:  &body.Username,
		Password:  string(hashedPassword),
		Name:      body.Name,
		Birthdate: nil,
	}
	us.DB.Create(&user)

	return user
}

func (us UsersServiceV1) GetUserList() []gin.H {
	result := []gin.H{
		{
			"id":        1,
			"createdAt": time.Now(),
			"updatedAt": time.Now(),
			"name":      "Sulthon Abdul Malik",
			"email":     "sulthon@mailsac.com",
			"birthdate": nil,
		},
		{
			"id":        1,
			"createdAt": time.Now(),
			"updatedAt": time.Now(),
			"name":      "Samara 98",
			"email":     "samara98@mailsac.com",
			"birthdate": nil,
		},
	}

	return result
}

func (us UsersServiceV1) GetUser() entities.User {
	// result := gin.H{
	// 	"id":        1,
	// 	"createdAt": time.Now(),
	// 	"updatedAt": time.Now(),
	// 	"name":      "Sulthon Abdul Malik",
	// 	"email":     "sulthon@mailsac.com",
	// 	"birthdate": nil,
	// }

	var user entities.User
	us.DB.First(&user, 1)

	return user
}
