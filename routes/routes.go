package routes

import (
	"admin_app_go/controllers"
	"admin_app_go/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)
	app.Get("/api/users", controllers.UserIndex)
	app.Post("/api/users", controllers.UserCreate)
	app.Get("/api/users/:id", controllers.UserShow)
}
