package main

import (
	"bytes"
	logger "github.com/Sirupsen/logrus"
	"github.com/deadcore/go-blog/controller"
	"github.com/deadcore/go-blog/filter"
	"net/http"
)

type Instrance struct {
	Port string
}

func (s Instrance) Start() error {
	controllers := controller.Router()

	loggingHandler := filter.LoggingHandler(controllers)

	bindAddress := s.bindAddress()

	logger.Info("Starting api.khhazix.co.uk on ", bindAddress)

	return http.ListenAndServe(bindAddress, loggingHandler)
}

func (s Instrance) bindAddress() string {
	var buffer bytes.Buffer
	buffer.WriteString("127.0.0.1:")
	buffer.WriteString(s.Port)
	return buffer.String()
}
