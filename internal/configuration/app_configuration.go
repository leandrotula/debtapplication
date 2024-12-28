package configuration

import (
	"debtsapp/internal/storage"
	"go.uber.org/zap"
)

type Application struct {
	Storage       *storage.Storage
	Configuration Configuration
	Logger        *zap.SugaredLogger
}

type Configuration struct {
	Port string
}
