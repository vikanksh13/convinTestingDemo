package serviceimpl

import (
	"convinTestingDemo/models"
	"errors"
	"math"
	"sort"
)

type NearestTestingCenter struct {
	CenterLocation  float64
	TestingCenterId string
}

func (tc *TestingCenters) BookTest(user *models.User, numberOfCenter int) (bool, error) {
	if user == nil {
		return false, errors.New("no user found")
	}

	if user.Booked {
		return false, errors.New("user is already booked the slot")
	}

	nearestTestingCenters, err := tc.getNearestTestingCenters(user, numberOfCenter)
	if err != nil {
		return false, err
	}

	var isBooked bool
	for _, center := range nearestTestingCenters {
		if center.NumberOfKits > 0 {
			center.NumberOfKits--
			user.Booked = true
			isBooked = true
			break
		}
	}

	if !isBooked {
		for _, center := range nearestTestingCenters {
			center.WaitList = make(map[string]models.User)
			center.WaitList[user.UserId] = *user
		}
	}

	return isBooked, nil
}

func (tc *TestingCenters) getNearestTestingCenters(user *models.User, numberOfCenter int) ([]*models.TestingCenter, error) {

	var nearestCenter []NearestTestingCenter
	var availableCenters []*models.TestingCenter

	for centerId, center := range tc.TestingCenters {
		xCenter := center.XCoordinate
		yCenter := center.YCoordinate
		xUser := user.UserXCoordinate
		yUser := user.UserYCoordinate

		xVal := (float64)((xCenter - xUser) * (xCenter - xUser))
		yVal := (float64)((yCenter - yUser) * (yCenter - yUser))
		val := xVal + yVal

		nearestCenter[centerId].CenterLocation = math.Sqrt(val)
		nearestCenter[centerId].TestingCenterId = center.TestingCenterId
	}

	sort.SliceStable(nearestCenter, func(i, j int) bool {
		return nearestCenter[i].CenterLocation < nearestCenter[j].CenterLocation
	})

	for numberOfCenter > 0 {
		numberOfCenter--
		if entry, ok := tc.TestingCenterMap[nearestCenter[numberOfCenter].TestingCenterId]; ok {
			availableCenters = append(availableCenters, &entry)
		} else {
			return nil, errors.New("test center not found")
		}

	}

	return availableCenters, nil
}
