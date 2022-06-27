package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsersListV1(c *gin.Context) {
	result := []gin.H{
		{
			"id":        1,
			"createdAt": time.Now(),
			"updatedAt": time.Now(),
			"name":      "noname",
			"email":     "noname@mailsac.com",
			"birthdate": time.Now(),
		},
		{
			"id":        2,
			"createdAt": time.Now(),
			"updatedAt": time.Now(),
			"name":      "nn",
			"email":     "nn@mailsac.com",
			"birthdate": time.Now(),
		},
	}
	c.JSON(http.StatusOK, result)
}
