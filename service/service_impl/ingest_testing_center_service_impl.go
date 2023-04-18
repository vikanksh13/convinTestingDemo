package serviceimpl

import (
	"convinTestingDemo/models"
)

type TestingCenters struct {
	TestingCenters []models.TestingCenter `json:"testing_center"`
	// TestingCenterMap map[string]models.TestingCenter `json:"testing_center_map"`
}

var TestingCenterMap map[string]models.TestingCenter

func init() {
	TestingCenterMap = make(map[string]models.TestingCenter)
}

func (testingCenter *TestingCenters) IngestTestingCenter(tc []models.TestingCenter) {

	for _, center := range tc {
		TestingCenterMap[center.TestingCenterId] = center
	}

	testingCenter.TestingCenters = append(testingCenter.TestingCenters, tc...)
}

func (testingCenter *TestingCenters) UpdateNumberOfKits(kits int, tc models.TestingCenter) []*models.User {

	var userDetails []*models.User
	for userId, user := range tc.WaitList {
		if kits > 0 && len(tc.WaitList) > 0 {
			delete(tc.WaitList, userId)
			userDetails = append(userDetails, &user)
			kits--
		}
	}

	if entry, ok := TestingCenterMap[tc.TestingCenterId]; ok {
		entry.NumberOfKits = kits

		TestingCenterMap[tc.TestingCenterId] = entry
	}
	return userDetails
}
