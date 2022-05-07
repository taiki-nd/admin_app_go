package logic

import (
	"admin_app_go/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProductFromId(c *fiber.Ctx) models.Product {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	return product
}
