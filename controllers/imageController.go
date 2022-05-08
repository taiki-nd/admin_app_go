package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("uploads images error: %s", err)
		return err
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = file.Filename
		err := c.SaveFile(file, "./uploads/"+filename)
		if err != nil {
			log.Printf("failed to save uploads images: %s", err)
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:3000/api/uploads/" + filename,
	})
}
