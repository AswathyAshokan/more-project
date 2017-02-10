
/* Author :Aswathy Ashok */

package controllers

import (

	//"github.com/astaxie/beegae"
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"
	"reflect"
	"app/passporte/helpers"
	"log"
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
		user :=models.User{}

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



		dbStatus, taskUserValue :=user.RetrieveUserFromDB(c.AppEngineCtx)
		log.Println("user database value",taskUserValue)
		switch dbStatus {

		case true:

			dataValue := reflect.ValueOf(taskUserValue)
			var keySlice []string

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			userValue := user.RetrieveUserNameFromDB(c.AppEngineCtx, keySlice)
			log.Println("user name",userValue)

			viewModel.GroupNameArray = userValue
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
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		// To perform the opertion you want
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, tasks[k].JobName)
			tempValueSlice = append(tempValueSlice, tasks[k].TaskName)
			tempValueSlice = append(tempValueSlice, tasks[k].TaskLocation)
			tempValueSlice = append(tempValueSlice, tasks[k].StartDate)
			tempValueSlice = append(tempValueSlice, tasks[k].EndDate)
			tempValueSlice = append(tempValueSlice,  tasks[k].LoginType)
			tempValueSlice = append(tempValueSlice,  tasks[k].Status)
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		viewModel.Keys = keySlice
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/task-details.html"

	case false:

	}


}
func (c *TaskController)LoadDeleteTask() {


	w :=c.Ctx.ResponseWriter
	taskId :=c.Ctx.Input.Param(":taskId")

	task := models.Task{}
	dbStatus := task.DeleteTaskFromDB(c.AppEngineCtx, taskId)

	switch dbStatus {

	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("true"))
	}


}
func (c *TaskController)LoadEditTask() {
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter

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

		task:=models.Task{}
		job :=models.Job{}
		taskId := c.Ctx.Input.Param(":taskId")
		dbStatus, taskDetail := task.RetrieveTaskDetailFromDB(c.AppEngineCtx, taskId)
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
			viewModel.PageType = helpers.SelectPageForEdit
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
			viewModel.TaskId=taskId
			c.Data["array"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-task.html"

		case false:

		}



	}

}

