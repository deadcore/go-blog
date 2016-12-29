package controller

import (
	"fmt"
	"github.com/deadcore/go-blog/api/dao"
	"github.com/deadcore/go-blog/api/dao/memory"
	"github.com/deadcore/go-blog/api/dao/mongo"
	"github.com/facebookgo/inject"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func Router() http.Handler {

	var g inject.Graph

	var postController PostController
	var pingController PingController
	var postDao dao.PostDao = &memory.InMemoryPostDao{}

	postDao.Get("1")

	err := g.Provide(
		&inject.Object{Name: "postDao", Value: postDao},
		&inject.Object{Value: &postController},
		&inject.Object{Value: &pingController},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	router := httprouter.New()

	router.GET("/post/:id", postController.Get)
	router.POST("/post", postController.Post)

	router.GET("/ping", pingController.Get)

	return router
}

func getMongoContext() mongo.MongoContext {
	return mongo.NewMongoContext("127.0.0.1", "khazix")
}
