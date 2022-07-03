package users

type CreateUserDto struct {
	Email    string `form:"name"`
	Username string `form:"username"`
	Password string `form:"password"`
	Name     string `form:"name"`
}
