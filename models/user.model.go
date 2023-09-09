package models

import (
	"time"
)

type User struct {
	// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v3();primary_key"`
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
