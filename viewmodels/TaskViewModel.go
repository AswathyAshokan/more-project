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
	FitToWork			[]string
	TaskId				string
	ContactNameToEdit		[]string
	CompanyTeamName			string
	GroupMembersAndUserToEdit	[]string
	UsersToEdit			[]string
	GroupsToEdit			[]string
	GroupMembers			[][]string
	CompanyPlan			string
	StartTime			string
	EndTime				string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	FitToWorkCheck			string
	JobId				string


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
	GroupMembers			[][]string
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string


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
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	NoCustomer			string
	JobMatch			string
	UserArray			[][]string
}