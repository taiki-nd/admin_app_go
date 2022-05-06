package models

import (
	"time"
)

type Role struct {
	Id          uint         `json:"id" gorm:"primarykey"`
	Name        string       `json:"name"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}
