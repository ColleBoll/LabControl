package database

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Hostname  string
	UpdatedAt time.Time
	CreatedAt time.Time
}
