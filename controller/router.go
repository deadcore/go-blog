package controller

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/dao/memory"
	"github.com/deadcore/go-blog/dao/mongo"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Router() http.Handler {

	var postDao dao.PostDao = &memory.InMemoryPostDao{}
	var postController = PostController{
		postDao: postDao,
	}
	var pingController PingController

	router := httprouter.New()

	router.GET("/post/:id", postController.Get)
	router.POST("/post", postController.Post)

	router.GET("/ping", pingController.Get)

	return router
}

func getMongoContext() mongo.MongoContext {
	return mongo.NewMongoContext("127.0.0.1", "khazix")
}
