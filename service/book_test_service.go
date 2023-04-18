package service

import "convinTestingDemo/models"

type IBookTestService interface {
	BookTest(user *models.User, numberOfCenter int) (bool, error)
}
