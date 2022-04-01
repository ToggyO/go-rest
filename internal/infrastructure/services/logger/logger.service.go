package logger

import (
	"go-rest/internal/application/contracts"
	"go-rest/internal/config"
	"go.uber.org/zap"
)

type loggerService struct {
	logger *zap.Logger
}

// NewLoggerService - creates new logger instance
func NewLoggerService(configuration *config.Configuration) contracts.ILogger {
	loggerBuilder := NewLoggerBuilder(configuration)
	return &loggerService{
		loggerBuilder.BuildAdvancedLogger(),
	}
}

// TODO: implement second param of type zap.Field...
func (l *loggerService) Debug(message string) {
	l.logger.Debug(message)
}

// TODO: implement second param of type zap.Field...
func (l *loggerService) Info(message string) {
	l.logger.Info(message)
}

// TODO: implement second param of type zap.Field...
func (l *loggerService) Warn(message string) {
	l.logger.Warn(message)
}

// TODO: implement second param of type zap.Field...
func (l *loggerService) Error(message string) {
	l.logger.Error(message)
}

// TODO: implement second param of type zap.Field...
func (l *loggerService) Fatal(message string) {
	l.logger.Fatal(message)
}

func (l *loggerService) Flush() {
	l.logger.Sync()
}
