package server

import (
	"bytes"
	logger "github.com/Sirupsen/logrus"
	"github.com/deadcore/go-blog/controller"
	"github.com/deadcore/go-blog/filter"
	"net/http"
	"github.com/deadcore/go-blog/context"
)

type Instance struct {
	Port string
}

func (s Instance) Start() error {

	applicationContext := context.NewApplicationContext()

	controllers := controller.Router(applicationContext)

	jsonHeaderSettingFilter := filter.JsonContentTypeHandler(controllers)
	corsHandler := filter.CorsHandler(jsonHeaderSettingFilter)
	loggingHandler := filter.LoggingHandler(corsHandler)

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
