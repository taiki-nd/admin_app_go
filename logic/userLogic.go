package logic

import (
	"admin_app_go/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUserFromId(c *fiber.Ctx) models.User {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	return user
}
