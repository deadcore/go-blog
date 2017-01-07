package service

import "github.com/deadcore/go-blog/model"

type AuthenticationService interface {
	Authenticate(username string, password string) model.Session
}

type authenticationService struct {
	userService    UserService
	sessionService SessionService
}

func NewAuthenticationService(userService UserService, sessionService SessionService) AuthenticationService {
	return &authenticationService{
		userService: userService,
		sessionService: sessionService,
	}
}

func (s authenticationService) Authenticate(username string, password string) model.Session {
	user, err := s.userService.Validate(username, password)
	if err != nil {
		panic(err)
	}
	return s.sessionService.Create(user)
}