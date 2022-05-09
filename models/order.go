package models

import (
	"time"
)

type Order struct {
	Id        uint        `json:"id" gorm:"primarykey"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email" gorm:"unique"`
	OrderItem []OrderItem `json:""order_item gorm:"foreignKey:OrderId"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type OrderItem struct {
	Id           uint `json:"id" gorm:"primarykey"`
	OrderId      uint
	ProductTitle string
	Price        float64
	Quantity     int
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
