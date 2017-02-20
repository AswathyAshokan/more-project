/*Created By Farsana*/
package models


import (
	"golang.org/x/net/context"
	"log"
)

type Customers struct {
	Info     CustomerData
	Settings CustomerSettings

}

type CustomerData struct {
	CustomerName 	 string
	ContactPerson	 string
	Address 	 string
	Phone		 string
	Email 		 string
	State		 string
	ZipCode		 string

}

type CustomerSettings struct {
	Status           string
	DateOfCreation   int64
}



// Add new customers to database
func(m *Customers) AddCustomersToDb(ctx context.Context) (bool){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("Customer").Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}

// Fetch all the details of customer from database
func GetAllCustomerDetails(ctx context.Context) (map[string]Customers,bool){
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	allCustomerDetails := map[string]Customers{}
	err = db.Child("Customer").Value(&allCustomerDetails)
	if err != nil {
		log.Fatal(err)
		return allCustomerDetails,false
	}
	return allCustomerDetails,true
}

// delete customer from database using customer id
func(m *Customers) DeleteCustomerById(ctx context.Context,customerKey string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+customerKey).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true
}

//get all the values of a customer using customer id for editing purpose
func(m *Customers) EditCustomer(ctx context.Context,customerId string) (Customers,bool){

	value := Customers{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+customerId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true
}

//update the customer profile
func(m *Customers) UpdateCustomerDetailsById(ctx context.Context,customerId string) (bool) {
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+ customerId).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true
}

//check customer name is already exist
func IsCustomerNameUsed(ctx context.Context,customerName string)(bool) {
	log.Println("customerName",customerName)
	customerDetails := map[string]Customers{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = db.Child("Customer").OrderBy("CustomerName").EqualTo(customerName).Value(&customerDetails)
	if err != nil {
		log.Fatal(err)
	}
	if len(customerDetails)==0{
		return true
	} else {
		return false
	}

}




