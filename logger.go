package logger

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger
var slogger *zap.SugaredLogger

var loggerMu sync.Mutex

func Logger() *zap.Logger {
	if logger == nil {
		loggerMu.Lock()
		defer loggerMu.Unlock()
		if IsDebug() {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction()
		}
	}

	return logger
}

func SLogger() *zap.SugaredLogger {
	if slogger == nil {
		slogger = Logger().Sugar()
	}

	return slogger
}

func Errorw(ctx context.Context, msg string, args ...interface{}) {
	SLogger().Errorw(msg, args...)
}

func Debugw(ctx context.Context, msg string, args ...interface{}) {
	SLogger().Debugw(msg, args...)
}

func IsDebug() bool {
	return os.Getenv("DEBUG_LOG") == "1" ||
		os.Getenv("DEBUG_LOG") == "true"
}
