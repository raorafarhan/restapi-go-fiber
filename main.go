package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"restapi-gofiber/config"
	"restapi-gofiber/factory"
	"restapi-gofiber/migration"
	"restapi-gofiber/utils/database/mysql"
	"strconv"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitMysqlDB(cfg)
	e := fiber.New()
	migration.InitMigrate(db)
	factory.InitFactory(e, db)

	fmt.Println("server running on port " + strconv.Itoa(cfg.SERVERPORT))
	log.Fatal(e.Listen("0.0.0.0.:" + strconv.Itoa(cfg.SERVERPORT)))

}
