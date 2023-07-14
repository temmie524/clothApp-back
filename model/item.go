package model

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string   `json:"name" gorm:"not null"`
	Price      uint     `json:"price"`
	User       User     `json:"id" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint     `json:"user_id" gorm:"not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId"`
	CategoryId uint     `json:"category_id" gorm:"not null"`
	Brand      Brand    `json:"brand" gorm:"foreignKey:BrandId"`
	BrandId    uint     `json:"brand_id" gorm:"not null"`
	Season     Season   `json:"season" gorm:"foreignKey:SeasonId"`
	SeasonId   uint     `json:"season_id"`
	BuyDate    uint     `json:"buy_date"`
	Picture    []byte   `json:"picture" gorm:"not null"`
}
