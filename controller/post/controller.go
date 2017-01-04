package post

import "github.com/deadcore/go-blog/dao"

type PostController struct {
	postDao dao.PostDao
}

func New(postDao dao.PostDao) *PostController {
	return &PostController{
		postDao: postDao,
	}
}