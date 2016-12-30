package main

import (
	logger "github.com/Sirupsen/logrus"
	"os"
)

func main() {
	configLogging()

	port, present := os.LookupEnv("PORT")
	if !present {
		port = "5000"
	}

	var server = Server{
		Port: port,
	}

	err := server.Start()

	logger.Panic(err)
}

func configLogging() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logger.TextFormatter{})

	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logger.InfoLevel)
}
