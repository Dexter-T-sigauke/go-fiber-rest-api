package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    uuid.UUID     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}
