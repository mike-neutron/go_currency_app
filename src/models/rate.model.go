package models

import (
	"gorm.io/datatypes"
	"time"
)

type Rate struct {
	ID   uint      `gorm:"primaryKey" json:"-"`
	Code string    `gorm:"varchar(3);not null;uniqueIndex:idx_unique" json:"code"`
	Rate float32   `gorm:"not null" json:"rate"`
	Date datatypes.Date `gorm:"not null;uniqueIndex:idx_unique" json:"date"`
}

type RateSwagger struct {
	ID   uint      `json:"-"`
	Code string    `json:"code"`
	Rate float32   `json:"rate"`
	Date time.Time `json:"date"`
}
