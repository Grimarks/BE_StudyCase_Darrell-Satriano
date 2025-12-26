package controllers

import (
	"StudyCase3/config"
	"StudyCase3/models"
	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event
	c.BodyParser(&event)

	config.DB.Create(&event)
	return c.JSON(event)
}

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	config.DB.Find(&events)
	return c.JSON(events)
}
