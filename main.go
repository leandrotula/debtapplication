package main

import (
	"database/sql"
	"debtsapp/cmd/handler"
	_ "debtsapp/docs"
	"debtsapp/internal/configuration"
	"debtsapp/internal/env"
	"debtsapp/internal/service"
	"debtsapp/internal/storage"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
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
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func(logger *zap.SugaredLogger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)

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
		Storage: storage.NewStorage(db),
		Configuration: configuration.Configuration{
			Port: env.GetString("PORT", "8080"),
		},
		Logger: logger,
	}

	router := handler.CreateRouterApp()

	userService := service.NewUserService(app)

	router.POST("/v1/users", userService.CreateAndInvite)

	err = router.Run(app.Configuration.Port)
	logger.Infow("Webserver started using port: ", app.Configuration.Port, "successfully")

	if err != nil {
		logger.Fatalw("Could not start webserver: ", err)
	}

}
