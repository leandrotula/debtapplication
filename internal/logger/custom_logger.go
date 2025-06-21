package logger

import (
	"go.uber.org/zap"
	"sync"
)

var (
	once     sync.Once
	instance *zap.SugaredLogger
)

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		instance = logger.Sugar()
	})
	return instance
}

func Cleanup() error {
	if instance != nil {
		return instance.Sync()
	}
	return nil
}
