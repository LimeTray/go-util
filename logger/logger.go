package logger

import (
	"fmt"
	"os"

	"github.com/bhoriuchi/go-bunyan/bunyan"
)

func RegisterLogger(name string) bunyan.Logger {
	var LoggerConfig bunyan.Config

	level := bunyan.LogLevelInfo
	if os.Getenv("SERVICE_NAME") != "" {
		name = os.Getenv("SERVICE_NAME")
	}
	if os.Getenv("LOG_LEVEL") != "" {
		level = os.Getenv("LOG_LEVEL")
	}
	if name == "" {
		name = "application-log"
	}
	LoggerConfig = bunyan.Config{
		Name:   name,
		Level:  level,
		Stream: os.Stdout,
	}
	log, err := bunyan.CreateLogger(LoggerConfig)
	if err != nil {
		fmt.Println("Error creating logger")
		panic(err)
	}
	log.Info("Logger registered")
	return log
}
