package posts

import (
	"go-photopost/src/entities"
	"go-photopost/src/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostsControllerV1Interface interface {
	Run(router *gin.RouterGroup)
	CreatePost(c *gin.Context)
	GetPostList(c *gin.Context)
	GetPost(c *gin.Context)
}

type PostsControllerV1 struct {
	Log               *log.Logger
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	PostsServiceV1    PostsServiceV1Interface
}

func NewPostsControllerV1(
	log *log.Logger,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	postsServiceV1 *PostsServiceV1,
) *PostsControllerV1 {
	return &PostsControllerV1{
		log,
		jwtAuthMiddleware,
		postsServiceV1,
	}
}

func (pc PostsControllerV1) Run(router *gin.RouterGroup) {
	router.POST("/", pc.JWTAuthMiddleware.Handler(), pc.CreatePost)
	router.GET("/", pc.GetPostList)
	router.GET("/p/:id", pc.GetPost)
}

func (pc PostsControllerV1) CreatePost(c *gin.Context) {
	var body CreatePostDto
	c.Bind(&body)

	userAny, _ := c.Get("user")
	user := userAny.(*entities.User)

	result := pc.PostsServiceV1.CreatePost(user, &body)
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetPostList(c *gin.Context) {
	result := pc.PostsServiceV1.GetPostList()
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetPost(c *gin.Context) {
	var uri GetPostByIdUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := pc.PostsServiceV1.GetPost(&uri)
	c.JSON(http.StatusOK, result)
}
