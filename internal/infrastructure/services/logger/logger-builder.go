package logger

import (
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go-rest/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"path/filepath"
)

type loggerBuilder struct {
	configuration *config.Configuration
}

func NewLoggerBuilder(c *config.Configuration) *loggerBuilder {
	return &loggerBuilder{c}
}

func (l *loggerBuilder) BuildLoggerFromJSON() *zap.Logger {
	var zapConfig *zap.Config

	pwd, _ := os.Getwd()
	configPath := filepath.Join(pwd, fmt.Sprintf("log.%s.json", l.configuration.GoEnv))
	bytes, err := ioutil.ReadFile(configPath)
	err = json.Unmarshal(bytes, &zapConfig)

	logger, err := zapConfig.Build()

	if err != nil {
		panic("Error occurred during logger configuring")
	}

	return logger
}

func (l *loggerBuilder) BuildAdvancedLogger() *zap.Logger {
	// TODO: check writer configuration and multi environment
	// TODO: separate logs into level named log files
	encoder := l.getEncoder()
	writeSyncerConsole := l.getConsoleWriter()

	var core zapcore.Core
	consoleCore := zapcore.NewCore(encoder, writeSyncerConsole, zapcore.DebugLevel)

	if !l.configuration.IsDev {
		writeSyncerFile := l.getFileWriter("logs")
		fileCore := zapcore.NewCore(encoder, writeSyncerFile, zapcore.DebugLevel)
		core = zapcore.NewTee(consoleCore, fileCore)
	} else {
		core = consoleCore
	}

	return zap.New(core)
}

func (l *loggerBuilder) getEncoder() zapcore.Encoder {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	if !l.configuration.IsDev {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderCfg)
}

func (l *loggerBuilder) getConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func (l *loggerBuilder) getFileWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./logs/%s.log", fileName),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
