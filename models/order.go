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
	Id           uint      `json:"id" gorm:"primarykey"`
	OrderId      uint      `json:"order_id"`
	ProductTitle string    `json:"product_title"`
	Price        float64   `json:"price"`
	Quantity     int       `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
