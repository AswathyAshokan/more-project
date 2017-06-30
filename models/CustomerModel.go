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


	//...update on job

	customerDetailForJobUpdation := map[string]Job{}
	taskCustomerForJobUpdate :=JobCustomer{}
	taskJobCustomerDetail :=JobCustomer{}

	err = db.Child("/Jobs/").Value(&customerDetailForJobUpdation)
	dataJobValue := reflect.ValueOf(customerDetailForJobUpdation)
	for _, key := range dataJobValue.MapKeys() {

		if customerDetailForJobUpdation[key.String()].Customer.CustomerId ==customerId{

			err = db.Child("Jobs/" + key.String()+"/Customer/").Value(&taskJobCustomerDetail)
			taskCustomerForJobUpdate.CustomerId =taskJobCustomerDetail.CustomerId
			taskCustomerForJobUpdate.CustomerName =m.Info.CustomerName
			err = db.Child("Jobs/" + key.String()+"/Customer/").Update(&taskCustomerForUpdate)

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


func (m *Customers) DeleteCustomerFromDB(ctx context.Context, customerId string,TaskSlice []string,companyTeamName string)(bool)  {

	customerDetailForUpdate :=TasksCustomer{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	customerDetailForJob := map[string]Job{}
	err = dB.Child("/Jobs/").Value(&customerDetailForJob)
	if len(customerDetailForJob) != 0 {
		dataValue := reflect.ValueOf(customerDetailForJob)
		for _, key := range dataValue.MapKeys() {
			if customerDetailForJob[key.String()].Customer.CustomerId == customerId {
				customerSettingsUpdation := CustomerSettings{}
				customerDeletion := CustomerSettings{}
				customerUpdationInJob := JobCustomer{}
				customerDeletionInJob := JobCustomer{}

				db,err :=GetFirebaseClient(ctx,"")
				err = db.Child("/Customers/"+ customerId+"/Settings").Value(&customerSettingsUpdation)
				if err != nil {
					log.Fatal(err)
					return  false
				}
				customerDeletion.Status = helpers.UserStatusDeleted
				customerDeletion.DateOfCreation = customerSettingsUpdation.DateOfCreation
				err = db.Child("Customers/"+customerId+"/Settings").Update(&customerDeletion)
				if err != nil {
					log.Fatal(err)
					return  false
				}
				err = db.Child("Jobs/"+key.String()+"/Customer").Value(&customerUpdationInJob)
				customerDeletionInJob.CustomerId =customerUpdationInJob.CustomerId
				customerDeletionInJob.CustomerName =customerUpdationInJob.CustomerName
				customerDeletionInJob.CustomerStatus =helpers.StatusInActive
				err = db.Child("Jobs/"+key.String()+"/Customer").Update(&customerDeletionInJob)
			}
		}

		customerDetailForUpdation := map[string]Tasks{}
		taskCustomerForUpdate :=TaskCustomer{}
		taskCustomerDetail :=TaskCustomer{}
		taskUpdate := TaskSetting{}
		taskDeletion :=TaskSetting{}
		taskDetailForUser :=Tasks{}

		err = dB.Child("/Tasks/").Value(&customerDetailForUpdation)
		dataValueOfTask := reflect.ValueOf(customerDetailForUpdation)
		for _, taskKey := range dataValueOfTask.MapKeys() {

			if customerDetailForUpdation[taskKey.String()].Customer.CustomerId ==customerId{
				log.Println("inside deletion of customer from taskkkkkkkk")

				err = dB.Child("Tasks/" + taskKey.String()+"/Customer/").Value(&taskCustomerDetail)
				taskCustomerForUpdate.CustomerId =taskCustomerDetail.CustomerId
				taskCustomerForUpdate.CustomerName =taskCustomerDetail.CustomerName
				taskCustomerForUpdate .CustomerStatus =helpers.StatusInActive

				err = dB.Child("Tasks/" + taskKey.String()+"/Customer/").Update(&taskCustomerForUpdate)
				taskDeletion.Status =helpers.StatusInActive
				err = dB.Child("/Tasks/"+ taskKey.String()+"/Settings").Value(&taskUpdate)
				taskDeletion.DateOfCreation =taskUpdate.DateOfCreation
				err = dB.Child("/Tasks/"+ taskKey.String()+"/Settings").Update(&taskDeletion)
				err = dB.Child("/Tasks/"+ taskKey.String()).Value(&taskDetailForUser)
				userData := reflect.ValueOf(taskDetailForUser.UsersAndGroups.User)
				for _, key := range userData.MapKeys() {
					userTaskDetail := UserTasks{}
					userTaskDetail.DateOfCreation = taskDetailForUser.Settings.DateOfCreation
					userTaskDetail.TaskName = taskDetailForUser.Info.TaskName
					userTaskDetail.CustomerName = taskDetailForUser.Customer.CustomerName
					userTaskDetail.EndDate = taskDetailForUser.Info.EndDate
					userTaskDetail.StartDate = taskDetailForUser.Info.StartDate
					userTaskDetail.JobName = taskDetailForUser.Job.JobName
					userTaskDetail.Status = helpers.StatusInActive
					userTaskDetail.CompanyId = companyTeamName
					userKey := key.String()
					err = dB.Child("/Users/" + userKey + "/Tasks/" + taskKey.String()).Update(&userTaskDetail)
					if err!=nil{
						log.Println("Deletion error:",err)
					}
				}
				log.Println("deleted successfully")

			}
		}
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
	customerDetailForJob := map[string]Job{}
	customerUpdationInJob := JobCustomer{}
	customerDeletionInJob := JobCustomer{}
	err = dB.Child("/Jobs/").Value(&customerDetailForJob)
	if len(customerDetailForJob) != 0 {
		dataValue := reflect.ValueOf(customerDetailForJob)
		for _, key := range dataValue.MapKeys() {
			if customerDetailForJob[key.String()].Customer.CustomerId == customerId {
				err = dB.Child("Jobs/"+key.String()+"/Customer").Value(&customerUpdationInJob)
				customerDeletionInJob.CustomerId =customerUpdationInJob.CustomerId
				customerDeletionInJob.CustomerName =customerUpdationInJob.CustomerName
				customerDeletionInJob.CustomerStatus =helpers.StatusInActive
				err = dB.Child("Jobs/"+key.String()+"/Customer").Update(&customerDeletionInJob)
			}
		}
	}



	return true,customerDetail
}
func (m *Job) IsCustomerUsedForJob( ctx context.Context, customerId string,companyTeamName string)(bool,map[string]Job)  {
	jobDetail := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	//jobStatus := "Active"

	err = dB.Child("Jobs").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&jobDetail)

	if err != nil {
		log.Fatal(err)
		return false, jobDetail
	}
	return true, jobDetail
	log.Println(jobDetail)

	return true,jobDetail
}