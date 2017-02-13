/*Created By Farsana*/
package models


import (
	"golang.org/x/net/context"
	"log"

)

type Customer struct {

	CustomerName 	 string
	ContactPerson	 string
	Address 	 string
	Phone		 string
	Email 		 string
	State		 string
	ZipCode		 string
}

// Add new customers to database
func(m *Customer) AddCustomersToDb(ctx context.Context) (bool){
	//log.Println("values in model",this)
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
func(m *Customer) GetAllCustomerDetails(ctx context.Context) map[string]Customer{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	values := map[string]Customer{}
	err = db.Child("Customer").Value(&values)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("%s\n", v)
	//log.Println(reflect.TypeOf(v))
	return values


}

// delete customer from database using customerid
func(m *Customer) DeleteCustomerById(ctx context.Context,customerKey string) bool{
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
func(m *Customer) EditCustomer(ctx context.Context,customerId string) (Customer,bool){

	value := Customer{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+customerId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}

//update the customer profile
func(m *Customer) UpdateCustomerDetailsById(ctx context.Context,customerId string) (bool) {


	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+ customerId).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}








