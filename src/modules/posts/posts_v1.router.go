package posts

import "github.com/gin-gonic/gin"

func PostRoutesV1(routerGroup *gin.RouterGroup) {
	postsV1 := routerGroup.Group("posts")

	postsV1.GET("/", GetPostListV1)
}
