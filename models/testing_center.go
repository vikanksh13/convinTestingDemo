package models

type TestingCenter struct {
	TestingCenterId string          `json:"testing_center_id"`
	XCoordinate     int             `json:"x_coordinate"`
	YCoordinate     int             `json:"y_coordinate"`
	NumberOfKits    int             `json:"number_of_kits"`
	WaitList        map[string]User `json:"wait_list"`
}
