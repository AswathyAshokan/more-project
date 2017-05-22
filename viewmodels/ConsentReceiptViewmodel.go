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
	InnerContent            [][]ConsentStruct
	CompanyTeamName          string

}
type ConsentStruct struct {
	InstructionArray        []string
	AcceptedUsers           []string
	RejectedUsers		[]string
	UserKey 		string
	Status 			string

}
type EditConsentReceipt struct {
	ReceiptName    		string
	GroupKey       		[]string
	GroupMembers		[]string
	CompanyTeamName		string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
	ProfilePicture		string
	PageType    		string
	UserNameToEdit         	[]string
	InstructionArrayToEdit  []string
	ConsentId   		string
}
