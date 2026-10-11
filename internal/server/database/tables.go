package database

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	ID        string
	Hostname  string
	LastSeen  time.Time
	CreatedAt time.Time
}
