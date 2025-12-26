package routes

import (
	"StudyCase3/controllers"
	"StudyCase3/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	app.Get("/events", controllers.GetEvents)

	app.Post("/events",
		middleware.AuthMiddleware,
		middleware.AdminOnly,
		controllers.CreateEvent,
	)

	app.Post("/transactions",
		middleware.AuthMiddleware,
		controllers.CreateTransaction,
	)
}
