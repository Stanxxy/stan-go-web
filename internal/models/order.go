package models

import (
	"fmt"
	"time"
	"errors"
	"github.com/shopspring/decimal"
	"database/sql/driver"
	"github.com/google/uuid"
	"github.com/rs/xid"
)

type OrderState string

const (
    ORDER_CREATED  OrderState = "order_created"
    WAITING_FOR_GROUPING OrderState = "waiting_for_grouping"
    WAITING_FOR_PAYMENT OrderState = "waiting_for_payment"
	WAITING_FOR_DELIVERY OrderState = "waiting_for_delivery"
	DELIVERING OrderState = "delivering"
	DELIVERED OrderState = "delivered"
	CANCELED OrderState = "canceled"
)

func (ost *OrderState) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan OrderState value:", value))
	}
    
	*ost = OrderState(bytes)
    return nil
}

func (ost OrderState) Value() (driver.Value, error) {
    return string(ost), nil
}

type DeliverMethod string

const (
    HOST_DELIVER  DeliverMethod = "host_deliver"
    PICK_UP DeliverMethod = "pick_up"
    BY_PLATFORM DeliverMethod = "by_platform"
	BY_THIRD_PARTY DeliverMethod = "by_third_party"
)

func (dm *DeliverMethod) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan DeliverMethod value:", value))
	}

    *dm = DeliverMethod(bytes)
    return nil
}

func (dm DeliverMethod) Value() (driver.Value, error) {
    return string(dm), nil
}

type Order struct {
	Oid        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	IDHost  xid.ID `gorm:"type:string(20)"`
	IDGuest  xid.ID `gorm:"type:string(20)"`
	State      OrderState `gorm:"type:string;not null"`
	TotalPrice   decimal.Decimal `gorm:"type:decimal(7,2);not null"`
	DeliverBy  DeliverMethod `gorm:"type:string"`
	IsGroup  bool `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
