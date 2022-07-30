package posts

import (
	"go-photopost/src/entities"
	"log"

	"gorm.io/gorm"
)

type PostsServiceV1Interface interface {
	CreatePost(user *entities.User, createPostDto *CreatePostDto) *entities.Post
	GetPostList() []entities.Post
	GetPost(uri *GetPostByIdUri) *entities.Post
}

type PostsServiceV1 struct {
	Log *log.Logger
	DB  *gorm.DB
}

func NewPostsServiceV1(
	log *log.Logger,
	db *gorm.DB,
) *PostsServiceV1 {
	return &PostsServiceV1{
		Log: log,
		DB:  db,
	}
}

func (ps PostsServiceV1) CreatePost(user *entities.User, createPostDto *CreatePostDto) *entities.Post {
	newPost := &entities.Post{
		AuthorID: user.ID,
		Caption:  createPostDto.Caption,
	}
	ps.DB.Create(newPost)

	return newPost
}

func (ps PostsServiceV1) GetPostList() []entities.Post {
	var postList []entities.Post
	ps.DB.Find(&postList)

	return postList
}

func (ps PostsServiceV1) GetPost(uri *GetPostByIdUri) *entities.Post {
	var post entities.Post
	ps.DB.Find(&post, uri.ID).First(&post)

	return &post
}
