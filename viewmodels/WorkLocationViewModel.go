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
	DateValues   			[][]string
	NotificationArray		[][]string
	NotificationNumber       	int
	FitToWorkArray			[]string
	FitToWorkKey			[]string
	FitToWorkForTask		[][]TaskFitToWork

}
type LoadWorkLocationViewModel struct {
	InnerContent            	[]WorkLocationStruct
	Values            		[][]string
	Keys              		[]string
	Users 				[][]WorkLocationUsers
	ExposureArray			[][]WorkExposure
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	CompanyTeamName			string
	NotificationArray		[][]string
	NotificationNumber       	int
	MinUserAndLoginTypeArray	[][]string

}


type WorkExposure struct {
	BreakMinute  	string
	WorkingHour	string
	TaskId 		string
}


type WorkLocationStruct struct {
	Description   		string
	AcceptedUsers 		[]string
	RejectedUsers 		[]string
	PendingUsers  		[]string
	InstructionKey  	string


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
	LatitudeForEditing	string
	LongitudeForEditing	string
	DateValues		[][]string
	NotificationArray	[][]string
	NotificationNumber      int
	FitToWorkArray		[]string
	FitToWorkKey		[]string
	FitToWorkForTask	[][]TaskFitToWork
	FitToWork		[]string
	WorkBreak		[]string
	NFCTagId		string
	FitToWorkCheck		string
	LoginType		string
	BreakTime		[]string
	WorkTime		[]string
	LogInMin		int64
	FitToWorkName		string

}
