package service

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/model"
	"github.com/google/uuid"
	"time"
)

type SessionService struct {
	SessionDao dao.SessionDao
}

func (s *SessionService) create(user model.User) model.Session {

	token, err := uuid.NewRandom()

	if err != nil {
		panic(err)
	}

	session := model.Session{
		UserId: user.Id,
		Token:  token.String(),
		Expiry: time.Now().Add(time.Hour * 2),
	}

	return s.SessionDao.Save(session)
}