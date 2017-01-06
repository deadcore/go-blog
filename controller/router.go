package controller

import (
	"github.com/deadcore/go-blog/dao/mongo"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/deadcore/go-blog/controller/post"
	"github.com/deadcore/go-blog/controller/ping"
	"github.com/deadcore/go-blog/controller/authentication"
)

func Router() http.Handler {

	mongoContext := getMongoContext()
	postDao := &mongo.MongoPostDao{
		Context: mongoContext,
	}

	postController := post.Controller(postDao)
	pingController := ping.Controller()
	authenticationController := authentication.Controller()

	router := httprouter.New()

	router.GET("/posts/:id", postController.Get)
	router.POST("/posts", postController.Post)
	router.GET("/posts", postController.List)

	router.GET("/ping", pingController.Get)

	router.POST("/authentication", authenticationController.Post)

	return router
}

func getMongoContext() mongo.MongoContext {
	return mongo.NewMongoContext("127.0.0.1", "khazix")
}
