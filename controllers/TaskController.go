
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
		ce := appengine.NewContext(r)
		log.Infof(ce, "requested struct: %+v", task)
		log.Infof(ce, "value of login type  ", task.LoginType)

		task.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		task.Status = "Completed"
		task.AddToDB(c.AppEngineCtx)

		}else {
		task:=models.Task{}
		_,tasks :=task.RetrieveFromUserDB(c.AppEngineCtx)
		ce := appengine.NewContext(r)
		log.Infof(ce, "all data", tasks)
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-task.html"


	}



}
func (c *TaskController)LoadTaskDetail() {
	task := models.Task{}
	dbStatus, tasks := task.RetrieveFromDB(c.AppEngineCtx)
	viewModel := viewmodels.TaskViewModel{}

	switch dbStatus {

	case true:
		r := c.Ctx.Request
		ce := appengine.NewContext(r)
		log.Infof(ce, "%s\n", task)
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
	ce := appengine.NewContext(r)
	id :=c.Ctx.Input.Param(":key")
	log.Infof(ce,"idddddddddd",id)
	task := models.Task{}
	dbStatus := task.DeleteFromDB(c.AppEngineCtx, id )

	switch dbStatus {

	case true:
		c.Redirect("/taskdetail", 302)
	case false :
	}


}