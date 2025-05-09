package configuration

import (
	"debtsapp/internal/configuration/token"
	"debtsapp/internal/storage"
	"go.uber.org/zap"
)

type Application struct {
	Storage            *storage.Storage
	Configuration      Configuration
	Logger             *zap.SugaredLogger
	ConfigurationToken *token.ConfigurationToken
}

type Configuration struct {
	Port string
}
