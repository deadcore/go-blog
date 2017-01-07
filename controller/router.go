package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/deadcore/go-blog/controller/post"
	"github.com/deadcore/go-blog/controller/ping"
	"github.com/deadcore/go-blog/controller/authentication"
	"github.com/deadcore/go-blog/context"
)

func Router(ctx context.ApplicationContext) http.Handler {

	postController := post.Controller(ctx.DaoContext().PostDao())
	pingController := ping.Controller()
	authenticationController := authentication.Controller(ctx.ServiceContext().AuthenticationService())

	router := httprouter.New()

	router.GET("/posts/:id", postController.Get)
	router.POST("/posts", postController.Post)
	router.GET("/posts", postController.List)

	router.GET("/ping", pingController.Get)

	router.POST("/authentication", authenticationController.Post)

	return router
}
