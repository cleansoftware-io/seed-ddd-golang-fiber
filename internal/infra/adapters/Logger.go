package adapters

import (
	"github.com/sirupsen/logrus"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/ports"
)

type LoggerImpl struct {
	logger *logrus.Logger
}

func NewLoggerImpl(logger *logrus.Logger) ports.Logger {
	return &LoggerImpl{
		logger: logger,
	}
}

func (l *LoggerImpl) Info(args ...interface{}) {
	l.logger.Info(&args)
}

func (l *LoggerImpl) Error(args ...interface{}) {
	l.logger.Error(&args)
}

func (l *LoggerImpl) Warning(args ...interface{}) {
	l.logger.Warning(&args)
}

func (l *LoggerImpl) Fatal(args ...interface{}) {
	l.logger.Fatal(&args)
}

func (l *LoggerImpl) Debug(args ...interface{}) {
	l.logger.Debug(&args)
}
