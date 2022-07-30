package users

type GetUserByIdUri struct {
	ID string `uri:"userId" binding:"required"`
}

type CreateUserDto struct {
	Email    string `form:"name"`
	Username string `form:"username"`
	Password string `form:"password"`
	Name     string `form:"name"`
}
