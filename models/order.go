package models

import (
	"time"
)

type Order struct {
	Id         uint        `json:"id" gorm:"primarykey"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Email      string      `json:"email" gorm:"unique"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
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
