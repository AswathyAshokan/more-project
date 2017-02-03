package viewmodels

import (
	"app/passporte/models"
)
type Customer struct {
	Customers 	[]models.Customer
	CustomerName  string
	ContactPerson string
	Address       string
	Phone         string
	Email         string
	State         string
	ZipCode       string
	Key 		[]string

}