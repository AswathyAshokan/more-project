package viewmodels

type InviteUserViewModel struct {
	FirstName      		string
	LastName      		string
	EmailId        		string
	UserType      		string
	Status         		string
	DateOfCreation 		int64
	InviteUserKey  		[]string
	PageType        	string
	InviteId        	string
	Values           	[][]string
	Keys             	[]string
	CompanyTeamName		string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
}

type EditInviteUserViewModel struct {
	FirstName      		string
	LastName      		string
	EmailId        		string
	UserType      		string
	Status         		string
	PageType        	string
	InviteId        	string
	CompanyTeamName		string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string

}
type AddInviteUserViewModel struct {
	CompanyTeamName		string
	CompanyPlan		string
	AllowInvitations	bool
	AdminFirstName		string
	AdminLastName		string
}
