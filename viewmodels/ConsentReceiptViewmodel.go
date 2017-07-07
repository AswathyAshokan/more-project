package viewmodels
type ConsentReceipt struct {
	GroupKey       		[]string
	GroupMembers		[]string
	CompanyTeamName		string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
	ProfilePicture		string
}

type LoadConsent struct {
	Values                  [][]string
	Keys			[]string
	InnerContent            []ConsentStruct
	CompanyTeamName         string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
	ProfilePicture		string

}
type ConsentStruct struct {
	Description   	string
	AcceptedUsers 	[]string
	RejectedUsers 	[]string
	PendingUsers  	[]string
	InstructionKey  string

}
type EditConsentReceipt struct {
	GroupKey       			[]string
	GroupMembers			[]string
	ReceiptName    			string
	ConsentKey       		[]string
	ConsentMembers			[]string
	CompanyTeamName			string
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	PageType    			string
	UserNameToEdit         		[]string
	InstructionArrayToEdit  	[]string
	ConsentId   			string
	UsersForEdit 			[]string
}
