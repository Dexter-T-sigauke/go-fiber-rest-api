package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Product struct {
	ID           uuid.UUID `json:"id" gorm:"type:text;primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}
