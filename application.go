package main

import (
	logger "github.com/Sirupsen/logrus"
	"os"
	"github.com/deadcore/go-blog/server"
)

func main() {
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "5000"
	}

	configLogging()

	startApi(port)
}

func startApi(port string) {
	var apiServer = server.Instance{
		Port: port,
	}

	err := apiServer.Start()

	logger.Panic(err)
}

func configLogging() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logger.TextFormatter{})

	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logger.InfoLevel)
}
