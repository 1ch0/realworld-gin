package log

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	logger := logrus.New()
	// Set log level to debug
	logger.SetLevel(logrus.DebugLevel)
	Logger = logger
}
