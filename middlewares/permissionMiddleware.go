package middlewares

import (
	"admin_app_go/db"
	"admin_app_go/logic"
	"admin_app_go/models"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	Id, err := logic.ParseJwt(cookie)
	if err != nil {
		return err
	}

	userID, _ := strconv.Atoi(Id)
	user := models.User{
		Id: uint(userID),
	}
	db.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}
	db.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("Unauthorized")
}
