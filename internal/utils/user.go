package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/Stanxxy/stan-go-web/internal/models"
)

type PasswordChangeTriplet struct {
	ReqID   string
	ReqTime time.Time
	UserId  string
}

func (c *PasswordChangeTriplet) Scan(value interface{}) error {
	// Convert the value to a byte slice
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to scan PasswordChangeTriplet")
	}

	// Unmarshal the byte slice into the AvailableTime struct
	err := json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}

	return nil
}

func (c PasswordChangeTriplet) Value() (driver.Value, error) {
	// Marshal the AvailableTime struct into a byte slice
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

func FindUserTimeZone(user *models.User) *time.Location {
	// TODO: find out user timezone
	return time.FixedZone("UTC-4", -4*60*60)
}
