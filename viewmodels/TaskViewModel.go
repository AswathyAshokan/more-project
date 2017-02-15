package viewmodels

type EditTaskViewModel  struct {

	JobNameArray		[]string
	JobCustomerNameArray	[]string
	ContactNameArray	[]string
	Key			[]string
	GroupNameArray		[]string

	PageType		string
	JobName			string
	TaskName		string
	TaskLocation		string
	StartDate		string
	EndDate			string
	LoginType		string
	Status			string
	TaskDescription		string
	UserNumber		string
	Log			string
	UserType		[]string
	Contact			string
	FitToWork		string
	TaskId			string
}

type AddTaskViewModel  struct {
	JobNameArray		[]string
	JobCustomerNameArray	[]string
	ContactNameArray	[]string
	Key			[]string
	GroupNameArray		[]string


}
type TaskDetailViewModel  struct {
	Values            	[][]string
	Keys              	[]string
}