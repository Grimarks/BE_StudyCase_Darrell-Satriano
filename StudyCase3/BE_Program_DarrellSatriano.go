package main

import (
	"StudyCase3/config"
	"StudyCase3/models"
	"StudyCase3/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Transaction{})

	routes.Setup(app)

	app.Listen(":8080")
}
