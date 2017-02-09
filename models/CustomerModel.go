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
func(this *Customer) AddCustomersToDb(ctx context.Context) (bool){
	//log.Println("values in model",this)
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("Customer").Push(this)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}
// Display details


func(this *Customer) DisplayCustomer(ctx context.Context) map[string]Customer{
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

// delete customer

func(this *Customer) DeleteCustomer(ctx context.Context,customerKey string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+customerKey).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

//edit a record



func(this *Customer) EditCustomer(ctx context.Context,customerKey string) (Customer,bool){

	value := Customer{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+customerKey).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}

func(this *Customer) UpdateCustomerDetails(ctx context.Context,customerKey string) (bool) {


	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+ customerKey).Update(&this)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}








