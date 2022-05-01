package routes

import (
	"admin_app_go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Get("/other", controllers.Bye)
}
