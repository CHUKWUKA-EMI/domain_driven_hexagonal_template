package logger

import "github.com/sirupsen/logrus"

// New creates a new logger
func New() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	return logger
}
