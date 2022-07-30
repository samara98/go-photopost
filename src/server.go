package src

import (
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"gorm.io/gorm"
)

type Server struct {
	DB                *gorm.DB
	JWTAuthHelper     *lib.JWTAuthHelper
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	AppModule         *AppModule
	UsersModule       *users.UsersModule
	PostsModule       *posts.PostsModule
}

func NewServer(
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	appModule *AppModule,
	usersModule *users.UsersModule,
	postsModule *posts.PostsModule,
) *Server {
	return &Server{
		db,
		jwtAuthHelper,
		jwtAuthMiddleware,
		appModule,
		usersModule,
		postsModule,
	}
}

func (server Server) Start() {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))

	// version 1
	apiV1 := r.Group("v1")

	// routes
	server.AppModule.Router(r)
	server.UsersModule.Router(apiV1)
	server.PostsModule.Router(apiV1)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"statusCode": http.StatusNotFound, "message": "Not Found"})
	})

	r.Run()
}
