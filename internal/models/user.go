package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/lib/pq"
	"github.com/rs/xid"
)

type User struct {
	ID        xid.ID `gorm:"type:string(20);primaryKey"`
	Username  string `gorm:"type:string(30);unique;not null"`
	Password  string `gorm:"type:string(64);not null"`
	Name      string `gorm:"type:string(30);not null"`
	AddressState  string `gorm:not null`
	AddressCity  string `gorm:not null`
	AddressStreet  string `gorm:not null`
	AddressUnit  string
	PhoneNum  string `gorm:"type:string(10);not null"`
	Email	  string `gorm:"type:string(50);not null"`
	Waiver	  bool
	PaymentMethod	pq.StringArray `gorm:"type:text[];not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = xid.New()
	return nil
}