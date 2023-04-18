package main

import (
	"convinTestingDemo/models"
	serviceimpl "convinTestingDemo/service/service_impl"
)

func main() {
	testingCenter1 := models.TestingCenter{
		TestingCenterId: "1",
		XCoordinate:     2,
		YCoordinate:     5,
		NumberOfKits:    10,
	}
	testingCenter2 := models.TestingCenter{
		TestingCenterId: "2",
		XCoordinate:     2,
		YCoordinate:     5,
		NumberOfKits:    10,
	}

	testCenters := []models.TestingCenter{
		testingCenter1,
		testingCenter2,
	}

	user := models.User{
		UserId:          "43",
		UserXCoordinate: 2,
		UserYCoordinate: 5,
	}

	testingCenterSvc := serviceimpl.TestingCenters{TestingCenters: testCenters}
	// bookTestSvc := serviceimpl.NearestTestingCenter{}

	testingCenterSvc.TestingCenterMap = make(map[string]models.TestingCenter)
	testingCenterSvc.IngestTestingCenter(testCenters)

	testingCenterSvc.BookTest(&user, 2)
}
