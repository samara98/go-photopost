package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersControllerV1Interface interface {
	Run(router *gin.RouterGroup)
}

type UsersControllerV1 struct {
	UsersService UsersServiceV1Interface
}

func NewUsersControllerV1(
	usersService *UsersServiceV1,
) *UsersControllerV1 {
	return &UsersControllerV1{
		usersService,
	}
}

func (uc UsersControllerV1) Run(router *gin.RouterGroup) {
	router.POST("/", uc.CreateUser)
	router.GET("/", uc.GetUserList)
	router.GET("/:userId", uc.GetUser)
}

func (uc UsersControllerV1) CreateUser(c *gin.Context) {
	var body CreateUserDto
	c.Bind(&body)

	result := uc.UsersService.CreateUser(body)

	c.JSON(http.StatusCreated, result)
}

func (uc UsersControllerV1) GetUserList(c *gin.Context) {
	result := uc.UsersService.GetUserList()
	c.JSON(http.StatusOK, result)
}

func (uc UsersControllerV1) GetUser(c *gin.Context) {
	result := uc.UsersService.GetUser()
	c.JSON(http.StatusOK, result)
}
