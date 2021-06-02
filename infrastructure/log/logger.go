package log

import (
	"log"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	var logger = logrus.New()
	// set standard logger to logrus instance
	log.SetOutput(logger.Writer())
	logrus.SetOutput(logger.Writer())
	return logger
}
