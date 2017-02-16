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
}