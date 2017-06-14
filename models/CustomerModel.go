/*Created By Farsana*/
package models


import (
	"golang.org/x/net/context"
	"log"
	"app/passporte/helpers"
	"reflect"
)

type Customers struct {
	Info     	CustomerData
	Settings 	CustomerSettings
	Tasks		map[string] TasksCustomer

}
type TasksCustomer struct {
	TasksCustomerStatus	string
}

type CustomerData struct {
	CustomerName		string
	ContactPerson		string
	Address			string
	Phone			string
	Email			string
	State			string
	ZipCode			string
	CompanyTeamName		string

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
	_,err = db.Child("Customers").Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}

// Fetch all the details of customer from database
func GetAllCustomerDetails(ctx context.Context,companyTeamName string) (map[string]Customers,bool){
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	allCustomerDetails := map[string]Customers{}
	err = db.Child("Customers").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&allCustomerDetails)
	if err != nil {
		log.Fatal(err)
		return allCustomerDetails,false
	}
	return allCustomerDetails,true
}

// delete customer from database using customer id
func(m *Customers) DeleteCustomerFromDBForNonTask(ctx context.Context,customerKey string) bool{
	log.Println("id",customerKey)
	customerSettingsUpdation := CustomerSettings{}
	customerDeletion := CustomerSettings{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Customers/"+ customerKey+"/Settings").Value(&customerSettingsUpdation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	customerDeletion.Status = helpers.UserStatusDeleted
	customerDeletion.DateOfCreation = customerSettingsUpdation.DateOfCreation
	err = db.Child("Customers/"+customerKey+"/Settings").Update(&customerDeletion)
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
	err = db.Child("/Customers/"+customerId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true
}

//update the customer profile
func(m *Customers) UpdateCustomerDetailsById(ctx context.Context,customerId string) (bool) {
	db,err :=GetFirebaseClient(ctx,"")
	customerSettingsDetails := CustomerSettings{}
	err = db.Child("/Customers/"+ customerId+"/Settings").Value(&customerSettingsDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	m.Settings.Status = customerSettingsDetails.Status
	m.Settings.DateOfCreation = customerSettingsDetails.DateOfCreation
	err = db.Child("/Customers/"+ customerId).Update(&m)
	if err != nil {
		log.Fatal(err)
		return  false
	}



	//....updation in task
	customerDetailForUpdation := map[string]Tasks{}
	taskCustomerForUpdate :=TaskCustomer{}
	taskCustomerDetail :=TaskCustomer{}

	err = db.Child("/Tasks/").Value(&customerDetailForUpdation)
	dataValue := reflect.ValueOf(customerDetailForUpdation)
	for _, key := range dataValue.MapKeys() {

		if customerDetailForUpdation[key.String()].Customer.CustomerId ==customerId{

			err = db.Child("Tasks/" + key.String()+"/Customer/").Value(&taskCustomerDetail)
			taskCustomerForUpdate.CustomerId =taskCustomerDetail.CustomerId
			taskCustomerForUpdate.CustomerName =m.Info.CustomerName
			taskCustomerForUpdate .CustomerStatus =taskCustomerDetail.CustomerStatus
			err = db.Child("Tasks/" + key.String()+"/Customer/").Update(&taskCustomerForUpdate)

		}
	}
	return true
}

//check customer name is already exist
func IsCustomerNameUsed(ctx context.Context,customerName string)(bool) {
	customerDetails := map[string]Customers{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = db.Child("Customers").OrderBy("Info/CustomerName").EqualTo(customerName).Value(&customerDetails)
	if err != nil {
		log.Fatal(err)
	}
	if len(customerDetails)==0{
		return true
	} else {
		dataValue := reflect.ValueOf(customerDetails)
		for _, key := range dataValue.MapKeys() {
			if customerDetails[key.String()].Settings.Status == helpers.StatusActive {
				return false
			}
		}

	}
	return true


}


func (m *Customers) DeleteCustomerFromDB(ctx context.Context, customerId string,TaskSlice []string)(bool)  {

	customerDetailForUpdate :=TasksCustomer{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	customerDetailForUpdate.TasksCustomerStatus =helpers.StatusInActive
	for i:=0;i<len(TaskSlice);i++{
		log.Println(TaskSlice[i])
		err = dB.Child("/Customers/"+ customerId+"/Tasks/"+TaskSlice[i]).Update(&customerDetailForUpdate)

	}
	taskCustomerDetail :=TaskCustomer{}
	taskCustomerForUpdate :=TaskCustomer{}
	for i:=0;i<len(TaskSlice);i++ {
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Customer/").Value(&taskCustomerDetail)
		log.Println("details from task job",)
		taskCustomerForUpdate.CustomerStatus =helpers.StatusInActive
		taskCustomerForUpdate.CustomerId =taskCustomerDetail.CustomerId
		taskCustomerForUpdate.CustomerName =taskCustomerDetail.CustomerName

		log.Println("fhsgjs",taskCustomerForUpdate)
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Customer/").Update(&taskCustomerForUpdate)

	}
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}
func (m *TasksCustomer) IsCustomerUsedForTask( ctx context.Context, customerId string)(bool,map[string]TasksCustomer)  {
	customerDetail := map[string]TasksCustomer{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Customers/"+ customerId+"/Tasks/").Value(&customerDetail)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false,customerDetail
	}
	log.Println(customerDetail)

	return true,customerDetail
}
//func (m *Customers) DeleteCustomerFromDBForNonTask(ctx context.Context, customerId string)(bool) {
//	customerDetail := Customers{}
//	updatedCustomerDetail :=Customers{}
//	log.Println("gggg")
//
//	dB, err := GetFirebaseClient(ctx,"")
//	err = dB.Child("/Customers/"+ customerId).Value(&customerDetail)
//	updatedCustomerDetail.Settings.DateOfCreation =customerDetail.Settings.DateOfCreation
//	updatedCustomerDetail.Settings.Status =helpers.StatusInActive
//	updatedCustomerDetail.Info.CustomerName=customerDetail.Info.CustomerName
//	updatedCustomerDetail.Info.CompanyTeamName =customerDetail.Info.CompanyTeamName
//	updatedCustomerDetail.Info.Address =customerDetail.Info.Address
//	updatedCustomerDetail.Info.ContactPerson =customerDetail.Info.CustomerName
//	updatedCustomerDetail.Info.Email =customerDetail.Info.Email
//	updatedCustomerDetail.Info.Phone =customerDetail.Info.Phone
//	updatedCustomerDetail.Info.State =customerDetail.Info.State
//	updatedCustomerDetail.Info.ZipCode =customerDetail.Info.ZipCode
//
//	log.Println("dfkfj",updatedCustomerDetail)
//
//	err = dB.Child("/Customers/"+ customerId).Update(&updatedCustomerDetail)
//	if err != nil {
//		log.Fatal(err)
//		return false
//	}
//	return true
//}
//

