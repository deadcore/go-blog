package context

import "github.com/deadcore/go-blog/service"

type ServiceContext interface {
	UserService() service.UserService
	SessionService() service.SessionService
	AuthenticationService() service.AuthenticationService
}

type serviceContext struct {
	userService           service.UserService
	sessionService        service.SessionService
	authenticationService service.AuthenticationService
}

func (s serviceContext) UserService() service.UserService {
	return s.userService
}

func (s serviceContext) SessionService() service.SessionService {
	return s.sessionService
}

func (s serviceContext) AuthenticationService() service.AuthenticationService {
	return s.authenticationService
}

func NewServiceContext(daoContext DaoContext) ServiceContext {
	userService := service.NewUserService(daoContext.UserDao())
	sessionService := service.NewSessionService(daoContext.SessionDao(), daoContext.UserDao())
	authenticationService := service.NewAuthenticationService(userService, sessionService)

	return &serviceContext{
		userService: userService,
		sessionService: sessionService,
		authenticationService: authenticationService,
	}

}