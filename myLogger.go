package myLogger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logLevel = logrus.DebugLevel


func getLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.ErrorLevel:
		return "ERROR"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.PanicLevel:
		return "PANIC"
	case logrus.FatalLevel:
		return "FATAL"
	}
	return "UNKNOWN"
}

// GetLogger returs logger
func GetLogger() *logrus.Logger {
	logger := logrus.New()
	// logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logger.SetLevel(logLevel)
	logger.SetOutput(os.Stdout)

	return logger
}
