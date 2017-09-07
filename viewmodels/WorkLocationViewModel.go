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
	Users 				[][]WorkLocationUsers

	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	CompanyTeamName			string

}
type WorkLocationUsers struct {
	Name	string
	UserKey string
}

type EditWorkLocation struct {
	WorkLogId 		string
	GroupNameArray 		[]string
	UsersKey     		[]string
	UserNameToEdit		[]string
	WorkLocation    	string
	PageType        	string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
	ProfilePicture		string
	CompanyTeamName		string
	GroupMembers		[][]string
	UserAndGroupKey		[]string
	StartDate		int64
	EndDate 		int64
	DailyStartTime		string
	DailyEndTime 		string
}
