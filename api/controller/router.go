package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/deadcore/go-blog/api/dao/mongo"
)

func Router() http.Handler {

	mongoContext := mongo.NewMongoContext("127.0.0.1", "khazix")

	var mongoPostDao = new(mongo.MongoPostDao)

	mongoPostDao.Context = mongoContext

	postController := PostController{
		postDao: mongoPostDao,
	}

	pingController := PingController{}

	router := httprouter.New()

	router.GET("/post/:id", postController.Get)
	router.POST("/post", postController.Post)

	router.GET("/ping", pingController.Get)

	return router
}