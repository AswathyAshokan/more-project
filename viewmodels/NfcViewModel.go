package viewmodels

type NfcViewModel struct {
	Values		[][]string
	Keys 		[]string
	CompanyTeamName string
	CompanyPlan	string
	AdminFirstName	string
	AdminLastName	string
	ProfilePicture	string
	NotificationArray	[][]string
	NotificationNumber       int

}

type EditNfcViewModel struct {
	PageType	string
	NfcId		string
	CustomerName	string
	Site      	string
	Location 	string
	NFCNumber	string
	CompanyTeamName string
	CompanyPlan	string
	AdminFirstName	string
	AdminLastName	string
	ProfilePicture	string
	NotificationArray	[][]string
	NotificationNumber       int

}