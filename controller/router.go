package controller

import (
	"github.com/deadcore/go-blog/dao/memory"
	"github.com/deadcore/go-blog/dao/mongo"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/deadcore/go-blog/controller/post"
	"github.com/deadcore/go-blog/controller/ping"
)

func Router() http.Handler {

	postDao := &memory.InMemoryPostDao{}

	postController := post.New(postDao)
	pingController := ping.New()

	router := httprouter.New()

	router.GET("/posts/:id", postController.Get)
	router.POST("/posts", postController.Post)
	router.GET("/posts", postController.List)

	router.GET("/ping", pingController.Get)

	return router
}

func getMongoContext() mongo.MongoContext {
	return mongo.NewMongoContext("127.0.0.1", "khazix")
}
