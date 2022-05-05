package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
