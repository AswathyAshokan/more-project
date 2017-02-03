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
func(this *Customer) AddToDb(ctx context.Context) {
	//log.Println("values in model",this)
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("Customer").Push(this)
	if err != nil {
		log.Println(err)
	}
}
// Display details


func(this *Customer) DisplayCustomer(ctx context.Context) map[string]Customer{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	v := map[string]Customer{}
	err = db.Child("Customer").Value(&v)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("%s\n", v)
	//log.Println(reflect.TypeOf(v))
	return v


}

// delete customer

func(this *Customer) DeleteCustomer(ctx context.Context,key string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+key).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

//edit a record

func(this *Customer) EditCustomer(ctx context.Context,key string) (Customer,bool){
	value := Customer{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customer/"+key).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}







