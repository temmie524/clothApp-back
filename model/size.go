package model

import "gorm.io/gorm"

type Size struct {
	gorm.Model
	Size_jp string `json:"size_jp" gorm:"not null"`
}
