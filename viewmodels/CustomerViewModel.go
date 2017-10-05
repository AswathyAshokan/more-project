package viewmodels

type Customer struct {

	CustomerName  	string
	ContactPerson 	string
	Address       	string
	Phone         	string
	Email         	string
	State        	string
	ZipCode       	string
	CustomerId    	string
	Values	      	[][]string
	Keys	      	[]string
	CompanyTeamName	string
	CompanyPlan	string
	AdminFirstName	string
	AdminLastName	string
	ProfilePicture	string
	Country		string
	NotificationArray	[][]string
	NotificationNumber       int
}
type EditCustomerViewModel struct {

	CustomerName  	string
	ContactPerson 	string
	Address       	string
	Phone         	string
	Email         	string
	State        	string
	ZipCode       	string
	PageType      	string
	CustomerId    	string
	CompanyTeamName	string
	CompanyPlan	string
	AdminFirstName	string
	AdminLastName	string
	ProfilePicture	string
	Country		string
	NotificationArray	[][]string
	NotificationNumber       int

}

type AddCustomerViewModel struct {
	CompanyTeamName		string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
	ProfilePicture		string
	NotificationArray	[][]string
	NotificationNumber       int
	DocumentExpiryNotification [][]string
}