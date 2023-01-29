package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"restapi-gofiber/config"
	"restapi-gofiber/factory"
	"restapi-gofiber/migration"
	"restapi-gofiber/utils/database/mysql"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitMysqlDB(cfg)
	e := fiber.New()
	migration.InitMigrate(db)
	factory.InitFactory(e, db)

	log.Fatal(e.Listen(os.Getenv(":/" + "SERVERPORT")))
}
