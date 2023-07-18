package models

import (
	"time"
	"github.com/rs/xid"
)

type CartItem struct {
	IDHost  xid.ID `gorm:"type:string(20);primaryKey;not null"`
	IDGuest  xid.ID `gorm:"type:string(20);primaryKey;not null"`
	Fid  string `gorm:"type:serial;primaryKey;not null"`
	Number      int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
