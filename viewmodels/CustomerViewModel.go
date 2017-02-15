package viewmodels

type Customer struct {

	CustomerName  string
	ContactPerson string
	Address       string
	Phone         string
	Email         string
	State         string
	ZipCode       string
	PageType      string
	CustomerId    string
	Values	      [][]string
	Keys	      []string
}
type ListJobDetailsOfCustomer struct {

	CustomerName      	string
	JobName           	string
	JobNumber         	string
	NumberOfTask      	string
	Status                  string
	CustomerId              string
}