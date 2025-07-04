package main

import (
	"database/sql"
	"debtsapp/cmd/handler"
	_ "debtsapp/docs"
	"debtsapp/internal/configuration"
	token2 "debtsapp/internal/configuration/token"
	"debtsapp/internal/env"
	clogger "debtsapp/internal/logger"
	"debtsapp/internal/service"
	"debtsapp/internal/service/encription"
	"debtsapp/internal/storage"
	"debtsapp/internal/token"
	_ "github.com/lib/pq"
)

// @title          debs API
// @version        1.0
// @description    debts api.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email ingleantula@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @schemes http
func main() {

	logger := clogger.GetLogger()

	logger.Infow("initializing debtsapp microservice")

	db, err := storage.New()
	if err != nil {
		logger.Fatalw("There was an error trying to configure db", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatalw("Error Trying to close db", err)
		}
	}(db)

	app := &configuration.Application{
		Storage: storage.InitDB(db),
		Configuration: configuration.Configuration{
			Port: env.GetString("PORT", "8080"),
		},
		ConfigurationToken: token2.NewConfigurationToken(
			env.GetExpirationDuration(),
			env.GetString("SECRET_TOKEN", "test"),
			env.GetString("AUDIENCE_TOKEN", "testaudience"),
			env.GetString("ISSUER_TOKEN", "testissuer")),
	}

	router := handler.CreateRouterApp()
	encryptionService := encription.NewDebtEncryption()

	userService := service.NewUserService(app, encryptionService)
	tokenService := token.NewTokenService(app)

	router.POST("/v1/users", userService.CreateAndInvite)
	router.PATCH("/v1/users", userService.ActivateUser)
	router.POST("/v1/token", tokenService.GenerateJwtToken)

	err = router.Run(app.Configuration.Port)
	logger.Infow("Webserver started using port: ", app.Configuration.Port, "successfully")

	if err != nil {
		logger.Fatalw("Could not start webserver: ", err)
	}

}
