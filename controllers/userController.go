package controllers

import (
	"admin_app_go/db"
	"admin_app_go/logic"
	"admin_app_go/middlewares"
	"admin_app_go/models"
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserIndex(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		log.Println(err)
		return err
	}

	var users []models.User
	limit := 5
	page, _ := strconv.Atoi(c.Query("page", "1"))
	offset := (page - 1) * limit
	db.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)

	var total int64
	db.DB.Model(&models.User{}).Count(&total)

	log.Println("show all users")

	lastPage := math.Ceil(float64(total) / float64(limit))

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	})
}

func UserCreate(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		log.Println(err)
		return err
	}

	var user models.User

	err = c.BodyParser(&user)
	if err != nil {
		log.Printf("POST method error: %s", err)
		return err
	}

	user.SetPassword(string(user.Password))

	db.DB.Create(&user)
	log.Printf("create new user: id = %v", user.Id)

	return c.JSON(&user)
}

func UserShow(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		log.Println(err)
		return err
	}

	user := logic.GetUserFromId(c)

	db.DB.Preload("Role").Find(&user)
	log.Printf("show user: id = %v", user.Id)

	return c.JSON(user)
}

func UserUpdate(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		log.Println(err)
		return err
	}

	user := logic.GetUserFromId(c)

	err = c.BodyParser(&user)
	if err != nil {
		log.Printf("PUT method error: %s", err)
		return err
	}

	db.DB.Model(&user).Updates(user)
	log.Printf("update user: id = %v", user.Id)

	return c.JSON(user)
}

func UserDelete(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		log.Println(err)
		return err
	}

	user := logic.GetUserFromId(c)

	db.DB.Delete(user)
	log.Printf("delete user")

	return nil
}
