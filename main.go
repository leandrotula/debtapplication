package main

import (
	"database/sql"
	"debtsapp/cmd/handler"
	_ "debtsapp/docs"
	"debtsapp/internal/env"
	"debtsapp/internal/service"
	"debtsapp/internal/service/ping"
	"debtsapp/internal/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title          debs API
// @version        1.0
// @description    Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @schemes http
func main() {

	db, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	app := &handler.Application{
		Storage: storage.NewStorage(db),
		Configuration: handler.Configuration{
			Port: env.GetString("PORT", "8080"),
		},
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userService := service.NewUserService(app.Storage)

	router.POST("/v1/users", userService.Save)

	router.GET("/ping", ping.Ping())

	err = router.Run(app.Configuration.Port)
	if err != nil {
		log.Fatal("Could not start webserver: ", err)
	}

}
