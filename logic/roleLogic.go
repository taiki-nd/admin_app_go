package logic

import (
	"admin_app_go/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RolePermission struct {
	Name        string
	Permissions []int
}

func GetRoleFromId(c *fiber.Ctx) models.Role {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	return role
}

func GetPermissions(rolePermission RolePermission) []models.Permission {
	permissions := make([]models.Permission, len(rolePermission.Permissions))

	for i, permissionId := range rolePermission.Permissions {
		permissions[i] = models.Permission{
			Id: uint(permissionId),
		}
	}

	return permissions
}
