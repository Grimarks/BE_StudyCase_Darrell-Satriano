package controllers

import (
	"StudyCase3/config"
	"StudyCase3/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm/clause"
)

func CreateTransaction(c *fiber.Ctx) error {
	var trx models.Transaction

	if err := c.BodyParser(&trx); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// ambil data user dair jwt
	user := c.Locals("user").(jwt.MapClaims)
	userID := uint(user["id"].(float64))
	trx.UserID = userID

	tx := config.DB.Begin()

	var event models.Event
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&event, trx.EventID).Error; err != nil {
		tx.Rollback()
		return c.Status(404).JSON(fiber.Map{
			"error": "Event not found",
		})
	}

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
