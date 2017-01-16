package dao

import "github.com/deadcore/go-blog/model"

type UserDao interface {
	Get(id string) (model.User, error)
	Save(user model.User) model.User
	FindByEmailAndPassword(password string, email string) (model.User, error)
}