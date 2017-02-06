package viewmodels
import (
	"app/passporte/models"
)
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