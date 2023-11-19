package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"encoding/json"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

type AvailableTime struct {
	WeekDay     string `json:"day"`
	StartHour   int    `json:"start_hour"`
	StartMinute int    `json:"start_min"`
	EndHour     int    `json:"end_hour"`
	EndMinute   int    `json:"end_min"`
}

func (c *AvailableTime) Scan(value interface{}) error {
	// Convert the value to a byte slice
	// if value == nil {
	// 	c = nil
	// 	return nil
	// }
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan AvailableTime")
	}

	// Unmarshal the byte slice into the AvailableTime struct
	err := json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}

	return nil
}

// Value return json value, implement driver.Valuer interface
func (c AvailableTime) Value() (driver.Value, error) {
	// Marshal the AvailableTime struct into a byte slice
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

type User struct {
	ID            xid.ID `gorm:"type:string(20);primaryKey"`
	Account       string `gorm:"type:string(30);unique;not null"`
	Password      string `gorm:"type:string(64);not null"`
	Username      string `gorm:"type:string(30);not null"`
	AddressState  string `gorm:"type:string(64)"`
	AddressCity   string `gorm:"type:string(64)"`
	AddressStreet string `gorm:"type:string(128)"`
	AddressUnit   string `gorm:"type:string(32)"`
	Zipcode       string `gorm:"type:string(8)"`
	// In order to reduce the number of google map api calls
	// we store latitude and longitude of the user whenever
	// his address is updated.
	Lat                float64           // latitude of the user
	Lon                float64           // longitude of the user
	PhoneNum           string            `gorm:"type:string(10);not null"`
	Email              string            `gorm:"type:string(50);not null"`
	AvailableTimeSlots []AvailableTime   `gorm:"type:jsonb;"`
	PaymentMethod      map[string]string `gorm:"type:jsonb;"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = xid.New()
	return nil
}
