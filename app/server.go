package app

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/learninNdi/go-toko-aki/app/controllers"
)

var server controllers.Server
var appConfig controllers.AppConfig
var dbConfig controllers.DBConfig

func setGeneralConfig() {

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppURL = os.Getenv("APP_URL")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	dbConfig.DBDriver = os.Getenv("DB_DRIVER")
	dbConfig.DBHost = os.Getenv("DB_HOST")
	dbConfig.DBPort = os.Getenv("DB_PORT")
	dbConfig.DBUser = os.Getenv("DB_USER")
	dbConfig.DBPass = os.Getenv("DB_PASS")
	dbConfig.DBName = os.Getenv("DB_NAME")

}

func Run() {

	server = controllers.Server{}
	appConfig = controllers.AppConfig{}
	dbConfig = controllers.DBConfig{}

	if err := godotenv.Load(); err != nil {
		panic("Error loading env file")
	}

	setGeneralConfig()

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)

}
