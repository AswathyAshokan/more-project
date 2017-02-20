
/* Author :Aswathy Ashok */

package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"
	"reflect"
	"app/passporte/helpers"
	"log"
	"bytes"
)

type TaskController struct {
	BaseController
}
/*Add task details to DB*/
func (c *TaskController)AddNewTask() {
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter

	if r.Method == "POST" {
		task:=models.Task{}
		task.TaskName= c.GetString("taskName")
		task.JobName= c.GetString("jobName")
		task.CustomerName = c.GetString("customerName")
		task.StartDate = c.GetString("startDate")
		task.EndDate = c.GetString("endDate")
		task.TaskLocation = c.GetString("taskLocation")
		task.TaskDescription = c.GetString("taskDescription")
		task.UserNumber = c.GetString("users")
		task.Log = c.GetString("log")
		task.UsersOrGroups = c.GetStrings("userOrGroup")
		tempContactId := c.GetStrings("contacts")

		for i := 0; i < len(tempContactId); i++ {
			task.ContactId = append(task.ContactId, tempContactId[i])

		}
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
		viewModel  := viewmodels.AddTaskViewModel{}
		user :=models.User{}
		var keySlice []string
		//Getting Jobs
		dbStatus, allJobs := models.GetAllJobs(c.AppEngineCtx)
		switch dbStatus {
		case true:

			log.Println("1")
			dataValue := reflect.ValueOf(allJobs)
			log.Println(dataValue)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for _, k := range dataValue.MapKeys() {
				viewModel.JobNameArray   = append(viewModel.JobNameArray, allJobs[k.String()].JobName)
				log.Println(viewModel.JobNameArray)
				viewModel.JobCustomerNameArray = append(viewModel.JobCustomerNameArray, allJobs[k.String()].CustomerName)
				log.Println(viewModel.JobCustomerNameArray)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		//Getting users and groups
		dbStatus, allUsers := user.GetAllUsers(c.AppEngineCtx)
		switch dbStatus {
		case true:

			log.Println("2")
			dataValue := reflect.ValueOf(allUsers)
			for _, k := range dataValue.MapKeys() {
				viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,  allUsers[k.String()].FirstName+""+ allUsers[k.String()].LastName+"(User)")
			}
			allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx)
			switch dbStatus {
			case true:
				log.Println("3")
				dataValue = reflect.ValueOf(allGroups)

				for _, k := range dataValue.MapKeys() {
					viewModel.GroupNameArray = append(viewModel.GroupNameArray, allGroups[k.String()].GroupName+"(Group)")
				}
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		dbStatus, contacts := models.GetAllContact(c.AppEngineCtx)
		switch dbStatus {
		case true:
			log.Println("4")
			dataValue := reflect.ValueOf(contacts)
			for _, k := range dataValue.MapKeys() {
				viewModel.ContactNameArray  = append(viewModel.ContactNameArray , contacts[k.String()].Name)
			}

			log.Println("5: ", keySlice)
			viewModel.Key = keySlice
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		log.Println("Data: ", viewModel)
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-task.html"


	}

}

/* display all task details*/
func (c *TaskController)LoadTaskDetail() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	task := models.Task{}
	dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx)
	viewModel := viewmodels.TaskDetailViewModel{}

	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(tasks)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			var buffer bytes.Buffer
			tempJobAndCustomer := ""
			buffer.WriteString(tasks[k].JobName)
			buffer.WriteString(" (")
			buffer.WriteString(tasks[k].CustomerName)
			buffer.WriteString(")")
			tempJobAndCustomer = buffer.String()
			buffer.Reset()
			tempValueSlice = append(tempValueSlice, tempJobAndCustomer)

			if !helpers.StringInSlice(tasks[k].CustomerName, viewModel.UniqueCustomerNames) {
				viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames, tasks[k].CustomerName)
			}
			/*if customerId == jobs[k].CustomerId{
				viewModel.SelectedCustomer = jobs[k].CustomerName
			}*/

			if !helpers.StringInSlice(tasks[k].JobName, viewModel.UniqueJobNames) {
				viewModel.UniqueJobNames = append(viewModel.UniqueJobNames, tasks[k].JobName)
			}

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
			log.Println(helpers.ServerConnectionError)

	}


}
/*delete task details from DB*/
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

/* Edit task details*/
func (c *TaskController)LoadEditTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		taskId := c.Ctx.Input.Param(":taskId")
		task := models.Task{}
		task.TaskName = c.GetString("taskName")
		task.JobName = c.GetString("jobName")
		task.CustomerName = c.GetString("customerName")
		task.StartDate = c.GetString("startDate")
		task.EndDate = c.GetString("endDate")
		task.TaskLocation = c.GetString("taskLocation")
		task.TaskDescription = c.GetString("taskDescription")
		task.UserNumber = c.GetString("users")
		task.Log = c.GetString("log")
		task.UsersOrGroups = c.GetStrings("userOrGroup")
		tempContactId := c.GetStrings("contacts")
		for i := 0; i < len(tempContactId); i++ {
			task.ContactId = append(task.ContactId, tempContactId[i])
		}
		task.LoginType = c.GetString("loginType")
		task.FitToWork = c.GetString("fitToWork")
		task.CurrentDate = time.Now().UnixNano() / int64(time.Millisecond)
		task.Status = "Completed"
		dbStatus := task.UpdateTaskToDB(c.AppEngineCtx, taskId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		viewModel  := viewmodels.EditTaskViewModel{}
		task := models.Task{}
		user := models.User{}
		taskId := c.Ctx.Input.Param(":taskId")
		dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, taskId)
		switch dbStatus {
		case true:
			dbStatus, allJobs := models.GetAllJobs(c.AppEngineCtx)
			var keySlice []string
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allJobs)
				for _, key := range dataValue.MapKeys() {
					keySlice = append(keySlice, key.String())
				}
				for _, k := range dataValue.MapKeys() {
					viewModel.JobNameArray = append(viewModel.JobNameArray, allJobs[k.String()].JobName)
					viewModel.JobCustomerNameArray = append(viewModel.JobCustomerNameArray, allJobs[k.String()].CustomerName)
				}
			case false:
				log.Println(helpers.ServerConnectionError)
			}
			dbStatus, taskUserValue := user.GetAllUsers(c.AppEngineCtx)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(taskUserValue)
				for _, k := range dataValue.MapKeys() {
					viewModel.GroupNameArray = append(viewModel.GroupNameArray, taskUserValue[k.String()].FirstName + "" + taskUserValue[k.String()].LastName+"(User)")
				}
				allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx)
				switch dbStatus {
				case true:
					dataValue = reflect.ValueOf(allGroups)
					for _, k := range dataValue.MapKeys() {
						viewModel.GroupNameArray = append(viewModel.GroupNameArray, allGroups[k.String()].GroupName+"(Group)")
					}
				case false:
					log.Println(helpers.ServerConnectionError)
				}
			case false:
				log.Println(helpers.ServerConnectionError)
			}
			dbStatus, contacts := models.GetAllContact(c.AppEngineCtx)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(contacts)
				for _, k := range dataValue.MapKeys() {
					viewModel.ContactNameArray = append(viewModel.ContactNameArray, contacts[k.String()].Name)
				}
				viewModel.Key = keySlice
			case false:
				log.Println(helpers.ServerConnectionError)
			}
			viewModel.Key = keySlice
			viewModel.PageType = helpers.SelectPageForEdit
			viewModel.JobName = taskDetail.JobName
			viewModel.TaskName = taskDetail.TaskName
			viewModel.TaskLocation = taskDetail.TaskLocation
			viewModel.StartDate = taskDetail.StartDate
			viewModel.EndDate = taskDetail.EndDate
			viewModel.TaskDescription = taskDetail.TaskDescription
			viewModel.UserNumber = taskDetail.UserNumber
			viewModel.Log = taskDetail.Log
			viewModel.UserType = taskDetail.UsersOrGroups
			viewModel.FitToWork = taskDetail.FitToWork
			viewModel.TaskId = taskId
			c.Data["vm"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-task.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}


