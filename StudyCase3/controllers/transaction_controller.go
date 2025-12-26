package controllers

import (
	"StudyCase3/config"
	"StudyCase3/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func CreateTransaction(c *fiber.Ctx) error {
	var trx models.Transaction
	c.BodyParser(&trx)

	tx := config.DB.Begin()

	var event models.Event
	tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&event, trx.EventID)

	if event.TicketsSold+trx.Quantity > event.Capacity {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{
			"error": "Ticket sold out",
		})
	}

	event.TicketsSold += trx.Quantity
	tx.Save(&event)
	tx.Create(&trx)
	tx.Commit()

	return c.JSON(trx)
}
