package src

import (
	"errors"
	"go-photopost/src/entities"
	"go-photopost/src/helpers"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"gorm.io/gorm"
)

type App struct {
	DB                *gorm.DB
	JWTAuthHelper     *lib.JWTAuthHelper
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	UsersModule       *users.UsersModule
	PostsModule       *posts.PostsModule
}

func NewApp(
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	usersModule *users.UsersModule,
	postsModule *posts.PostsModule,
) *App {
	return &App{
		db,
		jwtAuthHelper,
		jwtAuthMiddleware,
		usersModule,
		postsModule,
	}
}

func (app App) Start() {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	r.GET("/", app.greet)
	r.POST("/register", app.register)
	r.POST("/login", app.login)
	r.GET("/me", app.JWTAuthMiddleware.Handler(), app.me)

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

func (app App) greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

type RegisterUserDto struct {
	Email    string `form:"name"`
	Username string `form:"username"`
	Password string `form:"password"`
	Name     string `form:"name"`
}

func (app App) register(c *gin.Context) {
	var body RegisterUserDto
	c.Bind(&body)

	hashedPassword := helpers.HashPassword([]byte(body.Password))

	user := entities.User{
		Email:    &body.Email,
		Username: &body.Username,
		Password: string(hashedPassword),
		Name:     body.Name,
	}
	app.DB.Create(&user)

	c.JSON(http.StatusCreated, user)
}

type LoginUserDto struct {
	UserSession string `form:"userSession"`
	Password    string `form:"password"`
}

func (app App) login(c *gin.Context) {
	var body LoginUserDto
	c.Bind(&body)

	var user entities.User
	res := app.DB.Where("email = ? OR username = ?", body.UserSession, body.UserSession).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "Email/Username or Password",
		})
		return
	}

	err := helpers.CompareHash([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "Email/Username or Password",
		})
		return
	}

	// create token
	token := app.JWTAuthHelper.CreateToken(user)

	c.JSON(http.StatusCreated, token)
}

func (app App) me(c *gin.Context) {
	user, _ := c.Get("user")

	log.Default().Println(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "authorized",
	})
}
