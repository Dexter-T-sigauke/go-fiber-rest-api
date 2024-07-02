package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Order struct {
	ID           uuid.UUID `json:"id" gorm:"type:text;primaryKey"`
	CreatedAt    time.Time
	ProductRefer uuid.UUID `json:"product_id" gorm:"type:text"`
	Product      Product   `gorm:"foreignKey:ProductRefer"`
	UserRefer    uuid.UUID `json:"user_id" gorm:"type:text"`
	User         User      `gorm:"foreignKey:UserRefer"`
}