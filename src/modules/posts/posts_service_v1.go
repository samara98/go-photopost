package posts

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostsServiceV1Interface interface {
	GetPostList() []gin.H
	GetPost() gin.H
}

type PostsServiceV1 struct {
	DB *gorm.DB
}

func NewPostsServiceV1(
	db *gorm.DB,
) *PostsServiceV1 {
	return &PostsServiceV1{
		DB: db,
	}
}

func (ps PostsServiceV1) GetPostList() []gin.H {
	result := []gin.H{
		{
			"id":          1,
			"createdAt":   time.Now(),
			"updatedAt":   time.Now(),
			"authorId":    "noname",
			"caption":     "nn",
			"isPublished": true,
		},
		{
			"id":          2,
			"createdAt":   time.Now(),
			"updatedAt":   time.Now(),
			"authorId":    "noname",
			"caption":     "nn",
			"isPublished": true,
		},
	}

	return result
}

func (ps PostsServiceV1) GetPost() gin.H {
	result := gin.H{
		"id":          1,
		"createdAt":   time.Now(),
		"updatedAt":   time.Now(),
		"authorId":    "noname",
		"caption":     "nn",
		"isPublished": true,
	}

	return result
}
