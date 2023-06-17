package models

import (
	"time"
)

type Rate struct {
	ID   uint      `gorm:"primaryKey" json:"id,omitempty"`
	Code string    `gorm:"varchar(3);not null" json:"code"`
	Rate float32   `gorm:"not null" json:"rate"`
	Date time.Time `gorm:"not null" json:"date"`
}