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

	log.Println("show all users")

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
	log.Printf("create new user: id = %v", user.Id)

	return c.JSON(&user)
}

func UserShow(c *fiber.Ctx) error {
	user := logic.GetUserFromId(c)

	db.DB.Find(&user)
	log.Printf("show user: id = %v", user.Id)

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
	log.Printf("update user: id = %s", err)

	return c.JSON(user)
}

func UserDelete(c *fiber.Ctx) error {
	user := logic.GetUserFromId(c)

	db.DB.Delete(user)
	log.Printf("delete user")

	return nil
}
