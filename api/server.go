package api

import (
	"bytes"
	logger "github.com/Sirupsen/logrus"
	"github.com/deadcore/go-blog/api/controller"
	"github.com/deadcore/go-blog/api/filter"
	"net/http"
)

type Server struct {
	Port string
}

func (s Server) Start() error {
	controllers := controller.Router()

	loggingHandler := filter.LoggingHandler(controllers)

	bindAddress := s.bindAddress()

	logger.Info("Starting api.khhazix.co.uk on ", bindAddress)

	return http.ListenAndServe(bindAddress, loggingHandler)
}

func (s Server) bindAddress() string {
	var buffer bytes.Buffer
	buffer.WriteString("127.0.0.1:")
	buffer.WriteString(s.Port)
	return buffer.String()
}
