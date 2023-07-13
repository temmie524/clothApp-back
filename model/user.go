package model

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	Sex       int       `json:"sex"`
	Length    uint      `json:"length"`
	Weight    uint      `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
