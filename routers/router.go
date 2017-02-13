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
	//beegae.Router("/datatable",&controllers.NfcController{},"*:Datatable")
	beegae.Router("/nfc/:nfcId/delete",&controllers.NfcController{},"*:DeleteNFC")
	beegae.Router("/nfc/:nfcId/edit",&controllers.NfcController{},"*:EditNFC")

	//Farsana
	beegae.Router("/customer/add", &controllers.CustomerController{}, "*:AddCustomer")
	beegae.Router("/customer", &controllers.CustomerController{}, "*:CustomerDetails")
	beegae.Router("/customer/:customerid/delete", &controllers.CustomerController{}, "*:DeleteCustomer")
	beegae.Router("/customer/:customerid/edit", &controllers.CustomerController{}, "*:EditCustomer")

	beegae.Router("/group/add", &controllers.GroupController{}, "*:AddGroup")
	beegae.Router("/group", &controllers.GroupController{}, "*:GroupDetails")
	beegae.Router("/group/:groupId/delete", &controllers.GroupController{}, "*:DeleteGroup")
	beegae.Router("/group/:groupId/edit", &controllers.GroupController{}, "*:EditGroup")

	beegae.Router("/invite/add", &controllers.InviteUserController{}, "*:AddInvitation")
	beegae.Router("/invite", &controllers.InviteUserController{}, "*:InvitationDetails")
	beegae.Router("/invite/:inviteuserid/delete", &controllers.InviteUserController{}, "*:DeleteInvitation")
	beegae.Router("/invite/:inviteuserid/edit", &controllers.InviteUserController{}, "*:EditInvitation")


	//Aswathy
	beegae.Router("/contact/add", &controllers.ContactUserController{},"*:LoadContact")
	beegae.Router("/contact", &controllers.ContactUserController{},"*:LoadContactdetail")
	beegae.Router("/task/add", &controllers.TaskController{},"*:AddNewTask")
	beegae.Router("/task", &controllers.TaskController{},"*:LoadTaskDetail")
	beegae.Router("/job/add", &controllers.JobController{},"*:LoadJob")
	beegae.Router("/job", &controllers.JobController{},"*:LoadJobDetail")
	beegae.Router("/contact/:contactId/edit", &controllers.ContactUserController{},"*:LoadEditContact")
	beegae.Router("/job/:jobId/edit", &controllers.JobController{},"*:LoadEditJob")
	beegae.Router("/task/:taskId/edit", &controllers.TaskController{},"*:LoadEditTask")
	beegae.Router("/contact/:contactId/delete", &controllers.ContactUserController{},"*:LoadDeleteContact")
	beegae.Router("/job/:jobId/delete", &controllers.JobController{},"*:LoadDeleteJob")
	beegae.Router("/task/:taskId/delete", &controllers.TaskController{},"*:LoadDeleteTask")
}
