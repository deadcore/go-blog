package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"api.khazix.co.uk/dao/memory"
)

func Router() http.Handler {

	var postController = PostController {
		postDao: new(memory.InMemoryPostDao),
	}

	router := httprouter.New()

	router.GET("/post/:id", postController.Get)
	router.POST("/post", postController.Post)

	return router
}
