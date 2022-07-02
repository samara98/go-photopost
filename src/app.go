package src

import (
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

type App struct {
	UsersModule *users.UsersModule
	PostsModule *posts.PostsModule
}

func NewApp(
	usersModule *users.UsersModule,
	postsModule *posts.PostsModule,
) *App {
	return &App{
		usersModule,
		postsModule,
	}
}

func (app App) Start() {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	r.GET("/", greet)

	// version 1
	apiV1 := r.Group("v1")

	// routes
	app.UsersModule.Router(apiV1)
	app.PostsModule.Router(apiV1)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"statusCode": http.StatusNotFound, "message": "Not Found"})
	})

	r.Run()
}

func greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
