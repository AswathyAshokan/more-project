package viewmodels

import (
	"app/passporte/models"
)

type UserViewModel struct {
	Users		[]models.InviteUser
	FirstName      string
	LastName       string
	EmailId        string
	UserType       string
	Status         string
	DateOfCreation int64
	Key 		[]string
}
