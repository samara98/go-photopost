package users

import (
	"go-photopost/src/entities"
	"go-photopost/src/helpers"

	"gorm.io/gorm"
)

type UsersServiceV1Interface interface {
	CreateUser(body CreateUserDto) entities.User
	GetUserList() []entities.User
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

func (us UsersServiceV1) GetUserList() []entities.User {
	var users []entities.User
	us.DB.Find(&users)

	return users
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
