package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	SERVERPORT  int
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()

	}

	return appConfig
}

func initConfig() *AppConfig {
	var defConfig AppConfig
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	defConfig.DB_USERNAME = os.Getenv("MYSQLUSER")
	defConfig.DB_PASSWORD = os.Getenv("MYSQLPASSWORD")
	defConfig.DB_HOST = os.Getenv("MYSQLHOST")
	defConfig.DB_NAME = os.Getenv("MYSQLDATABASE")

	dbPortConv, errDBPort := strconv.Atoi(os.Getenv("MYSQLPORT"))

	if errDBPort != nil {
		log.Fatal(errDBPort)
		return nil
	}

	defConfig.DB_PORT = dbPortConv

	serverPortConv, errServerPort := strconv.Atoi(os.Getenv("SERVERPORT"))

	if errServerPort != nil {
		log.Fatal(errServerPort)
		return nil
	}

	defConfig.SERVERPORT = serverPortConv

	return &defConfig
}
