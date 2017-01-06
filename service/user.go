package service

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/model"
	"crypto/sha512"
)

type UserService struct {
	UserDao dao.UserDao
}

func (s UserService) validate(email string, password string) (model.User, error) {
	hasedPassword := sha512.Sum512([]byte(password))
	return s.UserDao.FindByEmailAndPassword(hasedPassword, email)
}