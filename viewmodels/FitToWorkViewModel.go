package viewmodels
type FitToWork struct {
	CompanyTeamName		string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
	ProfilePicture		string
	Values			[][]string
	InnerContent            []FitToWorkStruct
	Keys			[]string
	NotificationArray	[][]string
	NotificationNumber       int
}
type FitToWorkStruct struct {
	Description   		string
	InstructionKey  	string


}
type EditFitToWork struct {

	FitToWorkName    		string
	FitToWorkKey       		[]string
	FitToWorkMembers		[]string
	CompanyTeamName			string
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	PageType    			string
	UserNameToEdit         		[]string
	InstructionArrayToEdit  	[]string
	FitToWorkId   			string
	NotificationArray		[][]string
	NotificationNumber       	int

}