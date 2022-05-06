package logic

import (
	"admin_app_go/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetRoleFromId(c *fiber.Ctx) models.Role {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	return role
}
