package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name		string    `json:"name"`
	Price		uint      `json:"price"`
	ID User