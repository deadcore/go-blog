package post

import (
	"github.com/deadcore/go-blog/dao"
)

type postController struct {
	postDao dao.PostDao
}

func Controller(postDao dao.PostDao) *postController {
	return &postController{
		postDao: postDao,
	}
}