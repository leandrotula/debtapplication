package handler

import (
	"debtsapp/internal/storage"
)

type Application struct {
	Storage       *storage.Storage
	Configuration Configuration
}

type Configuration struct {
	Port string
}
