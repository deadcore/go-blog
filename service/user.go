package service

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/model"
	"crypto/sha512"
)

type UserService interface {
	Validate(email string, password string) (model.User, error)
}

type userService struct {
	userDao dao.UserDao
}

func NewUserService(userDao dao.UserDao) UserService {
	return &userService{
		userDao: userDao,
	}
}

func (s userService) Validate(email string, password string) (model.User, error) {
	hasedPassword := sha512.Sum512([]byte(password))
	return s.userDao.FindByEmailAndPassword(hasedPassword, email)
}