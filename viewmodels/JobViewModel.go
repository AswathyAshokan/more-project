package viewmodels
import (
	"app/passporte/models"
)
type JobViewModel  struct {

	Job			[]models.Job
	CustomerNameArray 	[]string
	CustomerName		string
	JobName			string
	JobNumber		string
	NumberOfTask		string
	Status			string
	CurrentDate		int64
	PageType		string
	Key			[]string
}