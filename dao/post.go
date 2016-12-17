package dao

import (
	"api.khazix.co.uk/model"
)

type EntityNotFoundError struct {}

type PostDao interface {
	Get(id uint64) (model.Post, error)
	Save(post model.Post) model.Post
}