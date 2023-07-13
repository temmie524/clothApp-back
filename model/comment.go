package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User    User   `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId  uint   `json:"user_id" gorm:"not null"`
	Item    Item   `json:"item_id" gorm:"foreignKey:ItemId; constraint:OnDelete:CASCADE"`
	ItemId  uint   `json:"-" gorm:"column:item_id"`
	Comment string `json:"comment"`
}

type CommentResponse struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
