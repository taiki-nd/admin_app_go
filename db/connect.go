package db

import (
	"admin_app_go/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Config.UserDevelop, config.Config.PasswordDevelop,
		config.Config.HostDevelop, config.Config.PortDevelop,
		config.Config.NameDevelop)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	log.Printf("success db connection: %v", db)
}