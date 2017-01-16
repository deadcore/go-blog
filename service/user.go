package service

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/model"
	"crypto/sha512"
	"encoding/base64"
)

type UserService interface {
	Validate(email string, password string) (model.User, error)
	Register(email string, password string) model.User
	Resolve(session model.Session) (model.User, error)
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
	return s.userDao.FindByEmailAndPassword(computeHash(password), email)
}

func (s userService) Register(email string, password string) model.User {
	user := model.User{
		Email: email,
		Password: string(computeHash(password)),
	}
	return s.userDao.Save(user)
}

func (s userService) Resolve(session model.Session) (model.User, error) {
	return s.userDao.Get(session.UserId)
}

func computeHash(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}