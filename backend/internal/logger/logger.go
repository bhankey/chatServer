package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(c *Config) (*Logger, error) {
	l := &Logger{
		logrus.New(),
	}
	level, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return nil, err
	}
	l.SetLevel(level)
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return l, nil
}
