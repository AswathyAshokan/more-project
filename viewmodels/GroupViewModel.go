package viewmodels

type GroupList struct {
	Values       [][]string
	Keys         []string
	CompanyTeamName		string
}

type AddGroupViewModel struct {
	GroupMembers 		[]string
	GroupKey     		[]string
	PageType		string
	CompanyTeamName		string
}

type EditGroupViewModel struct {
	GroupMembers    	[]string
	GroupKey        	[]string
	GroupMembersToEdit	[]string
	PageType        	string
	GroupNameToEdit 	string
	GroupId                 string
	CompanyTeamName		string
}