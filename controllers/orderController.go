package controllers

import (
	"admin_app_go/db"
	"admin_app_go/models"
	"encoding/csv"
	"log"
	"math"
	"os"
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

func ExportCsv(c *fiber.Ctx) error {
	filepath := "./csv/orders.csv"

	err := CreateCsv(filepath)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.Download(filepath)
}

func CreateCsv(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var orders []models.Order
	db.DB.Preload("OrderItems").Find(&orders)
	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	for _, order := range orders {
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"",
			"",
			"",
		}
		err := writer.Write(data)
		if err != nil {
			log.Println(err)
			return err
		}

		for _, orderItem := range order.OrderItems {
			data := []string{
				"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(orderItem.Quantity),
			}
			err := writer.Write(data)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	return nil
}
