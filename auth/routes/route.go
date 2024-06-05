package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rolandnii/roland-auth/auth/handlers"
)


func Init(app *fiber.App) {
	app.Post("/register",handlers.RegisterUser)
}

