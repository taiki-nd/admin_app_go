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
	app.Put("/api/user/info", controllers.UpdatesInfo)
	app.Put("/api/user/password", controllers.UpdatesPassword)
	app.Post("api/logout", controllers.Logout)

	app.Get("/api/users", controllers.UserIndex)
	app.Post("/api/users", controllers.UserCreate)
	app.Get("/api/users/:id", controllers.UserShow)
	app.Put("/api/users/:id", controllers.UserUpdate)
	app.Delete("/api/users/:id", controllers.UserDelete)

	app.Get("/api/roles", controllers.RoleIndex)
	app.Post("/api/roles", controllers.RoleCreate)
	app.Get("/api/roles/:id", controllers.RoleShow)
	app.Put("/api/roles/:id", controllers.RoleUpdate)
	app.Delete("/api/roles/:id", controllers.RoleDelete)

	app.Get("/api/products", controllers.ProductIndex)
	app.Post("/api/products", controllers.ProductCreate)
	app.Get("/api/products/:id", controllers.ProductShow)
	app.Put("/api/products/:id", controllers.ProductUpdate)
	app.Delete("/api/products/:id", controllers.ProductDelete)

	app.Get("/api/permissions", controllers.PermissionIndex)
}
