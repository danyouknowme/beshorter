package main

import (
	"log"

	"github.com/danyouknowme/beshorter/config"
	"github.com/danyouknowme/beshorter/database"
	"github.com/danyouknowme/beshorter/handler"
	"github.com/danyouknowme/beshorter/httpserver"
	"github.com/danyouknowme/beshorter/logger"
	"github.com/danyouknowme/beshorter/repository"
	"github.com/danyouknowme/beshorter/service"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}

	logger.InitLogger(logger.LoggerConfig{
		Env: config.Env,
	})

	db, err := database.NewDatabaseConnection(database.DatabaseConfig{
		Driver:       config.DB_Driver,
		Hostname:     config.DB_Hostname,
		Port:         config.DB_Port,
		Username:     config.DB_Username,
		Password:     config.DB_Password,
		DatabaseName: config.DB_Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	shortUrlRepository := repository.NewShortUrlRepository(db)

	shortUrlService := service.NewShortUrlService(shortUrlRepository)

	shortUrlHandler := handler.NewShortUrlHandler(shortUrlService)

	server := httpserver.NewHTTPServer()
	server.GET("/healthcheck", handler.HealthCheckHandler)
	server.GET("/:url", shortUrlHandler.GetShortenerUrlHandler)
	server.POST("/url/shorter", shortUrlHandler.CreateShortenerUrlHandler)

	server.Run(config.ServerAddress)
}
