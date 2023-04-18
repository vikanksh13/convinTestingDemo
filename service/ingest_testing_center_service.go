package service

import "convinTestingDemo/models"

type IIngestTestingCenterService interface {
	IngestTestingCenter(tc []models.TestingCenter)
	UpdateNumberOfKits(kits int, tc models.TestingCenter) []*models.User
}
