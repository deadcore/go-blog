package dao

import (
	"api.khazix.co.uk/api/model"
)

type EntityNotFoundError struct {}

type PostDao interface {
	Get(id string) (model.Post, error)
	Save(post model.Post) model.Post
}