package main

import (
	logger "github.com/Sirupsen/logrus"
	"github.com/deadcore/go-blog/api"
	"os"
)

func main() {
	configLogging()

	var server = api.Server{
		Port: 3000,
	}

	err := server.Start()

	logger.Panic(err)
}

func configLogging() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logger.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	logger.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	logger.SetLevel(logger.InfoLevel)
}
