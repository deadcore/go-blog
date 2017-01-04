package server

import (
	"bytes"
	logger "github.com/Sirupsen/logrus"
	"github.com/deadcore/go-blog/controller"
	"github.com/deadcore/go-blog/filter"
	"net/http"
)

type Instance struct {
	Port string
}

func (s Instance) Start() error {
	controllers := controller.Router()

	jsonHeaderSettingFilter := filter.JsonContentTypeHandler(controllers)
	loggingHandler := filter.LoggingHandler(jsonHeaderSettingFilter)

	bindAddress := s.bindAddress()

	logger.Info("Starting api.khhazix.co.uk on ", bindAddress)

	return http.ListenAndServe(bindAddress, loggingHandler)
}

func (s Instance) bindAddress() string {
	var buffer bytes.Buffer
	buffer.WriteString("127.0.0.1:")
	buffer.WriteString(s.Port)
	return buffer.String()
}
