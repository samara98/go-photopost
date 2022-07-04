//go:build wireinject
// +build wireinject

package main

import (
	"go-photopost/src"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"

	"github.com/google/wire"
)

func InitApp() *src.App {
	wire.Build(
		lib.NewDatabase,
		lib.NewJWTAuthHelper,
		middlewares.NewJWTAuthMiddleware,
		usersSvcV1,
		usersCtlV1,
		usersModule,
		postsSvcV1,
		postsCtlV1,
		postsModule,
		src.NewApp,
	)
	return &src.App{}
}

var usersModule = wire.NewSet(
	users.NewUsersModule,
	wire.Bind(
		new(users.UsersModuleInterface),
		new(*users.UsersModule),
	),
)

var usersCtlV1 = wire.NewSet(
	users.NewUsersControllerV1,
	wire.Bind(
		new(users.UsersControllerV1Interface),
		new(*users.UsersControllerV1),
	),
)

var usersSvcV1 = wire.NewSet(
	users.NewUsersServiceV1,
	wire.Bind(
		new(users.UsersServiceV1Interface),
		new(*users.UsersServiceV1),
	),
)

var postsModule = wire.NewSet(
	posts.NewPostsModule,
	wire.Bind(
		new(posts.PostsModuleInterface),
		new(*posts.PostsModule),
	),
)

var postsCtlV1 = wire.NewSet(
	posts.NewPostsControllerV1,
	wire.Bind(
		new(posts.PostsControllerV1Interface),
		new(*posts.PostsControllerV1),
	),
)

var postsSvcV1 = wire.NewSet(
	posts.NewPostsServiceV1,
	wire.Bind(
		new(posts.PostsServiceV1Interface),
		new(*posts.PostsServiceV1),
	),
)
