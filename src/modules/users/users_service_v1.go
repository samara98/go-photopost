package users

import (
	"go-photopost/src/entities"
	"go-photopost/src/helpers"
	"log"

	"gorm.io/gorm"
)

type UsersServiceV1Interface interface {
	CreateUser(body CreateUserDto) *entities.User
	GetUserList() []entities.User
	GetUser(uri *GetUserByIdUri) *entities.User
}

type UsersServiceV1 struct {
	Log *log.Logger
	DB  *gorm.DB
}

func NewUsersServiceV1(
	log *log.Logger,
	db *gorm.DB,
) *UsersServiceV1 {
	return &UsersServiceV1{
		Log: log,
		DB:  db,
	}
}

func (us UsersServiceV1) CreateUser(body CreateUserDto) *entities.User {
	hashedPassword := helpers.HashPassword([]byte(body.Password))

	user := entities.User{
		Email:     &body.Email,
		Username:  &body.Username,
		Password:  string(hashedPassword),
		Name:      body.Name,
		Birthdate: nil,
	}
	us.DB.Create(&user)

	return &user
}

func (us UsersServiceV1) GetUserList() []entities.User {
	var users []entities.User
	us.DB.Find(&users)

	return users
}

func (us UsersServiceV1) GetUser(uri *GetUserByIdUri) *entities.User {
	var user entities.User
	us.DB.First(&user, uri.ID)

	return &user
}
