/*Created by Arjun Ajith on 21/02/2017*/

package models

type Users struct {
	Info			UserInfo
	NextOfKin		NextOfKin
	Settings		UserSettings
	SocialNetworks		UserSocialNetworks
	Tasks			map[string]UserTasks
	Company          	map[string]UsersCompany
	ConsentReceipts          map[string]ConsentReceiptDetails
}
type ConsentReceiptDetails struct {
	UserResponse		string
	CompanyId		string

}
type UserInfo struct {
	//Address			string
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