
/* Author :Aswathy Ashok */

package controllers

import (

	//"github.com/astaxie/beegae"
	"app/passporte/models"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"time"
	"app/passporte/viewmodels"
	"reflect"
)

type TaskController struct {
	BaseController
}

func (c *TaskController)LoadTask() {
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter
	viewModel  := viewmodels.TaskViewModel{}
	if r.Method == "POST" {

		task:=models.Task{}
		task.TaskName= c.GetString("taskName")
		task.JobName= c.GetString("jobName")
		task.StartDate = c.GetString("startDate")
		task.EndDate = c.GetString("endDate")
		task.TaskLocation = c.GetString("taskLocation")
		task.TaskDescription = c.GetString("taskDescription")
		task.UserNumber = c.GetString("users")
		task.Log = c.GetString("log")
		task.UserType = c.GetString("userType")
		task.Contact = c.GetString("contacts")
		task.LoginType=c.GetString("loginType")
		task.FitToWork = c.GetString("fitToWork")
		task.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		task.Status = "Completed"
		dbStatus :=task.AddTaskToDB(c.AppEngineCtx)
		switch dbStatus {

			case true:
				w.Write([]byte("true"))

			case false:
				w.Write([]byte("false"))
		}

	}else {
		task:=models.Task{}
		job :=models.Job{}

		dbStatus,tasks :=task.RetrieveJobFromDB(c.AppEngineCtx)
		switch dbStatus {

		case true:

			dataValue := reflect.ValueOf(tasks)
			var keySlice []string

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			jobValue := job.RetrieveJobValueFromDB(c.AppEngineCtx,keySlice)

			viewModel.JobNameArray  =jobValue

		case false:

		}


		contact :=models.ContactUser{}
		dbStatus,contacts :=contact.RetrieveContactFromDB(c.AppEngineCtx)
		switch dbStatus {

		case true:

			dataValue := reflect.ValueOf(contacts)
			var keySlice []string

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			contactsName := contact.RetrieveContactNameFromDB(c.AppEngineCtx, keySlice)

			viewModel.ContactNameArray = contactsName
			viewModel.Key=keySlice
		case false:
		}
		c.Data["array"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-task.html"


	}

}
func (c *TaskController)LoadTaskDetail() {
	task := models.Task{}
	dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx)
	viewModel := viewmodels.TaskViewModel{}

	switch dbStatus {

	case true:

		//var valueSlice []models.User
		dataValue := reflect.ValueOf(tasks)
		var keySlice []string
		var valueSlice []models.Task
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		// To perform the opertion you want
		for _, k := range keySlice {
			valueSlice = append(valueSlice, tasks[k])
			viewModel.Task = append(viewModel.Task, tasks[k])
			viewModel.Key=keySlice

		}

		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/task-details.html"

	case false:

	}


}
func (c *TaskController)LoadDeleteTask() {

	r := c.Ctx.Request
	context := appengine.NewContext(r)
	taskId :=c.Ctx.Input.Param(":taskId")
	log.Infof(context,"idddddddddd", taskId)
	task := models.Task{}
	dbStatus := task.DeleteTaskFromDB(c.AppEngineCtx, taskId)

	switch dbStatus {

	case true:
		c.Redirect("/task", 302)
	case false :
	}


}
func (c *TaskController)LoadEditTask() {
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter
	context := appengine.NewContext(r)

	viewModel  := viewmodels.TaskViewModel{}
	if r.Method == "POST" {
		taskId := c.Ctx.Input.Param(":taskId")
		task:=models.Task{}
		task.TaskName= c.GetString("taskName")
		task.JobName= c.GetString("jobName")
		task.StartDate = c.GetString("startDate")
		task.EndDate = c.GetString("endDate")
		task.TaskLocation = c.GetString("taskLocation")
		task.TaskDescription = c.GetString("taskDescription")
		task.UserNumber = c.GetString("users")
		task.Log = c.GetString("log")
		task.UserType = c.GetString("userType")
		task.Contact = c.GetString("contacts")
		task.LoginType=c.GetString("loginType")
		task.FitToWork = c.GetString("fitToWork")
		task.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		task.Status = "Completed"
		dbStatus :=task.UpdateTaskToDB(c.AppEngineCtx,taskId)
		switch dbStatus {

		case true:
			w.Write([]byte("true"))

		case false:
			w.Write([]byte("false"))
		}


	} else {
		log.Infof(context, "insideee edittt")
		task:=models.Task{}
		job :=models.Job{}
		taskId := c.Ctx.Input.Param(":taskId")
		dbStatus, taskDetail := task.RetrieveTaskDetailFromDB(c.AppEngineCtx, taskId)
		log.Infof(context, "our task details", taskDetail)
		switch dbStatus {

		case true:
			_,tasks :=task.RetrieveJobFromDB(c.AppEngineCtx)
			dataValue := reflect.ValueOf(tasks)
			var keySlice []string

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			jobValue := job.RetrieveJobValueFromDB(c.AppEngineCtx,keySlice)

			viewModel.JobNameArray  =jobValue
			contact :=models.ContactUser{}
			_,contacts :=contact.RetrieveContactFromDB(c.AppEngineCtx)
			dataValue = reflect.ValueOf(contacts)

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			contactsName := contact.RetrieveContactNameFromDB(c.AppEngineCtx, keySlice)
			viewModel.ContactNameArray = contactsName
			viewModel.Key=keySlice
			viewModel.PageType = "2"
			viewModel.JobName = taskDetail.JobName
			viewModel.TaskName = taskDetail.TaskName
			viewModel.TaskLocation = taskDetail.TaskLocation
			viewModel.StartDate = taskDetail.StartDate
			viewModel.EndDate = taskDetail.EndDate
			viewModel.TaskDescription= taskDetail.TaskDescription
			viewModel.UserNumber = taskDetail.UserNumber
			viewModel.Log = taskDetail.Log
			viewModel.UserType = taskDetail.UserType
			viewModel.Contact = taskDetail.Contact
			viewModel.FitToWork = taskDetail.FitToWork
			c.Data["array"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-task.html"

		case false:

		}



	}

}

