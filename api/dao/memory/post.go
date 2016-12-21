package memory

import (
	"errors"
	"github.com/deadcore/go-blog/api/model"
)

type (
	InMemoryPostDao struct {
	}
)

var (
	arr = make([]model.Post, 16)
	c   count64
)

func (m *InMemoryPostDao) Get(id uint64) (model.Post, error) {
	return findOne(func(post model.Post) bool {
		return post.Id == id
	})
}

func (m *InMemoryPostDao) Save(post model.Post) model.Post {
	post.Id = c.increment()
	arr = append(arr, post)
	return post
}

func findOne(f func(post model.Post) bool) (model.Post, error) {
	for _, v := range arr {
		if f(v) {
			return v, nil
		}
	}
	return model.Post{}, errors.New("NOT FOUND")
}
