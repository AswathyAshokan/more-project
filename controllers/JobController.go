
/* Author :Aswathy Ashok */

package controllers

import (

	//"github.com/astaxie/beegae"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"time"

	"reflect"

	"app/passporte/helpers"
)

type JobController struct {
	BaseController
}

func (c *JobController)LoadJob() {
	viewModel := viewmodels.JobViewModel{}
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter
	if r.Method == "POST" {

		job:=models.Job{}
		job.CustomerName= c.GetString("customerName")
		job.JobName= c.GetString("jobName")
		job.JobNumber = c.GetString("jobNumber")
		job.NumberOfTask = c.GetString("numberOfTask")
		job.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		job.Status = "Open"
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
		switch dbStatus {

		case true:

			dataValue := reflect.ValueOf(jobs)
			var keySlice []string

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			customerValue := job.RetrieveCustomerNameFromDB(c.AppEngineCtx,keySlice)

			viewModel.CustomerNameArray  =customerValue
			viewModel.PageType="1"
		case false:

		}
		c.Data["array"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-job.html"

	}



}
func (c *JobController)LoadJobDetail() {
	job := models.Job{}
	dbStatus, jobs := job.RetrieveJobFromDB(c.AppEngineCtx)
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
	}


}
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
func (c *JobController)LoadEditJob() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		jobId := c.Ctx.Input.Param(":jobId")
		job := models.Job{}
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

		dbStatus, jobDetail := job.RetrieveJobDetailFromDB(c.AppEngineCtx, jobId)
		switch dbStatus {
		case true:
			_, jobs := job.RetrieveCustomerFromDB(c.AppEngineCtx)
			dataValue := reflect.ValueOf(jobs)
			var keySlice []string

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			customerValue := job.RetrieveCustomerNameFromDB(c.AppEngineCtx, keySlice)
			viewModel.CustomerNameArray = customerValue
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

		}

	}
}