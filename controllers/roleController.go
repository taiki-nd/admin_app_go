package controllers

import (
	"admin_app_go/db"
	"admin_app_go/logic"
	"admin_app_go/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RolePermission struct {
	Name        string
	Permissions []int
}

func RoleIndex(c *fiber.Ctx) error {
	var roles []models.Role
	db.DB.Find(&roles)

	log.Println("show all roles")

	return c.JSON(roles)
}

func RoleCreate(c *fiber.Ctx) error {
	var rolePermission RolePermission

	err := c.BodyParser(&rolePermission)
	if err != nil {
		log.Printf("POST method error: %s", err)
		return err
	}

	permissions := make([]models.Permission, len(rolePermission.Permissions))

	for i, permissionId := range rolePermission.Permissions {
		permissions[i] = models.Permission{
			Id: uint(permissionId),
		}
	}

	role := models.Role{
		Name:        rolePermission.Name,
		Permissions: permissions,
	}

	db.DB.Create(&role)
	log.Printf("create new role: id = %v", role.Id)

	return c.JSON(&role)

	/*
		var role models.Role

		err := c.BodyParser(&role)
		if err != nil {
			log.Printf("POST method error: %s", err)
			return err
		}

		db.DB.Create(&role)
		log.Printf("create new role: id = %v", role.Id)

		return c.JSON(&role)
	*/
}

func RoleShow(c *fiber.Ctx) error {
	role := logic.GetRoleFromId(c)

	db.DB.Find(&role)
	log.Printf("show Role: id = %v", role.Id)

	return c.JSON(role)
}

func RoleUpdate(c *fiber.Ctx) error {
	role := logic.GetRoleFromId(c)

	err := c.BodyParser(&role)
	if err != nil {
		log.Printf("PUT method error: %s", err)
		return err
	}

	db.DB.Model(&role).Updates(role)
	log.Printf("update role: id = %s", err)

	return c.JSON(role)
}

func RoleDelete(c *fiber.Ctx) error {
	role := logic.GetRoleFromId(c)

	db.DB.Delete(role)
	log.Printf("delete role")

	return nil
}
