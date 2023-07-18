package models

import (
	"time"
	"github.com/rs/xid"
)

type UserQA struct {
	ID  xid.ID `gorm:"type:string(20);primaryKey;not null"`
	Qid int `gorm:"type:int;primaryKey;not null"`
	Question  string `gorm:"type:string(256);not null"`
	Answer  string `gorm:"type:string(256);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
