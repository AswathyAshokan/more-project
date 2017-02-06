package viewmodels

import (
	"app/passporte/models"
)

type InviteUserViewModel struct {
	Users          []models.InviteUser
	FirstName      string
	LastName       string
	EmailId        string
	UserType       string
	Status         string
	DateOfCreation int64
	InviteUserKey  []string
}
