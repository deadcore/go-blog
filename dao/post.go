package dao

import (
	"github.com/deadcore/go-blog/model"
)

type PostDao interface {
	Get(id string) (model.Post, error)
	Save(post model.Post) model.Post
	FindAll() ([]model.Post, error)
}