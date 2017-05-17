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
	//InstructionArray	[][]string
	//ReceiptName             string
	Values                  [][]string
	Keys			[]string
	InnerContent            [][]ConsentStruct
	CompanyTeamName          string

}
type ConsentStruct struct {
	InstructionArray        []string
	UserName     		string
	UserKey 		string

}
