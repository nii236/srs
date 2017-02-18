package log

import "github.com/Sirupsen/logrus"

var logger *Logger

type Logger struct {
	*logrus.Logger
}

// New returns a new logger
func New() *Logger {
	if logger == nil {
		logger = &Logger{
			Logger: logrus.New(),
		}
		logger.Formatter = &logrus.TextFormatter{
			ForceColors: true,
		}

	}
	return logger
}
