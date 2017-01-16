package authentication

import "github.com/deadcore/go-blog/service"

type authenticationController struct {
	authenticationService service.AuthenticationService
	sessionService        service.SessionService
}

func Controller(authenticationService service.AuthenticationService, sessionService service.SessionService) *authenticationController {
	return &authenticationController{
		authenticationService: authenticationService,
		sessionService: sessionService,
	}
}