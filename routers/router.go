/*Author: Sarath
Date:01/02/2017*/
package routers

import (
	"app/passporte/controllers"
	"github.com/astaxie/beegae"
)

func init() {

	//Sarath
	beegae.Router("/", &controllers.LoginController{}, "*:Login")
	beegae.Router("/register", &controllers.RegisterController{}, "*:Register")
	beegae.Router("/nfc", &controllers.NfcController{},"*:NFCDetails")
	beegae.Router("/nfc/add",&controllers.NfcController{},"*:AddNFC")
	beegae.Router("/datatable",&controllers.NfcController{},"*:Datatable")
	beegae.Router("/nfc/delete",&controllers.NfcController{},"*:DeleteNFC")

	//Farsana
	beegae.Router("/add-customer", &controllers.CustomerController{}, "*:AddCustomer")
	beegae.Router("/customer", &controllers.CustomerController{}, "*:CustomerDetails")
	beegae.Router("/add-group", &controllers.GroupController{}, "*:AddGroup")
	beegae.Router("/group", &controllers.GroupController{}, "*:GroupDetails")
	beegae.Router("/delete-group/:Key", &controllers.GroupController{}, "*:GroupDelete")
	beegae.Router("/add-user", &controllers.InviteUserController{}, "*:AddUser")
	beegae.Router("/invite", &controllers.InviteUserController{}, "*:UserDetails")
	beegae.Router("/delete-user/:Key", &controllers.InviteUserController{}, "*:UserDelete")
	beegae.Router("/edit-user/:Key", &controllers.InviteUserController{}, "*:UserEdit")
	beegae.Router("/view-user/:Key", &controllers.InviteUserController{}, "*:UserView")
	beegae.Router("/delete-customer/:Key", &controllers.CustomerController{}, "*:CustomerDelete")
	beegae.Router("/edit-customer/:Key", &controllers.InviteUserController{}, "*:UserEdit")
	beegae.Router("/view-customer/:Key", &controllers.InviteUserController{}, "*:UserView")

	//Aswathy
	beegae.Router("/contact/add", &controllers.ContactUserController{},"*:LoadContact")
	beegae.Router("/contact", &controllers.ContactUserController{},"*:LoadContactdetail")
	beegae.Router("/task/add", &controllers.TaskController{},"*:LoadTask")
	beegae.Router("/task", &controllers.TaskController{},"*:LoadTaskDetail")
	beegae.Router("/project/add", &controllers.ProjectController{},"*:LoadProject")
	beegae.Router("/project", &controllers.ProjectController{},"*:LoadProjectDetail")
	beegae.Router("/contact/:contactId/edit", &controllers.ContactUserController{},"*:LoadEditContact")
	beegae.Router("/contact/:contactId/delete", &controllers.ContactUserController{},"*:LoadDeleteContact")
	//beegae.Router("/editProject", &controllers.ProjectController{},"*:LoadEditProject")
	beegae.Router("/project/:projectId/delete", &controllers.ProjectController{},"*:LoadDeleteProject")
	beegae.Router("/task/:taskId/delete", &controllers.TaskController{},"*:LoadDeleteTask")
}
