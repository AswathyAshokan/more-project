package viewmodels

import (
	//"app/passporte/models"
)
type AddLocationViewModel struct {
	GroupNameArray			[]string
	UserAndGroupKey			[]string
	GroupMembers			[][]string
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	CompanyTeamName			string

}
type LoadWorkLocationViewModel struct {
	Values            		[][]string
	Keys              		[]string
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	CompanyTeamName			string

}

