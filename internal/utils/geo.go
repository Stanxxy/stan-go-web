package utils

import (
	"math"
	"sort"

	"github.com/Stanxxy/stan-go-web/internal/models"
)

type Location struct {
	Lat float64
	Lon float64
}
type Closness struct {
	Item     *models.User
	Distance float64
}

func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	R := 6371.0 // Earth's radius in kilometers

	// Convert coordinates to radians
	lat1Rad := lat1 * math.Pi / 180.0
	lon1Rad := lon1 * math.Pi / 180.0
	lat2Rad := lat2 * math.Pi / 180.0
	lon2Rad := lon2 * math.Pi / 180.0

	// Calculate differences in coordinates
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Apply Haversine formula
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c

	return distance
}

func SortLocationsBasedOnCloseness(referencePoint Location, comparatees []models.User) []*models.User {
	closnessList := make([]Closness, len(comparatees))

	for i, comparatee := range comparatees {
		currentLoc := Location{Lat: comparatee.Lat, Lon: comparatee.Lon}
		distance := HaversineDistance(referencePoint.Lat, referencePoint.Lon, currentLoc.Lat, currentLoc.Lon)
		closnessList[i] = Closness{Item: &comparatee, Distance: distance}
	}

	sort.Slice(closnessList, func(i, j int) bool {
		return closnessList[i].Distance < closnessList[j].Distance
	})

	ans := make([]*models.User, len(closnessList))

	for i, closness := range closnessList {
		ans[i] = closness.Item
	}
	return ans
}
