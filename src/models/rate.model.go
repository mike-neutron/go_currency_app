package models

import (
	"gorm.io/datatypes"
)

type Rate struct {
	ID   uint      `gorm:"primaryKey" json:"-"`
	Code string    `gorm:"varchar(3);not null;uniqueIndex:idx_unique" json:"code"`
	Rate float32   `gorm:"not null" json:"rate"`
	Date datatypes.Date `gorm:"not null;uniqueIndex:idx_unique" json:"date"`
}