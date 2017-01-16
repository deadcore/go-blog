package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/deadcore/go-blog/controller/post"
	"github.com/deadcore/go-blog/controller/ping"
	"github.com/deadcore/go-blog/controller/authentication"
	"github.com/deadcore/go-blog/controller/authentication/register"
	"github.com/deadcore/go-blog/context"
)

type Router interface {
	Build() http.Handler
}

type router struct {
	context context.ApplicationContext
}

func NewRouter(context context.ApplicationContext) Router {
	return &router{
		context: context,
	}
}

func (r router) Build() http.Handler {

	postController := post.Controller(r.context.DaoContext().PostDao())
	pingController := ping.Controller()
	authenticationController := authentication.Controller(r.context.ServiceContext().AuthenticationService(), r.context.ServiceContext().SessionService())
	registerController := register.Controller(r.context.ServiceContext().UserService())

	router := httprouter.New()

	router.GET("/posts/:id", postController.Get)
	router.POST("/posts", r.authenticatedRoute("ADMIN", postController.Post))
	router.GET("/posts", postController.List)

	router.GET("/ping", pingController.Get)

	router.POST("/authentication", authenticationController.Post)
	router.DELETE("/authentication/:token", authenticationController.Delete)
	router.POST("/authentication/register", registerController.Post)

	return router
}

func (ro router) authenticatedRoute(role string, h httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		sessionToken := request.Header.Get("x-session-token")
		if sessionToken == "" {
			http.Error(writer, "", 401)
			return;
		}
		session, err := ro.context.ServiceContext().SessionService().Resolve(sessionToken)

		if err != nil {
			http.Error(writer, "", 401)
			return;
		}

		user, err := ro.context.ServiceContext().UserService().Resolve(session)

		if !contains(user.Roles, role) {
			http.Error(writer, "", 403)
			return;
		}

		h(writer, request, params)
	}
}

func contains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}