package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserId    uint    `json:"user_id" gorm:"not null"`
	Comment   Comment `json:"comment" gorm:"foreignKey:CommentId; constraint:OnDelete:CASCADE"`
	CommentId uint    `json:"comment_id" gorm:"not null"`
}

type LikeResponse struct {
	Id uint `json:"id" gorm:"primary_key"`
}
