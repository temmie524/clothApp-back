package model

import "gorm.io/gorm"

type Season struct {
	gorm.Model
	Season string `json:"season" gorm:"not null"`
}
