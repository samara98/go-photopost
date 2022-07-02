package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostsControllerV1Interface interface {
	Run(router *gin.RouterGroup)
	GetPostList(c *gin.Context)
}

type PostsControllerV1 struct {
	postsServiceV1 PostsServiceV1Interface
}

func NewPostsControllerV1(
	postsServiceV1 *PostsServiceV1,
) *PostsControllerV1 {
	return &PostsControllerV1{
		postsServiceV1,
	}
}

func (pc PostsControllerV1) Run(router *gin.RouterGroup) {
	router.GET("/", pc.GetPostList)
	router.GET("/:id", pc.GetPost)
}

func (pc PostsControllerV1) GetPostList(c *gin.Context) {
	result := pc.postsServiceV1.GetPostList()
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetPost(c *gin.Context) {
	result := pc.postsServiceV1.GetPost()
	c.JSON(http.StatusOK, result)
}
