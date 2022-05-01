package logic

import (
	"admin_app_go/db"
	"admin_app_go/models"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	jwt.StandardClaims
}

func GetUserFromCookie(cookie string, c *fiber.Ctx) (user models.User) {
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*Claims)

	db.DB.Where("id =?", claims.Issuer).First(&user)
	return user
}
