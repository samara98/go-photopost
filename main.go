package main

import (
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", Greet)

	// version 1
	apiV1 := r.Group("v1")

	// routes
	posts.PostRoutesV1(apiV1)
	users.UsersRoutesV1(apiV1)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"statusCode": http.StatusNotFound, "message": "Not Found"})
	})

	r.Run()
}
