package controller

import (
	"net/http"

	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/Stanxxy/stan-go-web/internal/core"
	"github.com/Stanxxy/stan-go-web/internal/core/errors"
	"github.com/Stanxxy/stan-go-web/internal/models"
	"github.com/Stanxxy/stan-go-web/internal/utils"
	echo "github.com/labstack/echo/v4"
)

// TODO:
// consider to allow search by address string.
// search and return the potential places.
// get the nearst address for Lat and Lon.
// seek to solve it with front end and only pass lat and lon in backend.

func GetBusiness(c echo.Context) error {
	cc := c.(*context.AppContext)

	var getBusinessRequest GetBusinessRequest

	err := cc.Bind(&getBusinessRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	businesses := []models.User{}

	_, err = cc.UserStore.RetrieveManyNoCondition(&businesses)

	if err != nil {
		b := errors.NewBoom(errors.EntityQueryError, errors.ErrorText(errors.EntityQueryError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// use SortLocationsBasedOnCloseness to Sort the distance of each business to current location
	currentLocation := utils.Location{
		Lat: getBusinessRequest.LocationLat,
		Lon: getBusinessRequest.LocationLon,
	}
	userSortedByDistance := utils.SortLocationsBasedOnCloseness(currentLocation, businesses)

	// convert to business slice
	businessList := make([]Business, getBusinessRequest.Quantity) // we only need to get certain number of business
	// filter with open flag, startNumber and quantity

	for i, user := range userSortedByDistance {
		// check open

		isOpen := utils.CheckBusinessOpenRightNow(user)
		if !isOpen && getBusinessRequest.Open != 0 {
			continue
		}

		// return a list of busispace to store ness
		// We could leave an extra table to store reason and available time
		reason := "No Reason"
		if !isOpen {
			reason = "Not in operation time."
		}

		if i >= getBusinessRequest.StartNum {
			businessList[i-getBusinessRequest.StartNum] = Business{
				Uid:              user.ID,
				BusinessName:     user.Username,
				BusinessPhoneNum: user.PhoneNum,
				BusinessLocation: GeologicalLocation{
					Lat: user.Lat,
					Lon: user.Lon,
				},
				BusinessAvailableTime: user.AvailableTimeSlots,
				IsAvailable:           utils.Btoi(isOpen),
				Reason:                reason,
			}
		}
	}

	// Do something with the user object
	resp := GetBusinessResponse{
		Code:    0,
		Message: "Get businesses succeeded",
		Data:    businessList}
	return c.JSON(http.StatusOK, resp)
}

func GetBusinessByName(c echo.Context) error {
	cc := c.(*context.AppContext)

	var getBusinessByNameRequest GetBusinessByNameRequest

	err := cc.Bind(&getBusinessByNameRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Based on what condition is given, we go for different query logic
	queryCondition := make(map[string]any)
	// For cases when BusinessName is given, we use UserName to match BusinessName
	if len(getBusinessByNameRequest.BusinessName) != 0 {
		queryCondition["Username"] = "%" + getBusinessByNameRequest.BusinessName + "%"
	}
	if len(getBusinessByNameRequest.Zipcode) != 0 {
		queryCondition["Zipcode"] = getBusinessByNameRequest.Zipcode
	}
	if len(getBusinessByNameRequest.State) != 0 {
		queryCondition["AddressState"] = getBusinessByNameRequest.State
	}
	if len(getBusinessByNameRequest.City) != 0 {
		queryCondition["AddressCity"] = getBusinessByNameRequest.City
	}
	businesses := []models.User{}
	// For cass when zipcode is enabled, we search based on zipcode,
	// the same logic applies to city and state

	_, err = cc.UserStore.RetrieveManyWithCondition(&queryCondition, &businesses)

	if err != nil {
		b := errors.NewBoom(errors.EntityUpdateError, errors.ErrorText(errors.EntityUpdateError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	businessList := make([]Business, getBusinessByNameRequest.Quantity)

	for i, user := range businesses {
		// check open

		isOpen := utils.CheckBusinessOpenRightNow(&user)
		if !isOpen && getBusinessByNameRequest.Open != 0 {
			continue
		}

		// return a list of busispace to store ness
		// We could leave an extra table to store reason and available time
		reason := "No Reason"
		if !isOpen {
			reason = "Not in operation time."
		}

		if i >= getBusinessByNameRequest.StartNum {
			businessList[i-getBusinessByNameRequest.StartNum] = Business{
				Uid:              user.ID,
				BusinessName:     user.Username,
				BusinessPhoneNum: user.PhoneNum,
				BusinessLocation: GeologicalLocation{
					Lat: user.Lat,
					Lon: user.Lon,
				},
				BusinessAvailableTime: user.AvailableTimeSlots,
				IsAvailable:           utils.Btoi(isOpen),
				Reason:                reason,
			}
		}
	}

	// Do something with the user object
	resp := GetBusinessByNameResponse{
		Code:    0,
		Message: "Get business by name succeeded",
		Data:    businessList,
	}
	return c.JSON(http.StatusOK, resp)
}

// RegisterBusinessRoutes registers the authentication routes with the provided router.
func RegisterBusinessRoutes(server *core.Server) {

	g := server.Echo.Group("/api")
	g.POST("/business/Getusiness", GetBusiness)
	g.POST("/business/GetBusinessByName", GetBusinessByName)
}
