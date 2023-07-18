package models

import (
	"time"
    "errors"
	"github.com/shopspring/decimal"
    "database/sql/driver"
    "encoding/json"
    "github.com/google/uuid"
)

type DetailedOrderItem struct {
    Fid string
    Number int
	Price decimal.Decimal
}

func (c *DetailedOrderItem) Scan(value interface{}) error {
    // Convert the value to a byte slice
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("Failed to scan DetailedOrderItem")
    }

    // Unmarshal the byte slice into the DetailedOrderItem struct
    err := json.Unmarshal(bytes, c)
    if err != nil {
        return err
    }

    return nil
}

func (c DetailedOrderItem) Value() (driver.Value, error) {
    // Marshal the DetailedOrderItem struct into a byte slice
    bytes, err := json.Marshal(c)
    if err != nil {
        return nil, err
    }

    return string(bytes), nil
}

type StateHistoryItem struct {
    State OrderState
    Time time.Time
}

func (c *StateHistoryItem) Scan(value interface{}) error {
    // Convert the value to a byte slice
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("Failed to scan StateHistoryItem")
    }

    // Unmarshal the byte slice into the CustomType struct
    err := json.Unmarshal(bytes, c)
    if err != nil {
        return err
    } 	

    return nil
}

func (c StateHistoryItem) Value() (driver.Value, error) {
    // Marshal the CustomType struct into a byte slice
    bytes, err := json.Marshal(c)
    if err != nil {
        return nil, err
    }

    return string(bytes), nil
}


type OrderDetails struct{
	Oid        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderItems   []DetailedOrderItem  `gorm:"type:jsonb;default:'[]'"`
	StateHistory   []StateHistoryItem  `gorm:"type:jsonb;default:'[]'"`
	Note	  string `gorm:"type:string(256)"`
    CreatedAt time.Time
	UpdatedAt time.Time
}