package db

import (
	"admin_app_go/config"
	"admin_app_go/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config.UserDevelop, config.Config.PasswordDevelop,
		config.Config.HostDevelop, config.Config.PortDevelop,
		config.Config.NameDevelop)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	log.Printf("success db connection: %v", db)

	DB = db

	db.AutoMigrate(&models.Role{}, &models.User{}, &models.Permission{})
	//db.Migrator().DropTable(&models.User{}, &models.Role{})
	//db.Migrator().CreateTable(&models.User{}, &models.Role{})
}
