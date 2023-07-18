package models

import (
	"time"
	"github.com/lib/pq"
	"github.com/rs/xid"
)

type Food struct {
	ID        xid.ID `gorm:"type:string(20);primaryKey;not null"`
	Fid  	  string `gorm:"type:serial;primaryKey;not null"`
	Name      string `gorm:"type:string(30);not null"`
	Ingradients  pq.StringArray `gorm:"type:text[]"`
	Notes  string `gorm:"type:string(256)"`
	OrderCutTime  time.Time `gorm:"not null"`
	Number  int `gorm:"not null"`
	Pics  pq.StringArray `gorm:"type:text[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
