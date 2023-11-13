package models

import (
	"time"

	"github.com/lib/pq"
	"github.com/rs/xid"
)

// TODO: food should be added with an available time. This is for
// one time meal providers.
type Food struct {
	ID           xid.ID         `gorm:"type:string(20);primaryKey;not null"`
	Fid          string         `gorm:"type:serial;primaryKey;not null"`
	Name         string         `gorm:"type:string(30);not null"`
	Ingradients  pq.StringArray `gorm:"type:text[]"`
	Notes        string         `gorm:"type:string(256)"`
	OrderCutTime time.Time      `gorm:"not null"`
	Number       int            `gorm:"not null"`
	Pics         pq.StringArray `gorm:"type:text[]"`
	StartTime    time.Time      // The start time for the food to be available. Served for one time meal. Could be null
	EndTime      time.Time      // The end time for the food to be available
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
