package middlewares

import (
	"go-photopost/src/lib"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
	Log           *log.Logger
	JWTAuthHelper *lib.JWTAuthHelper
}

func NewJWTAuthMiddleware(
	log *log.Logger,
	jwtHelper *lib.JWTAuthHelper,
) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{
		log,
		jwtHelper,
	}
}

func (m JWTAuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		c.Set("user", gin.H{
			"id": 1,
		})
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := m.JWTAuthHelper.Authorize(authToken)
			if authorized {
				c.Next()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			m.Log.Println(err)
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not authorized",
		})
		c.Abort()
	}
}
