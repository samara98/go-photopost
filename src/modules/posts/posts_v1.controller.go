package posts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPostListV1(c *gin.Context) {
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
	c.JSON(http.StatusOK, result)
}
