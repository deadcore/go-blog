package service

import "github.com/deadcore/go-blog/model"

type AuthenticationService struct {
	UserService    UserService
	SessionService SessionService
}

func (s AuthenticationService) Authenticate(username string, password string) model.Session {
	user, err := s.UserService.validate(username, password)
	if err != nil {
		panic(err)
	}
	return s.SessionService.create(user)
}