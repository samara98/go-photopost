package posts

import "github.com/gin-gonic/gin"

type PostsModuleInterface interface {
	Router(rg *gin.RouterGroup)
}

type PostsModule struct {
	PostsControllerV1 PostsControllerV1Interface
}

func NewPostsModule(
	postsControllerV1 *PostsControllerV1,
) *PostsModule {
	return &PostsModule{
		postsControllerV1,
	}
}

func (pc PostsModule) Router(rg *gin.RouterGroup) {
	postsRoutesV1 := rg.Group("posts")

	pc.PostsControllerV1.Run(postsRoutesV1)
}
