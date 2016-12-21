package main

import (
	logger "github.com/Sirupsen/logrus"
	"github.com/deadcore/go-blog/api"
	"os"
)

func main() {
	configLogging()

	port, present := os.LookupEnv("PORT")
	if !present {
		port = "5000"
	}

	var server = api.Server{
		Port: port,
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
