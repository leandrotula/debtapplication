package main

import (
	"debtsapp/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func StartWebserver(app *Application) {

	router := gin.Default()

	userService := service.NewUserService(app.Storage)

	router.POST("/v1/users", userService.Save)

	err := router.Run(app.Configuration.Port)
	if err != nil {
		log.Fatal("Could not start webserver: ", err)
	}

}
