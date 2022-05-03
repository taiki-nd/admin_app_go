package controllers

import (
	"admin_app_go/db"
	"admin_app_go/logic"
	"admin_app_go/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	log.Println("start register")

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("POST method error: %s", err)
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		log.Println("password & password_confirm dose not match.")
		return c.JSON(fiber.Map{
			"message": "password & password_confirm dose not match.",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	db.DB.Create(&user)
	log.Printf("finish register: %v", user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	log.Println("start login")
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("POST method error: %s", err)
		return err
	}

	var user models.User
	db.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		log.Printf("user not found: email = %s", data["email"])
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.Status(400)
		log.Printf("incorrect password: ID = %v, email = %s", user.Id, user.Email)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := logic.GenerateJwt(int(user.Id))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	log.Printf("login success: %s", data["email"])

	return c.JSON(token)
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	issuer, _ := logic.ParseJwt(cookie)

	var user models.User
	db.DB.Where("id =?", issuer).First(&user)

	return c.JSON(&user)
}

func Logout(c *fiber.Ctx) error {
	log.Printf("start logout")
	issuer, err := logic.ParseJwt(c.Cookies("jwt"))
	if err != nil {
		return err
	}
	user := models.User{}
	db.DB.Where("id =?", issuer).First(&user)
	log.Printf("start logout: ID = %v, Email = %s", user.Id, user.Email)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	log.Println("logout success")

	return c.JSON(fiber.Map{
		"message": "success logout",
	})
}
