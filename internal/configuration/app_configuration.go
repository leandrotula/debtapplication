package configuration

import (
	"debtsapp/internal/configuration/token"
	"debtsapp/internal/storage"
)

type Application struct {
	Storage            *storage.Storage
	Configuration      Configuration
	ConfigurationToken *token.ConfigurationToken
}

type Configuration struct {
	Port string
}

func NewApplication(
	storage *storage.Storage,
	configurationToken *token.ConfigurationToken) *Application {
	return &Application{
		Storage:            storage,
		Configuration:      Configuration{},
		ConfigurationToken: configurationToken,
	}
}
