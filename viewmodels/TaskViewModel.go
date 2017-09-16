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
	WorkBreak			[]string
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
	WorkTime			[]string
	BreakTime			[]string
	NFCTagId			string
	FitToWorkName			string
	FitToWorkArray			[]string
	FitToWorkKey			[]string
	FitToWorkForTask		[][]TaskFitToWork
	ContactUser			[][]TaskContact
	CustomerNameToEdit		string
	ContactNameKeyToEdit		[]string
	JobNameFormUrl			string
	CustomerNameFormUrl		string
	WorkLocationForUser		[][]string


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
	FitToWorkArray			[]string
	FitToWorkKey			[]string
	FitToWorkForTask		[][]TaskFitToWork
	ContactUser			[][]TaskContact
	JobNameFormUrl			string
	CustomerNameFormUrl		string
	WorkLocationArray		[]string
	WorkLocationForUser		[][]string



}
type TaskContact struct {
	ContactName 	 string
	ContactId	string
	CustomerName	[]string
	CustomerId	[]string
}

type TaskUsers struct {
	TaskId 	string
	Name	string
	Status	string
}
type TaskMinUsers struct {
	TaskId 		string
	MinimumUser	string
	LoginType	string
}

type TaskDetailViewModel  struct {
	Values            			[][]string
	Keys              			[]string
	UniqueCustomerNames			[]string
	UniqueJobNames				[]string
	SelectedJob				string
	SelectedCustomer			string
	SelectedCustomerForJob			string
	CompanyTeamName				string
	CompanyPlan				string
	AdminFirstName				string
	AdminLastName				string
	ProfilePicture				string
	NoCustomer				string
	JobMatch				string
	UserArray				[][]TaskUsers
	MinUserAndLoginTypeArray		[][]string
	ExposureArray				[][]TaskExposure
}

type TaskFitToWork struct {
	FitToWorkName 	string
	Instruction	[]string
	InstructionKey	[]string

}
type TaskExposure struct {
	BreakMinute  	string
	WorkingHour	string
	TaskId 		string
}