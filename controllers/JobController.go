
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
	viewModel := viewmodels.JobViewModel{}
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter
	if r.Method == "POST" {
		job:=models.Job{}
		storedSession := ReadSession(w, r)
		job.Customer.CustomerId = c.GetString("customerId")
		job.Customer.CustomerName = c.GetString("customerName")
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
		storedSession := ReadSession(w, r)
		 customers ,dbStatus:=models.GetAllCustomerDetails(c.AppEngineCtx)
		var keySlice []string
		dataValue := reflect.ValueOf(customers)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		viewModel.Keys=keySlice
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(customers)
			for _, k := range dataValue.MapKeys() {
				viewModel.CustomerNameArray  = append(viewModel.CustomerNameArray, customers[k.String()].Info.CustomerName)
			}
			viewModel.PageType=helpers.SelectPageForAdd
			viewModel.CompanyTeamName = storedSession.CompanyTeamName
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-job.html"

	}

}
/*Display job details*/
func (c *JobController)LoadJobDetail() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	//companyName:=c.GetSession("companyName")
	//log.Println("companyName",companyName)
	customerId := ""
	customerId = c.Ctx.Input.Param(":customerId")
	job :=models.Job{}
	dbStatus, jobs := job.GetAllJobs(c.AppEngineCtx)
	viewModel := viewmodels.JobViewModel{}
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(jobs)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, jobs[k].Customer.CustomerName)
			if !helpers.StringInSlice(jobs[k].Customer.CustomerName, viewModel.UniqueCustomerNames) {
				viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames, jobs[k].Customer.CustomerName)
			}
			if customerId == jobs[k].Customer.CustomerId{
				viewModel.SelectedCustomer = jobs[k].Customer.CustomerName
			}
			tempValueSlice = append(tempValueSlice, jobs[k].Info.JobName)
			tempValueSlice = append(tempValueSlice, jobs[k].Info.JobNumber)
			tempValueSlice = append(tempValueSlice, jobs[k].Info.NumberOfTask)
			tempValueSlice = append(tempValueSlice, jobs[k].Settings.Status)
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}

		viewModel.Keys = keySlice
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/job-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}


}

/*Delete job details from DB*/
func (c *JobController)LoadDeleteJob() {
	jobId :=c.Ctx.Input.Param(":jobId")
	job := models.Job{}
	dbStatus := job.DeleteJobFromDB(c.AppEngineCtx, jobId)
	w := c.Ctx.ResponseWriter
	switch dbStatus {

		case true:
			w.Write([]byte("true"))
		case false :
			w.Write([]byte("false"))
	}


}

/*Edit job details*/
func (c *JobController)LoadEditJob() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		storedSession := ReadSession(w, r)
		jobId := c.Ctx.Input.Param(":jobId")
		job := models.Job{}
		job.Customer.CustomerId = c.GetString("customerId")
		job.Customer.CustomerName = c.GetString("customerName")
		job.Info.JobName = c.GetString("jobName")
		job.Info.JobNumber = c.GetString("jobNumber")
		job.Info.NumberOfTask = c.GetString("numberOfTask")
		job.Settings.DateOfCreation = time.Now().UnixNano() / int64(time.Millisecond)
		job.Settings.Status = "Open"
		job.Info.CompanyTeamName = storedSession.CompanyTeamName
		dbStatus := job.UpdateJobToDB(c.AppEngineCtx,jobId)
		switch dbStatus {

			case true:
				w.Write([]byte("true"))

			case false:
				w.Write([]byte("false"))
		}

	} else {
		storedSession := ReadSession(w, r)
		jobId := c.Ctx.Input.Param(":jobId")
		viewModel := viewmodels.JobViewModel{}
		job := models.Job{}
		dbStatus, jobDetail := job.GetJobDetailById(c.AppEngineCtx, jobId)
		switch dbStatus {
		case true:
			customers, dbStatus := models.GetAllCustomerDetails(c.AppEngineCtx)
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
				viewModel.JobName = jobDetail.Info.JobName
				viewModel.JobNumber = jobDetail.Info.JobNumber
				viewModel.NumberOfTask = jobDetail.Info.NumberOfTask
				viewModel.JobId = jobId
				viewModel.CompanyTeamName = storedSession.CompanyTeamName
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