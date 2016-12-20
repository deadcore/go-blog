package api

import (
	"net/http"
	"api.khazix.co.uk/api/controller"
	"api.khazix.co.uk/api/filter"
	"strconv"
	"bytes"
	logger "github.com/Sirupsen/logrus"
)

type Server struct {
	Port uint64
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
	var strPort = strconv.FormatUint(s.Port, 10)
	buffer.WriteString("127.0.0.1:")
	buffer.WriteString(strPort)
	return buffer.String()
}