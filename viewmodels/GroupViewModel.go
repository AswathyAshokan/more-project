package viewmodels

import (
	"app/passporte/models"
)
type Group struct {
	Groups       	[]models.Group
	GroupName    	string
	GroupMembers 	[]string
	GroupKey     	[]string
	PageType	string
	SelectedUser    string
}

//

type EditGroupViewModel struct {
	GroupMembers 	[]string
	GroupKey     	[]string
	PageType	string
	GroupName	string
	GroupId		string
}