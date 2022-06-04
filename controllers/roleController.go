package controllers

import (
	"admin_app_go/db"
	"admin_app_go/logic"
	"admin_app_go/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RolePermission struct {
	Name        string
	Permissions []int
}

func RoleIndex(c *fiber.Ctx) error {
	var roles []models.Role
	db.DB.Preload("Permissions").Find(&roles)

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

	permissions := logic.GetPermissions(logic.RolePermission(rolePermission))

	role := models.Role{
		Name:        rolePermission.Name,
		Permissions: permissions,
	}

	db.DB.Create(&role)
	log.Printf("create new role: id = %v", role.Id)

	return c.JSON(&role)
}

func RoleShow(c *fiber.Ctx) error {
	role := logic.GetRoleFromId(c)

	db.DB.Preload("Permissions").Find(&role)
	log.Printf("show Role: id = %v", role.Id)

	return c.JSON(role)
}

func RoleUpdate(c *fiber.Ctx) error {
	var rolePermission RolePermission

	err := c.BodyParser(&rolePermission)
	if err != nil {
		log.Printf("PUT method error: %s", err)
		return err
	}

	permissions := logic.GetPermissions(logic.RolePermission(rolePermission))

	id, _ := strconv.Atoi(c.Params("id"))

	db.DB.Table("role_permissions").Where("role_id", id).Delete("")

	role := models.Role{
		Id:          uint(id),
		Name:        rolePermission.Name,
		Permissions: permissions,
	}

	db.DB.Model(&role).Updates(role)
	log.Printf("update role: id = %v", id)

	return c.JSON(role)
}

func RoleDelete(c *fiber.Ctx) error {
	//アソシエーションのくんであるものをあらかじめ削除する必要がある
	id, _ := strconv.Atoi(c.Params("id"))
	db.DB.Table("role_permissions").Where("role_id", id).Delete("")

	role := logic.GetRoleFromId(c)
	db.DB.Delete(role)
	log.Printf("delete role")

	return nil
}
