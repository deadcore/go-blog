package service

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/model"
	"github.com/google/uuid"
	"time"
)

type SessionService interface {
	Create(user model.User) model.Session
}

type sessionService struct {
	sessionDao dao.SessionDao
}

func NewSessionService(sessionDao dao.SessionDao) SessionService {
	return &sessionService{
		sessionDao: sessionDao,
	}
}

func (s *sessionService) Create(user model.User) model.Session {

	token, err := uuid.NewRandom()

	if err != nil {
		panic(err)
	}

	session := model.Session{
		UserId: user.Id,
		Token:  token.String(),
		Expiry: time.Now().Add(time.Hour * 2),
	}

	return s.sessionDao.Save(session)
}