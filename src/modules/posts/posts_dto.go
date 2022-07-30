package posts

type GetPostByIdUri struct {
	ID string `uri:"id" binding:"required"`
}

type CreatePostDto struct {
	Caption string `form:"caption"`
}
