package service

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/model"
	"github.com/google/uuid"
	"time"
)

type SessionService interface {
	Create(user model.User) model.Session
	Resolve(sessionToken string) (model.Session, error)
	Delete(sessionToken string) error
}

type sessionService struct {
	sessionDao dao.SessionDao
	userDao    dao.UserDao
}

func NewSessionService(sessionDao dao.SessionDao, userDao dao.UserDao) SessionService {
	return &sessionService{
		sessionDao: sessionDao,
		userDao: userDao,
	}
}

func (s *sessionService) Create(user model.User) model.Session {

	token, err := uuid.NewRandom(); if err != nil {
		panic(err)
	}

	session := model.Session{
		UserId: user.Id,
		Token:  token.String(),
		Expiry: time.Now().Add(time.Hour * 2),
	}

	s.sessionDao.Save(&session)

	return session
}

func (s *sessionService) Resolve(sessionToken string) (model.Session, error) {
	return s.sessionDao.FindBySessionToken(sessionToken)
}

func (s *sessionService) Delete(sessionToken string) error {
	session, err := s.sessionDao.FindBySessionToken(sessionToken)
	if err != nil {
		return err
	}
	return s.sessionDao.Delete(session.Id)
}