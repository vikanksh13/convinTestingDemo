package models

type User struct {
	UserId          string `json:"user_id"`
	UserXCoordinate int    `json:"user_x_coordinate"`
	UserYCoordinate int    `json:"user_y_coordinate"`
	Booked          bool   `json:"booked"`
}
