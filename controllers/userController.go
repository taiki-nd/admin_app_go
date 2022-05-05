package controllers

import (
	"admin_app_go/db"
	"admin_app_go/logic"
	"admin_app_go/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserIndex(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)

	return c.JSON(users)
}

func UserCreate(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		log.Printf("POST method error: %s", err)
		return err
	}

	user.SetPassword(string(user.Password))

	db.DB.Create(&user)

	return c.JSON(&user)
}

func UserShow(c *fiber.Ctx) error {
	user := logic.GetUserFromId(c)

	db.DB.Find(&user)

	return c.JSON(user)
}

func UserUpdate(c *fiber.Ctx) error {
	user := logic.GetUserFromId(c)

	err := c.BodyParser(&user)
	if err != nil {
		log.Printf("PUT method error: %s", err)
		return err
	}

	db.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func UserDelete(c *fiber.Ctx) error {
	user := logic.GetUserFromId(c)

	db.DB.Delete(user)

	return nil
}
