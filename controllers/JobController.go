
/* Author :Aswathy Ashok */

package controllers

import (

	"app/passporte/models"
	"app/passporte/viewmodels"
	"time"
	"reflect"
	"app/passporte/helpers"
	"log"
)

type JobController struct {
	BaseController
}
/*Add job details to DB*/
func (c *JobController)AddNewJob() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	viewModel := viewmodels.JobViewModel{}
	if r.Method == "POST" {
		job:=models.Job{}
		job.Customer.CustomerId = c.GetString("customerId")
		job.Customer.CustomerName = c.GetString("customerName")
		job.Customer.CustomerStatus =helpers.StatusActive
		job.Info.JobName = c.GetString("jobName")
		job.Info.JobNumber = c.GetString("jobNumber")
		job.Info.NumberOfTask = c.GetString("numberOfTask")
		job.Settings.DateOfCreation = time.Now().UnixNano() / int64(time.Millisecond)
		job.Settings.Status = helpers.StatusActive
		job.Info.CompanyTeamName = storedSession.CompanyTeamName
		dbStatus :=job.AddJobToDB(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}else {

		//Getting all customer details
		customers ,dbStatus:=models.GetAllCustomerDetails(c.AppEngineCtx,companyTeamName)
		var keySlice []string
		dataValue := reflect.ValueOf(customers)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(customers)
			for _, k := range dataValue.MapKeys() {
				if customers[k.String()].Settings.Status =="Active"{
					viewModel.CustomerNameArray  = append(viewModel.CustomerNameArray, customers[k.String()].Info.CustomerName)
					viewModel.Keys=append(viewModel.Keys, k.String())
				}

			}
			log.Println("customer name array",viewModel.CustomerNameArray)
			viewModel.PageType=helpers.SelectPageForAdd
			viewModel.CompanyTeamName = storedSession.CompanyTeamName
			viewModel.CompanyPlan = storedSession.CompanyPlan
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		viewModel.AdminFirstName = storedSession.AdminFirstName
		viewModel.AdminLastName = storedSession.AdminLastName
		viewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-job.html"

	}

}
/*Display job details*/
func (c *JobController)LoadJobDetail() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	log.Println("company name",companyTeamName)
	storedSession := ReadSession(w, r, companyTeamName)
	customerId := ""
	customerId = c.Ctx.Input.Param(":customerId")
	job :=models.Job{}
	dbStatus, jobs := job.GetAllJobs(c.AppEngineCtx,companyTeamName)
	viewModel := viewmodels.JobViewModel{}
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(jobs)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			if jobs[k].Settings.Status == "Active" && jobs[k].Customer.CustomerStatus =="Active" {
				var tempValueSlice []string
				if jobs[k].Customer.CustomerStatus =="Active"{
					tempValueSlice = append(tempValueSlice, jobs[k].Customer.CustomerName)
					if !helpers.StringInSlice(jobs[k].Customer.CustomerName, viewModel.UniqueCustomerNames) {
						viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames, jobs[k].Customer.CustomerName)
					}
					if customerId == jobs[k].Customer.CustomerId {
						viewModel.SelectedCustomer = jobs[k].Customer.CustomerName
					}
				}

				tempValueSlice = append(tempValueSlice, jobs[k].Info.JobName)
				tempValueSlice = append(tempValueSlice, jobs[k].Info.JobNumber)
				tempValueSlice = append(tempValueSlice, jobs[k].Info.NumberOfTask)
				tempValueSlice = append(tempValueSlice, jobs[k].Settings.Status)
				viewModel.Values = append(viewModel.Values, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}
		}
		if len(customerId) ==0 {
			viewModel.SelectedCustomer= ""
			viewModel.CustomerMatch="true"

		}
		viewModel.Keys = keySlice
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.CompanyPlan = storedSession.CompanyPlan
		if  len(viewModel.SelectedCustomer) ==0 && len(customerId) !=0{

			viewModel.CustomerMatch ="false"
			viewModel.SelectedCustomer ="false"
		}
		//if  len(viewModel.SelectedCustomer) ==0{
		//	log.Println("dfdgfdgdgd")
		//	viewModel.SelectedCustomer = "No Customer"
		//
		//}
		viewModel.AdminFirstName = storedSession.AdminFirstName
		viewModel.AdminLastName = storedSession.AdminLastName
		viewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/job-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}


}

/*Delete job details from DB*/
//func (c *JobController)LoadDeleteJob() {
//	r := c.Ctx.Request
//	w := c.Ctx.ResponseWriter
//	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
//	ReadSession(w, r, companyTeamName)
//	jobId :=c.Ctx.Input.Param(":jobId")
//	job := models.Job{}
//	dbStatus := job.DeleteJobFromDB(c.AppEngineCtx, jobId)
//	switch dbStatus {
//
//		case true:
//			w.Write([]byte("true"))
//		case false :
//			w.Write([]byte("false"))
//	}
//
//
//}

/*Edit job details*/
func (c *JobController)LoadEditJob() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		jobId := c.Ctx.Input.Param(":jobId")
		job := models.Job{}
		job.Customer.CustomerId = c.GetString("customerId")
		job.Customer.CustomerName = c.GetString("customerName")
		job.Info.JobName = c.GetString("jobName")
		job.Info.JobNumber = c.GetString("jobNumber")
		job.Info.NumberOfTask = c.GetString("numberOfTask")
		job.Settings.DateOfCreation = time.Now().UnixNano() / int64(time.Millisecond)
		job.Settings.Status = helpers.StatusActive
		job.Info.CompanyTeamName = storedSession.CompanyTeamName
		dbStatus := job.UpdateJobToDB(c.AppEngineCtx,jobId)
		switch dbStatus {

			case true:
				w.Write([]byte("true"))

			case false:
				w.Write([]byte("false"))
		}

	} else {
		jobId := c.Ctx.Input.Param(":jobId")
		viewModel := viewmodels.JobViewModel{}
		job := models.Job{}
		dbStatus, jobDetail := job.GetJobDetailById(c.AppEngineCtx, jobId)
		log.Println("job details",jobDetail)
		switch dbStatus {
		case true:
			customers, dbStatus := models.GetAllCustomerDetails(c.AppEngineCtx,companyTeamName)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(customers)
				var keySlice []string
				for _, key := range dataValue.MapKeys() {
					keySlice = append(keySlice, key.String())
				}
				viewModel.Keys = keySlice
				for _, k := range dataValue.MapKeys() {
					viewModel.CustomerNameArray = append(viewModel.CustomerNameArray, customers[k.String()].Info.CustomerName)
				}
				viewModel.PageType = helpers.SelectPageForEdit
				viewModel.CustomerName = jobDetail.Customer.CustomerName
				viewModel.CustomerId =jobDetail.Customer.CustomerId
				viewModel.JobName = jobDetail.Info.JobName
				viewModel.JobNumber = jobDetail.Info.JobNumber
				viewModel.NumberOfTask = jobDetail.Info.NumberOfTask
				viewModel.JobId = jobId
				viewModel.CompanyTeamName = storedSession.CompanyTeamName
				viewModel.CompanyPlan = storedSession.CompanyPlan
				viewModel.AdminFirstName = storedSession.AdminFirstName
				viewModel.AdminLastName = storedSession.AdminLastName
				viewModel.ProfilePicture =storedSession.ProfilePicture
				c.Data["vm"] = viewModel
				c.Layout = "layout/layout.html"
				c.TplName = "template/add-job.html"
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}
//Function to check job name exists in DB
func (c *JobController)CheckJobName(){
	w := c.Ctx.ResponseWriter
	jobName := c.GetString("jobName")
	isJobNameUsed := models.CheckJobNameIsUsed(c.AppEngineCtx, jobName)
	switch isJobNameUsed{
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}
//Function to check job number exists in DB
func (c *JobController)CheckJobNumber(){
	w := c.Ctx.ResponseWriter
	jobNumber := c.GetString("jobNumber")
	isJobNumberUsed := models.CheckJobNumberIsUsed(c.AppEngineCtx, jobNumber)
	switch isJobNumberUsed{
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}
func (c *JobController)LoadDeleteJob() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	log.Println("inside delete")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	jobId := c.Ctx.Input.Param(":jobId")
	user := models.TasksJob{}
	dbStatus, contactDetail := user.IsJobUsedForTask(c.AppEngineCtx, jobId)
	log.Println("status", dbStatus)
	log.Println(contactDetail)
	switch dbStatus {
	case true:
		log.Println("true")
		if len(contactDetail) != 0 {
			dataValue := reflect.ValueOf(contactDetail)
			for _, key := range dataValue.MapKeys() {
				if contactDetail[key.String()].TasksJobStatus == helpers.StatusActive {
					log.Println("insideeee fgjgfjh")
					w.Write([]byte("true"))
					break
				} else {
					log.Println("false")
					w.Write([]byte("false"))
				}
			}
		} else {
			w.Write([]byte("false"))
		}
	case false :
		w.Write([]byte("false"))
	}
}
func (c *JobController) DeleteJobIfNotInTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	jobId := c.Ctx.Input.Param(":jobId")
	user :=models.Job{}
	log.Println("inside deletion of cotact")
	contact :=models.TasksJob{}
	var TaskSlice []string
	dbStatus,jobDetails := contact.IsJobUsedForTask(c.AppEngineCtx, jobId)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(jobDetails)
		for _, key := range dataValue.MapKeys() {
			TaskSlice = append(TaskSlice, key.String())
		}
		dbStatus := user.DeleteJobFromDB(c.AppEngineCtx, jobId,TaskSlice)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false :
			w.Write([]byte("false"))
		}
	}
}



func (c *JobController) RemoveJobFromTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	jobId := c.Ctx.Input.Param(":jobId")
	log.Println("hiiii")
	//contact :=models.TasksContact{}
	//var TaskSlice []string
	//dbStatus,contactDetails := contact.IsContactUsedForTask(c.AppEngineCtx, contactId)
	//switch dbStatus {
	//case true:
	//	dataValue := reflect.ValueOf(contactDetails)
	//	for _, key := range dataValue.MapKeys() {
	//		TaskSlice=append(TaskSlice,key.String())
	//	}
	//
	//	dbStatus := contact.DeleteContactFromTask(c.AppEngineCtx, contactId, TaskSlice)
	//	switch dbStatus {
	//	case true:
	//		w.Write([]byte("true"))
	//	case false:
	//		w.Write([]byte("false"))
	//	}
	//case false:
	//	log.Println("false")
	user :=models.Job{}
	dbStatus := user.DeleteJobFromDBForNonTask(c.AppEngineCtx, jobId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}
}
