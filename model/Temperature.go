package models

import "time"

type Temperature struct {
	ID        uint    `gorm:"primarykey"`
	Name      string  `gorm:"default:null"`
	Temp      float32 `gorm:"default:0"`
	CreatedAt time.Time
}
