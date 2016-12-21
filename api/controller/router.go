package controller

import (
	"github.com/deadcore/go-blog/api/dao/memory"
	"github.com/deadcore/go-blog/api/dao/mongo"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Router() http.Handler {

	//var mongoContext = getMongoContext()
	//
	//var mongoPostDao = new(mongo.MongoPostDao)
	//
	//mongoPostDao.Context = mongoContext

	postController := PostController{
		postDao: new(memory.InMemoryPostDao),
	}

	pingController := PingController{}

	router := httprouter.New()

	router.GET("/post/:id", postController.Get)
	router.POST("/post", postController.Post)

	router.GET("/ping", pingController.Get)

	return router
}

func getMongoContext() mongo.MongoContext {
	return mongo.NewMongoContext("127.0.0.1", "khazix")
}
