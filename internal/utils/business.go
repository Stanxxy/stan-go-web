package utils

import (
	"github.com/Stanxxy/stan-go-web/internal/models/user"
	"time"
)

// check if business is open at current time
func CheckBusinessOpenRightNow(business *User) bool {
	// get current time
	curTime := time.now()
	return CheckBusinessOpen(business, curTime)
}

// check if business is open at a given time
func CheckBusinessOpen(business *User, givenTime time.Time) bool {
	res := false
	for i, timeSlot := range business.AvailableTimeSlots {
		// first compare weekday
		if givenTime.Weekday().String() != timeSlot.WeekDay {
			continue
		}

		// TODO: we need to get time zone infor via user location
		// right now jsut hard coded as eastern time
		loc := *(FindUserTimeZone(business))
		
		startTime := time.Date(
			givenTime.Year(), 
			givenTime.Month(),
			givenTime.Day(),
			timeSlot.StartHour, // for hour
			timeSlot.StartMinute,
			0, 0, 0, 0, loc)
		endTime := time.Date(
			givenTime.Year(), 
			givenTime.Month(),
			givenTime.Day(),
			timeSlot.EndHour, // for hour
			timeSlot.EndMinute,
			0, 0, 0, 0, loc)
		// check logic
		if givenTime.Compare(startTime) >= 0 && givenTime.Compare(endTime) <= 0 {
			res = true
			break
		}
	}
	return res
}