
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
		task.Info.TaskName= c.GetString("taskName")
		task.Job.JobName= c.GetString("jobName")
		task.Job.JobId = c.GetString("jobId")
		task.Customer.CustomerName = c.GetString("customerName")
		log.Println("customer",task.Customer.CustomerName,task.Job.JobId)
		task.Customer.CustomerId =c.GetString("jobId")
		task.Info.StartDate = c.GetString("startDate")
		task.Info.EndDate = c.GetString("endDate")
		task.Info.TaskLocation = c.GetString("taskLocation")
		task.Info.TaskDescription = c.GetString("taskDescription")
		task.Info.UserNumber = c.GetString("users")
		task.Info.Log = c.GetString("log")
		//task.UserOrGroup= c.GetStrings("userOrGroup")
		//tempContactId := c.GetStrings("contacts")

		//for i := 0; i < len(tempContactId); i++ {
		//	task.ContactId = append(task.ContactId, tempContactId[i])
		//
		//}
		contacts := models.TaskContact{}
		tempContactName := c.GetStrings("contactName")
		tempContactId := c.GetStrings("contactId")

		for i := 0; i < len(tempContactId); i++ {
			contacts.ContactId= tempContactId[i]
			contacts.ContactName = tempContactName[i]
			task.Contact = append(task.Contact ,contacts)
		}
		task.Info.LoginType=c.GetString("loginType")
		task.Info.FitToWork = c.GetString("fitToWork")
		task.Settings.DateOfCreation =time.Now().UnixNano() / int64(time.Millisecond)
		task.Settings.Status = "Completed"
		dbStatus :=task.AddTaskToDB(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}else {
		viewModel  := viewmodels.AddTaskViewModel{}
		user :=models.Users{}
		var keySlice []string
		var keySliceForGroupAndUser 	[]string
		var keySliceForContact		[]string
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
				viewModel.JobNameArray   = append(viewModel.JobNameArray, allJobs[k.String()].Info.JobName)
				log.Println(viewModel.JobNameArray)
				viewModel.JobCustomerNameArray = append(viewModel.JobCustomerNameArray, allJobs[k.String()].Customer.CustomerName)
				log.Println(viewModel.JobCustomerNameArray)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		//Getting users and groups
		 allUsers,dbStatus := user.GetUsersForDropdown(c.AppEngineCtx)
		log.Println("users for task",allUsers)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(allUsers)
			log.Println(dataValue)
			for _, key := range dataValue.MapKeys() {
				keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
			}
			for _, k := range dataValue.MapKeys() {
				viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,  allUsers[k.String()].Info.FullName+"(User)")
			}
			allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allGroups)
				for _, key := range dataValue.MapKeys() {
					keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
				}
				dataValue = reflect.ValueOf(allGroups)
				for _, k := range dataValue.MapKeys() {
					viewModel.GroupNameArray = append(viewModel.GroupNameArray, allGroups[k.String()].Info.GroupName+"(Group)")
				}
				viewModel.UserAndGroupKey=keySliceForGroupAndUser
				log.Println("user and group key",keySliceForGroupAndUser)
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
			for _, key := range dataValue.MapKeys() {
				keySliceForContact = append(keySliceForContact, key.String())
			}
			for _, k := range dataValue.MapKeys() {
				viewModel.ContactNameArray  = append(viewModel.ContactNameArray , contacts[k.String()].Info.Name)
			}

			log.Println("5: ", keySlice)
			viewModel.Key = keySlice
			viewModel.ContactKey=keySliceForContact
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
	jobId := ""
	jobId = c.Ctx.Input.Param(":jobId")
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
			buffer.WriteString(tasks[k].Job.JobName)
			buffer.WriteString(" (")
			buffer.WriteString(tasks[k].Customer.CustomerName)
			buffer.WriteString(")")
			tempJobAndCustomer = buffer.String()
			buffer.Reset()
			tempValueSlice = append(tempValueSlice, tempJobAndCustomer)

			if !helpers.StringInSlice(tasks[k].Customer.CustomerName, viewModel.UniqueCustomerNames) {
				viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames, tasks[k].Customer.CustomerName)
			}
			if jobId == tasks[k].Job.JobId{
				viewModel.SelectedJob = tasks[k].Job.JobName
				viewModel.SelectedCustomerForJob=tasks[k].Customer.CustomerName
			}
			if !helpers.StringInSlice(tasks[k].Job.JobName, viewModel.UniqueJobNames) {
				viewModel.UniqueJobNames = append(viewModel.UniqueJobNames, tasks[k].Job.JobName)
			}
			tempValueSlice = append(tempValueSlice, tasks[k].Info.TaskName)
			tempValueSlice = append(tempValueSlice, tasks[k].Info.TaskLocation)
			tempValueSlice = append(tempValueSlice, tasks[k].Info.StartDate)
			tempValueSlice = append(tempValueSlice, tasks[k].Info.EndDate)
			tempValueSlice = append(tempValueSlice,  tasks[k].Info.LoginType)
			tempValueSlice = append(tempValueSlice,  tasks[k].Settings.Status)
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
		task.Info.TaskName = c.GetString("taskName")
		task.Job.JobName = c.GetString("jobName")
		task.Job.JobId = c.GetString("jobId")
		task.Customer.CustomerName = c.GetString("customerName")
		task.Customer.CustomerId =c.GetString("jobId")
		task.Info.StartDate = c.GetString("startDate")
		task.Info.EndDate = c.GetString("endDate")
		task.Info.TaskLocation = c.GetString("taskLocation")
		task.Info.TaskDescription = c.GetString("taskDescription")
		task.Info.UserNumber = c.GetString("users")
		task.Info.Log = c.GetString("log")
		//task.UsersOrGroups = c.GetStrings("userOrGroup")
		//tempContactId := c.GetStrings("contacts")
		//for i := 0; i < len(tempContactId); i++ {
		//	task.ContactId = append(task.ContactId, tempContactId[i])
		//}
		contacts := models.TaskContact{}
		tempContactName := c.GetStrings("contacts")
		tempContactId := c.GetStrings("contactId")

		for i := 0; i < len(tempContactId); i++ {
			contacts.ContactId= tempContactId[i]
			contacts.ContactName = tempContactName[i]
			task.Contact = append(task.Contact ,contacts)
		}
		task.Info.LoginType = c.GetString("loginType")
		task.Info.FitToWork = c.GetString("fitToWork")
		task.Settings.DateOfCreation = time.Now().UnixNano() / int64(time.Millisecond)
		task.Settings.Status = "Completed"
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
		user := models.Users{}
		taskId := c.Ctx.Input.Param(":taskId")
		dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, taskId)
		switch dbStatus {
		case true:
			dbStatus, allJobs := models.GetAllJobs(c.AppEngineCtx)
			var keySlice []string
			var keySliceForGroupAndUser 	[]string
			var keySliceForContact		[]string
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allJobs)
				for _, key := range dataValue.MapKeys() {
					keySlice = append(keySlice, key.String())
				}
				for _, k := range dataValue.MapKeys() {
					viewModel.JobNameArray = append(viewModel.JobNameArray, allJobs[k.String()].Info.JobName)
					viewModel.JobCustomerNameArray = append(viewModel.JobCustomerNameArray, allJobs[k.String()].Customer.CustomerName)
				}
			case false:
				log.Println(helpers.ServerConnectionError)
			}
			allUsers,dbStatus := user.GetUsersForDropdown(c.AppEngineCtx)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allUsers)
				for _, key := range dataValue.MapKeys() {
					keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
				}

				for _, k := range dataValue.MapKeys() {
					viewModel.GroupNameArray = append(viewModel.GroupNameArray, allUsers[k.String()].Info.FullName+"(User)")
				}
				allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx)
				switch dbStatus {
				case true:
					dataValue := reflect.ValueOf(allGroups)
					for _, key := range dataValue.MapKeys() {
						keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
					}
					for _, k := range dataValue.MapKeys() {
						viewModel.GroupNameArray = append(viewModel.GroupNameArray, allGroups[k.String()].Info.GroupName+"(Group)")
					}
					viewModel.UserAndGroupKey=keySliceForGroupAndUser
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
				for _, key := range dataValue.MapKeys() {
					keySliceForContact = append(keySliceForContact, key.String())
				}
				for _, k := range dataValue.MapKeys() {
					viewModel.ContactNameArray = append(viewModel.ContactNameArray, contacts[k.String()].Info.Name)
				}
				viewModel.ContactKey=keySliceForContact
				contactDetails, dbStatus := task.GetContactDetailById(c.AppEngineCtx, taskId)
				switch dbStatus {
				case true:
					for i := 0; i < len(contactDetails.Contact); i++ {
						viewModel.ContactNameToEdit = append(viewModel.ContactNameToEdit, contactDetails.Contact[i].ContactId)
					}
					viewModel.Key = keySlice
					viewModel.PageType = helpers.SelectPageForEdit
					viewModel.JobName = taskDetail.Job.JobName
					viewModel.TaskName = taskDetail.Info.TaskName
					viewModel.TaskLocation = taskDetail.Info.TaskLocation
					viewModel.StartDate = taskDetail.Info.StartDate
					viewModel.EndDate = taskDetail.Info.EndDate
					viewModel.TaskDescription = taskDetail.Info.TaskDescription
					viewModel.UserNumber = taskDetail.Info.UserNumber
					viewModel.Log = taskDetail.Info.Log
					//viewModel.UserType = taskDetail.UsersOrGroups
					viewModel.FitToWork = taskDetail.Info.FitToWork
					viewModel.TaskId = taskId
					c.Data["vm"] = viewModel
					c.Layout = "layout/layout.html"
					c.TplName = "template/add-task.html"

				case false:
					log.Println(helpers.ServerConnectionError)
				}
			}

		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}


