package register

import "github.com/deadcore/go-blog/service"

type registrationController struct {
	userService service.UserService
}

func Controller(userService service.UserService) *registrationController {
	return &registrationController{
		userService: userService,
	}
}