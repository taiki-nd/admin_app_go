package controllers

import (
	"admin_app_go/db"
	"admin_app_go/logic"
	"admin_app_go/models"
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ProductIndex(c *fiber.Ctx) error {
	var products []models.Product
	limit := 5
	page, _ := strconv.Atoi(c.Query("page", "1"))
	offset := (page - 1) * limit
	db.DB.Offset(offset).Limit(limit).Find(&products)

	var total int64
	db.DB.Model(&models.Product{}).Count(&total)

	log.Println("show all products")

	lastPage := math.Ceil(float64(total) / float64(limit))

	return c.JSON(fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	})
}

func ProductCreate(c *fiber.Ctx) error {
	var product models.Product

	err := c.BodyParser(&product)
	if err != nil {
		log.Printf("POST method error: %s", err)
		return err
	}

	db.DB.Create(&product)
	log.Printf("create new product: id = %v", product.Id)

	return c.JSON(&product)
}

func ProductShow(c *fiber.Ctx) error {
	product := logic.GetProductFromId(c)

	db.DB.Find(&product)
	log.Printf("show product: id = %v", product.Id)

	return c.JSON(product)
}

func ProductUpdate(c *fiber.Ctx) error {
	product := logic.GetProductFromId(c)

	err := c.BodyParser(&product)
	if err != nil {
		log.Printf("PUT method error: %s", err)
		return err
	}

	db.DB.Model(&product).Updates(product)
	log.Printf("update product: id = %s", err)

	return c.JSON(product)
}

func ProductDelete(c *fiber.Ctx) error {
	product := logic.GetProductFromId(c)

	db.DB.Delete(product)
	log.Printf("delete product")

	return nil
}
