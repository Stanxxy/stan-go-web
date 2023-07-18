package controller

import (
	"net/http"
	"github.com/rs/xid"
	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/Stanxxy/stan-go-web/internal/core/errors"
	"github.com/Stanxxy/stan-go-web/internal/models"
	"github.com/labstack/echo/v4"
)

// TODO:
// consider to allow search by address string.
// search and return the potential places.
// get the nearst address for Lat and Lon.
// seek to solve it with front end and only pass lat and lon in backend.

func (ctrl User) getBusiness(c echo.Context) {
	cc := c.(*context.AppContext)

	var getBusinessRequest GetBusinessRequest

	err := cc.Bind(&getBusinessRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	businesses := []models.User{}

	rowsAffected, err := cc.UserStore.RetrieveMany(&businesses)
	// use SortLocationsBasedOnCloseness to Sort the distance of each business to current location
	currentLocation := Location{
		Lat: getBusinessRequest.LocationLat, 
		Lon: getBusinessRequest.LocationLon
	}
	userSortedByDistance := SortLocationsBasedOnCloseness(currentLocation, businesses)

	// convert to business slice
	businessList := make([]Business, getBusinessRequest.Quantity) // we only need to get certain number of business
	// filter with open flag, startNumber and quantity

	for i, user := range userSortedByDistance {
		// check open
		
		isOpen := CheckBusinessOpenRightNow(user)
		if !isOpen && getBusinessRequest.Open {
			continue
		}
		
		// return a list of busispace to store ness
		// We could leave an extra table to store reason and available time
		reason := "No Reason"
		if !isOpen {
			reason = "Not in operation time."
		}

		if i >= getBusinessRequest.StartNum {
			businessList[i - StartNum] = Business {
				Uid: user.ID,
				BusinessName: user.Username,
				BusinessPhoneNum: user.PhoneNum,
				BusinessLocation: GeologicalLocation{
					Lat: User.Lat,
					Lon: User.Lon
				},
				BusinessAvailableTime: user.AvailableTimeSlots
				IsAvailable: isOpen,
				Reason: reason 
			}
		}
	}

	// Do something with the user object
	resp := GetBusinessResponse{
		Code: 0, 
		Message: "Address info updated",
		Data: businessList}
	return c.JSON(http.StatusOK, resp)
}

func (ctrl User) getBusinessByName(c echo.Context) {	
	cc := c.(*context.AppContext)

	var getBusinessByNameRequest GetBusinessByNameRequest

	err := cc.Bind(&getBusinessByNameRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Based on what condition is given, we go for different query logic
	queryCondition := make(map[string]string)
	// For cases when BusinessName is given, we use UserName to match BusinessName
	if len(getBusinessByNameRequest.BusinessName) != 0 {
		queryCondition["Username"] =  "%" + getBusinessByNameRequest.BusinessName + "%"
	}
	if len(getBusinessByNameRequest.Zipcode) != 0 {
		queryCondition["Zipcode"] =  getBusinessByNameRequest.Zipcode
	}
	if len(getBusinessByNameRequest.State) != 0 {
		queryCondition["AddressState"] =  getBusinessByNameRequest.State
	}
	if len(getBusinessByNameRequest.City) != 0 {
		queryCondition["AddressCity"] =  getBusinessByNameRequest.City
	}
	businesses := []models.User{}
	// For cass when zipcode is enabled, we search based on zipcode,
	// the same logic applies to city and state

	rowsAffected, err := cc.UserStore.RetrieveMany(queryCondition, &businesses)

	if err != nil {
		b := errors.NewBoom(errors.EntityUpdateError, errors.ErrorText(errors.EntityUpdateError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Do something with the user object
	resp := UpdatePaymentInfoResponse{user.ID, "payment info updated"}
	return c.JSON(http.StatusOK, resp)
}