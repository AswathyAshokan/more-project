
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
		task.ProjectName= c.GetString("projectName")
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
		context := appengine.NewContext(r)
		log.Infof(context, "requested struct: %+v", task)
		log.Infof(context, "value of login type  ", task.LoginType)

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
		project :=models.Project{}
		_,tasks :=task.RetrieveProjectFromDB(c.AppEngineCtx)
		context := appengine.NewContext(r)

		dataValue := reflect.ValueOf(tasks)
		var keySlice []string

		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}


		tasksValue := project.RetrieveProjectValueFromDB(c.AppEngineCtx,keySlice)
		log.Infof(context, "all data",tasksValue)
		viewModel.ProjectName  =tasksValue
		c.Data["vm"] = viewModel
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
		r := c.Ctx.Request
		context := appengine.NewContext(r)
		log.Infof(context, "%s\n", task)
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
		//log.Infof(ce,"Key:", keySlice, "Value:", valueSlice)
		//log.Infof(ce,"Value: ", valueSlice)
		//log.Infof(ce,"Value: ", valueSlice)
		//mvVar := map["Name"].(string)
		//m := f.(map[string]interface{}
		//viewModel.Name = contact[result[i]].Name
		//viewModel.PhoneNumber = contact["PhoneNumber"]
		//viewModel.Email = contact["Email"]
		//viewModel.Address = contact["Address"]
		//viewModel.State = contact["State"]
		//viewModel.ZipCode = contact["ZipCode"]
		//log.Infof(ce, "typeeee",(reflect.TypeOf(viewModel)))
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
