package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// create a logrus instance
var logger *logrus.Logger

// init function to initialize
func init() {
	logger = logrus.New()
	configureLogger()
}

// function to configure the logger
func configureLogger() {

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info" // if the log level is not yet set -> use info as standard
	}

	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Fatal("Error by parsing the log level: ", err)
	}

	logger.SetLevel(parsedLogLevel)
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

// logs Trace level and higher
func Trace(message interface{}) {
	logger.Trace("Trace: ", message)
}

// logs Debug level and higher
func Debug(message interface{}) {
	logger.Debug("Debug: ", message)
}

// logs Info level and higher
func Info(message interface{}) {
	logger.Info("Info: ", message)
}

// logs Warn level and higher
func Warning(message interface{}) {
	logger.Warn("Warning: ", message)
}

// logs Error level and higher
func Error(message interface{}) {
	logger.Error("Error: ", message)
}

// logs Fatal level and higher
func Fatal(message interface{}) {
	logger.Fatal("Fatal: ", message)
}

// logs Panic level and higher
func Panic(message interface{}) {
	logger.Panic("Panic: ", message)
}
