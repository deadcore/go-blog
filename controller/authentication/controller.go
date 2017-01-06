package authentication

import "github.com/deadcore/go-blog/service"

type authenticationController struct {
	authenticationService service.AuthenticationService
}

func Controller(authenticationService service.AuthenticationService) *authenticationController {
	return &authenticationController{
		authenticationService: authenticationService,
	}
}