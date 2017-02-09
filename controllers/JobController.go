
/* Author :Aswathy Ashok */

package controllers

import (

	//"github.com/astaxie/beegae"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"time"

	"reflect"

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
		context := appengine.NewContext(r)
		log.Infof(context, "requested struct: %+v", job)
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
		r := c.Ctx.Request
		context := appengine.NewContext(r)
		log.Infof(context, "%s\n", jobs)
		//var valueSlice []models.User
		dataValue := reflect.ValueOf(jobs)
		var keySlice []string
		var valueSlice []models.Job
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		for _, k := range keySlice {
			valueSlice = append(valueSlice, jobs[k])
			viewModel.Job = append(viewModel.Job, jobs[k])
			viewModel.Key=keySlice

		}
		log.Infof(context,"Key:", keySlice, "Value:", valueSlice)
		log.Infof(context, "typeeee",(reflect.TypeOf(viewModel)))
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/job-details.html"

	case false:
	}


}
func (c *JobController)LoadDeleteJob() {

	r := c.Ctx.Request
	context := appengine.NewContext(r)
	jobId :=c.Ctx.Input.Param(":jobId")
	log.Infof(context,"idddddddddd", jobId)
	job := models.Job{}
	dbStatus := job.DeleteJobFromDB(c.AppEngineCtx, jobId)

	switch dbStatus {

	case true:
		c.Redirect("/job", 302)
	case false :
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
		context := appengine.NewContext(r)
		log.Infof(context, "requested struct: %+v", job)
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

		context := appengine.NewContext(r)
		jobId := c.Ctx.Input.Param(":jobId")
		log.Infof(context, "idddddddddd", jobId)
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
			viewModel.PageType = "2"
			viewModel.CustomerName = jobDetail.CustomerName
			viewModel.JobName = jobDetail.JobName
			viewModel.JobNumber = jobDetail.JobNumber
			viewModel.NumberOfTask = jobDetail.NumberOfTask
			c.Data["array"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-job.html"
		case false:

		}

	}
}