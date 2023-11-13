package controller

import (
	"github.com/Stanxxy/stan-go-web/internal/models"

	"github.com/rs/xid"
)

// Note1: For available time, we tend to put the demand of long term
// business as our priority. Thus, it's reasonable to set up time
// slots for the day in a week and select a list of days in a single
// week.

// Note1(cont'd): For the transient business or social meal providers,
// we will not provide available time but focus on the meal itself.
// It's resonable that these meals are not regularly and can only be
// available on certain time period. Thus, we condier to leave the available
// time field empty and set up the is_available flag to be 1.
// The buisiness should be marked by reason field to indicate
// the existance in discover page is because of one time activity.

type GeologicalLocation struct {
	Lat float64 `json:"location_lat"`
	Lon float64 `json:"location_lng"`
}

type Business struct {
	Uid                   xid.ID                 `json:"uid"`
	BusinessName          string                 `json:"business_name"`
	BusinessAddress       string                 `json:"business_address"`
	BusinessPhoneNum      string                 `json:"business_phone_num"`
	BusinessLocation      GeologicalLocation     `json:"business_location"`
	BusinessAvailableTime []models.AvailableTime `json:"business_available_time"`
	IsAvailable           int                    `json:"business_is_available"`
	Reason                string                 `json:"reason"`
}

// Request Types
type (
	GetBusinessRequest struct {
		LocationLat float64 `json:"location_lat" validate:"required"`
		LocationLon float64 `json:"location_lng" validate:"required"`
		Open        int     `json:"open" validate:"required"`
		StartNum    int     `json:"start_num" validate:"required"`
		Quantity    int     `json:"quantity" validate:"required"`
	}
	GetBusinessByNameRequest struct {
		BusinessName string `json:"business_name" validate:"required"`
		Zipcode      string `json:"zipcode" validate:"required"`
		City         string `json:"city" validate:"required"`
		State        string `json:"street" validate:"required"`
		Open         int    `json:"open" validate:"required"`
		StartNum     int    `json:"start_num" validate:"required"`
		Quantity     int    `json:"quantity" validate:"required"`
	}
)

// Response Types
type (
	GetBusinessResponse struct {
		Code    int        `json:"code"`
		Message string     `json:"message"`
		Data    []Business `json:"data"`
	}
	GetBusinessByNameResponse struct {
		Code    int        `json:"code"`
		Message string     `json:"message"`
		Data    []Business `json:"data"`
	}
)
