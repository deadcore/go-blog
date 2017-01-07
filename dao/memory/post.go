package memory

import (
	"errors"
	"github.com/deadcore/go-blog/model"
	"strconv"
	"github.com/deadcore/go-blog/dao"
)

type inMemoryPostDao struct{}

var arr = make([]model.Post, 0)
var c count64

func PostDao() dao.PostDao {
	return &inMemoryPostDao{}
}

func (m *inMemoryPostDao) Get(id string) (model.Post, error) {
	return findOne(func(post model.Post) bool {
		return post.Id == id
	})
}

func (m *inMemoryPostDao) Save(post model.Post) model.Post {
	post.Id = strconv.FormatUint(c.increment(), 10)
	arr = append(arr, post)
	return post
}

func (m *inMemoryPostDao) FindAll() ([]model.Post, error) {
	return arr, nil
}

func findOne(f func(post model.Post) bool) (model.Post, error) {
	for _, v := range arr {
		if f(v) {
			return v, nil
		}
	}
	return model.Post{}, errors.New("NOT FOUND")
}
