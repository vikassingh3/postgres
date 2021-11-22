package main

import (
	"log"

	"github.com/vikas/config"
	"github.com/vikas/router"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())

	godotenv.Load()

	config.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))

	defer config.DB.Close()
}
