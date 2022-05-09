package controllers

import (
	"admin_app_go/db"
	"admin_app_go/models"
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func OrderIndex(c *fiber.Ctx) error {
	var orders []models.Order
	limit := 5
	page, _ := strconv.Atoi(c.Query("page", "1"))
	offset := (page - 1) * limit
	db.DB.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	var total int64
	db.DB.Model(&models.Order{}).Count(&total)

	log.Println("show all orders")

	lastPage := math.Ceil(float64(total) / float64(limit))

	for i, _ := range orders {
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
	}

	return c.JSON(fiber.Map{
		"data": orders,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	})
}
