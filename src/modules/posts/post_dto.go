package posts

type CreatePostDto struct {
	Caption string `form:"caption"`
}

type GetPostByIdUri struct {
	ID string `uri:"id" binding:"required"`
}
