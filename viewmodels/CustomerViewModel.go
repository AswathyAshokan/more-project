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

}

type AddCustomerViewModel struct {
	CompanyTeamName		string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
}