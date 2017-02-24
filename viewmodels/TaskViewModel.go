package viewmodels

type EditTaskViewModel  struct {

	JobNameArray			[]string
	JobCustomerNameArray		[]string
	ContactNameArray		[]string
	ContactKey			[]string
	Key				[]string
	UserAndGroupKey			[]string
	GroupNameArray			[]string
	PageType			string
	JobName				string
	TaskName			string
	TaskLocation			string
	StartDate			string
	EndDate				string
	LoginType			string
	Status				string
	TaskDescription			string
	UserNumber			string
	Log				string
	UserType			[]string
	Contact				string
	FitToWork			string
	TaskId				string
	ContactNameToEdit		[]string
	CompanyTeamName			string
}

type AddTaskViewModel  struct {
	JobNameArray			[]string
	JobCustomerNameArray		[]string
	ContactNameArray		[]string
	Key				[]string
	GroupNameArray			[]string
	UserAndGroupKey			[]string
	ContactKey			[]string
	CompanyTeamName			string


}

type TaskDetailViewModel  struct {
	Values            		[][]string
	Keys              		[]string
	UniqueCustomerNames		[]string
	UniqueJobNames			[]string
	SelectedJob			string
	SelectedCustomer		string
	SelectedCustomerForJob		string
	CompanyTeamName			string
}