package memory

import (
	"api.khazix.co.uk/model"
	"sync/atomic"
	"github.com/kataras/go-errors"
)

type (
	InMemoryPostDao struct{}
)

var arr = make([]model.Post, 16)

type count32 uint64

func (c *count32) increment() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

func (c *count32) get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}
var c count32

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
			return v, nil;
		}
	}
	return model.Post{}, errors.New("NOT FOUND")
}