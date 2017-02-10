package viewmodels
import (
	"app/passporte/models"
)
type TaskViewModel  struct {
	Task			[]models.Task
	Key			[]string
	JobNameArray	[]string
	ContactNameArray	[]string
	PageType		string
	JobName		string
	TaskName		string
	TaskLocation		string
	StartDate		string
	EndDate			string
	LoginType		string
	Status			string
	TaskDescription		string
	UserNumber		string
	Log			string
	UserType		string
	Contact			string
	FitToWork		string
	TaskId			string
	GroupNameArray		[]string

}