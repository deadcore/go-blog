package main

import (
	logger "github.com/Sirupsen/logrus"
	"os"
	"github.com/deadcore/go-blog/server"
	"github.com/deadcore/go-blog/context"
)

func main() {
	configLogging()

	var config = context.Build()

	var apiServer = server.Instance{
		Configuration: *config,
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
