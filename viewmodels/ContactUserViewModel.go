package viewmodels
import (
	"app/passporte/models"
)

type ContactUserViewModel  struct {
	User        	[]models.ContactUser
	Name        	string
	Address     	string
	State      	 string
	ZipCode     	string
	Email      	 string
	PhoneNumber 	string
	CurrentDate 	int64
	Status      	string
	PageType	string
	Key         	[]string
	ContactId	string

}