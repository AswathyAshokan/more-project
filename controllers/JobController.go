
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
		job.CustomerId = c.GetString("customerId")
		job.CustomerName = c.GetString("customerName")
		job.JobName = c.GetString("jobName")
		job.JobNumber = c.GetString("jobNumber")
		job.NumberOfTask = c.GetString("numberOfTask")
		job.CurrentDate = time.Now().UnixNano() / int64(time.Millisecond)
		job.Status = helpers.StatusActive
		dbStatus :=job.AddJobToDB(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}else {
		job :=models.Job{}
		dbStatus,jobs :=job.RetrieveCustomerFromDB(c.AppEngineCtx)
		var keySlice []string
		dataValue := reflect.ValueOf(jobs)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		viewModel.Keys=keySlice
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(jobs)
			for _, k := range dataValue.MapKeys() {
				viewModel.CustomerNameArray  = append(viewModel.CustomerNameArray, jobs[k.String()].CustomerName)
				}
			viewModel.PageType="add"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		c.Data["array"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-job.html"

	}

}
/*Display job details*/
func (c *JobController)LoadJobDetail() {
	customerId := ""
	customerId = c.Ctx.Input.Param(":customerId")
	job := models.Job{}
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
			tempValueSlice = append(tempValueSlice, jobs[k].CustomerName)
			if !helpers.StringInSlice(jobs[k].CustomerName, viewModel.UniqueCustomerNames) {
				viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames, jobs[k].CustomerName)
			}
			if customerId == jobs[k].CustomerId{
				viewModel.SelectedCustomer = jobs[k].CustomerName
			}
			tempValueSlice = append(tempValueSlice, jobs[k].JobName)
			tempValueSlice = append(tempValueSlice, jobs[k].JobNumber)
			tempValueSlice = append(tempValueSlice, jobs[k].NumberOfTask)
			tempValueSlice = append(tempValueSlice, jobs[k].Status)
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}

		viewModel.Keys = keySlice
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
		jobId := c.Ctx.Input.Param(":jobId")
		job := models.Job{}
		job.CustomerId = c.GetString("customerId")
		job.CustomerName = c.GetString("customerName")
		job.JobName = c.GetString("jobName")
		job.JobNumber = c.GetString("jobNumber")
		job.NumberOfTask = c.GetString("numberOfTask")
		job.CurrentDate = time.Now().UnixNano() / int64(time.Millisecond)
		job.Status = "Open"
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
		switch dbStatus {
		case true:
			_, jobs := job.RetrieveCustomerFromDB(c.AppEngineCtx)
			dataValue := reflect.ValueOf(jobs)
			for _, k := range dataValue.MapKeys() {
				viewModel.CustomerNameArray  = append(viewModel.CustomerNameArray, jobs[k.String()].CustomerName)
			}
			viewModel.PageType = helpers.SelectPageForEdit
			viewModel.CustomerName = jobDetail.CustomerName
			viewModel.JobName = jobDetail.JobName
			viewModel.JobNumber = jobDetail.JobNumber
			viewModel.NumberOfTask = jobDetail.NumberOfTask
			viewModel.JobId= jobId
			c.Data["array"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-job.html"
		case false:
			log.Println(helpers.ServerConnectionError)

		}

	}
}