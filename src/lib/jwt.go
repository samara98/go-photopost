package lib

import (
	"errors"
	"go-photopost/src/entities"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

// JWTAuthHelper service relating to authorization
type JWTAuthHelper struct {
	Log *log.Logger
}

// NewJWTAuthHelper creates a new auth service
func NewJWTAuthHelper() *JWTAuthHelper {
	return &JWTAuthHelper{}
}

// CreateToken creates jwt auth token
func (s JWTAuthHelper) CreateToken(user entities.User) *Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"name":     user.Name,
		"email":    *user.Email,
		"username": *user.Username,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		s.Log.Println("JWT validation failed: ", err)
	}

	return &Token{
		Type:  "Bearer",
		Token: tokenString,
	}
}

// Authorize authorizes the generated token
func (s JWTAuthHelper) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}
	return false, errors.New("couldn't handle token")
}
