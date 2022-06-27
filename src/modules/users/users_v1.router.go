package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRoutesV1(routerGroup *gin.RouterGroup) {
	postsV1 := routerGroup.Group("users")

	postsV1.GET("/", GetUsersListV1)
}
