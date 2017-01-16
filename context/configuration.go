package context

import (
	"os"
	"bytes"
)

type Configuration struct {
	MongoConfiguration  *MongoConfiguration
	ServerConfiguration *ServerConfiguration
}

type MongoConfiguration struct {
	Host     string
	Database string
}

type ServerConfiguration struct {
	Port string
}

func Build() *Configuration {
	config := &Configuration{
		MongoConfiguration: buildMongoConfiguration(),
		ServerConfiguration: buildServerConfiguration(),
	}

	return config
}

func buildServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		Port: lookupEnvOrDefault("SERVER_PORT", "5000"),
	}
}

func buildMongoConfiguration() *MongoConfiguration {
	return &MongoConfiguration{
		Host: lookupMandatoryEnv("MONGO_HOST"),
		Database: lookupMandatoryEnv("MONGO_DATABASE"),
	}
}

func lookupEnvOrDefault(key string, other string) string {
	value, present := os.LookupEnv(key)
	if !present {
		value = other
	}
	return value
}

func lookupMandatoryEnv(key string) string {
	mongoDatabase, present := os.LookupEnv(key)

	if !present {
		var buffer bytes.Buffer
		buffer.WriteString("Enviroment variable \"")
		buffer.WriteString(key)
		buffer.WriteString("\" has not been found")

		panic(buffer.String())
	}
	return mongoDatabase
}