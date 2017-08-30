/*Created by Arjun Ajith on 21/02/2017*/

package models

type Users struct {
	Info			UserInfo
	NextOfKin		NextOfKin
	Settings		UserSettings
	SocialNetworks		UserSocialNetworks
	Tasks			map[string]UserTasks
	Company          	map[string]UsersCompany
	ConsentReceipts         map[string]ConsentReceiptDetails
	Task			map[string]UserNotification
	Invitations		map[string]UserInvitations
	WorkLocation		map[string]WorkLocationInUser
	//Invitations        map[string]UserInvitations

}



type ConsentReceiptDetails struct {
	UserResponse		string
	CompanyId		string

}
type UserInfo struct {

	City      		string
	Country     		string
	DateOfBirth		int64
	DialCode		string
	Email			string
	FullName		string
	MedicalHistory		string
	MedicareNumber		string
	Medication		string
	State 			string
	TaxFileNumber 		string
	ZipCode 		string
	Phone			int64

}
type UsersCompany struct{
	CompanyName             string
	DateOfJoin		int64
	Status			string
}

type NextOfKin struct {
	KinEmail		string
	KinName			string
	KinPhone		string
	Relation		string
}

type UserSettings struct {
	AdminID			string
	DateOfCreation		int64
	ProfilePicture		string
	Status			string
	ThumbProfilePicture	string
	LastLatitude		float64
	LastLongitude		float64
}

type UserSocialNetworks struct {
	FacebookId		string
	SkypeId			string
}

type UserTasks struct {
	CompanyId		string
	CustomerName		string
	DateOfCreation		int64
	EndDate			int64
	Id 			string
	JobName			string
	StartDate		int64
	Status			string
	TaskName		string
}
type WorkLocationInUser struct {
	CompanyId		string
	DateOfCreation		int64
	WorkLocationForTask 	string
	StartDate               int64
	EndDate			int64
	DailyStartDate          int64
	DailyEndDate	        int64
	Latitude		string
	Longitude		string
	Status 			string
}
type UserNotification struct {
	TaskId		string
	TaskName	string
	Date		int64
	IsRead		bool
	IsViewed	bool
	Category	string
	Status		string
	IsDeleted       bool
}
type UserInvitations struct {
	Date		int64
	IsRead		bool
	IsViewed	bool
	CompanyAdmin	string
	CompanyName	string
	Category	string
	IsDeleted       bool
}