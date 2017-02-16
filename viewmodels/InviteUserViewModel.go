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
}

type EditInviteUserViewModel struct {
	FirstName      		string
	LastName      		string
	EmailId        		string
	UserType      		string
	Status         		string
	PageType        	string
	InviteId        	string
}
