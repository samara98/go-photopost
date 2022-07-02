package users

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UsersServiceV1Interface interface {
	GetUserList() []gin.H
	GetUser() gin.H
}

type UsersServiceV1 struct {
}

func NewUsersServiceV1() *UsersServiceV1 {
	return &UsersServiceV1{}
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

func (us UsersServiceV1) GetUser() gin.H {
	result := gin.H{
		"id":        1,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
		"name":      "Sulthon Abdul Malik",
		"email":     "sulthon@mailsac.com",
		"birthdate": nil,
	}

	return result
}
