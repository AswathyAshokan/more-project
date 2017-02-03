
/* Author : Aswathy Ashok */

package viewmodels
import (
	"app/passporte/models"
)

type ContactUserViewModel  struct {
	User        []models.ContactUser
	Name        string
	Address     string
	State       string
	ZipCode     string
	Email       string
	PhoneNumber string
	CurrentDate int64
	Status      string
	Key         []string

}

type ProjectViewModel  struct {

	Project		[]models.Project
	CustomerName	string
	ProjectName	string
	ProjectNumber	string
	NumberOfTask	string
	Status		string
	CurrentDate	int64
	Key		[]string
}

type TaskViewModel  struct {
	Task		[]models.Task
	Key		[]string

}