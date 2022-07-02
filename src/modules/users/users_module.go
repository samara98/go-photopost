package users

import "github.com/gin-gonic/gin"

type UsersModuleInterface interface {
	Router(rg *gin.RouterGroup)
}

type UsersModule struct {
	UsersControllerV1 UsersControllerV1Interface
}

func NewUsersModule(
	usersControllerV1 *UsersControllerV1,
) *UsersModule {
	return &UsersModule{
		usersControllerV1,
	}
}

func (um UsersModule) Router(rg *gin.RouterGroup) {
	usersRoutesV1 := rg.Group("users")

	um.UsersControllerV1.Run(usersRoutesV1)
}
