package models

import (
	"time"
	"github.com/gofrs/uuid"

)
type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:text;primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
