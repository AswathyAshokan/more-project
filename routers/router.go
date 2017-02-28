/*Author: Sarath
Date:01/02/2017*/
package routers

import (
	"app/passporte/controllers"
	"github.com/astaxie/beegae"
)

func init() {

	//Sarath
	beegae.Router("/", &controllers.LoginController{}, "*:Root")
	beegae.Router("/login", &controllers.LoginController{}, "*:Login")
	beegae.Router("/register", &controllers.RegisterController{}, "*:Register")
	beegae.Router("/isEmailUsed/:emailId",&controllers.RegisterController{},"*:CheckEmail")
	beegae.Router("/:companyTeamName/nfc", &controllers.NfcController{},"*:NFCDetails")
	beegae.Router("/:companyTeamName/nfc/add",&controllers.NfcController{},"*:AddNFC")
	//beegae.Router("/datatable",&controllers.NfcController{},"*:Datatable")
	beegae.Router("/:companyTeamName/nfc/:nfcId/delete",&controllers.NfcController{},"*:DeleteNFC")
	beegae.Router("/:companyTeamName/nfc/:nfcId/edit",&controllers.NfcController{},"*:EditNFC")
	beegae.Router("/logout",&controllers.LoginController{},"*:Logout")

	//Farsana
	beegae.Router("/:companyTeamName/customer/add", &controllers.CustomerController{}, "*:AddCustomer")
	beegae.Router("/:companyTeamName/customer", &controllers.CustomerController{}, "*:CustomerDetails")
	beegae.Router("/:companyTeamName/customer/:customerid/delete", &controllers.CustomerController{}, "*:DeleteCustomer")
	beegae.Router("/:companyTeamName/customer/:customerid/edit", &controllers.CustomerController{}, "*:EditCustomer")
	beegae.Router("/iscustomernameused/:customername/:type/:oldName", &controllers.CustomerController{}, "*:CustomerNameCheck")


	beegae.Router("/:companyTeamName/group/add", &controllers.GroupController{}, "*:AddGroup")
	beegae.Router("/:companyTeamName/group", &controllers.GroupController{}, "*:GroupDetails")
	beegae.Router("/:companyTeamName/group/:groupId/delete", &controllers.GroupController{}, "*:DeleteGroup")
	beegae.Router("/:companyTeamName/group/:groupId/edit", &controllers.GroupController{}, "*:EditGroup")

	beegae.Router("/:companyTeamName/invite/add", &controllers.InviteUserController{}, "*:AddInvitation")
	beegae.Router("/:companyTeamName/invite", &controllers.InviteUserController{}, "*:InvitationDetails")
	beegae.Router("/:companyTeamName/invite/:inviteuserid/delete", &controllers.InviteUserController{}, "*:DeleteInvitation")
	beegae.Router("/:companyTeamName/invite/:inviteuserid/edit", &controllers.InviteUserController{}, "*:EditInvitation")

	beegae.Router("/plan", &controllers.PlanController{}, "*:PlanDetails")
	beegae.Router("/planupdate", &controllers.PlanController{}, "*:PlanUpdate")


	beegae.Router("/:companyTeamName/customer-management", &controllers.CustomerManagementController{}, "*:CustomerManagement")

	//Aswathy
	beegae.Router("/:companyTeamName/contact/add", &controllers.ContactUserController{},"*:AddNewContact")
	beegae.Router("/:companyTeamName/contact", &controllers.ContactUserController{},"*:DisplayContactDetails")
	beegae.Router("/:companyTeamName/contact/:contactId/edit", &controllers.ContactUserController{},"*:LoadEditContact")
	beegae.Router("/:companyTeamName/contact/:contactId/delete", &controllers.ContactUserController{},"*:LoadDeleteContact")

	beegae.Router("/:companyTeamName/job/add", &controllers.JobController{},"*:AddNewJob")
	beegae.Router("/:companyTeamName/job", &controllers.JobController{},"*:LoadJobDetail")
	beegae.Router("/:companyTeamName/job/:jobId/edit", &controllers.JobController{},"*:LoadEditJob")
	beegae.Router("/:companyTeamName/job/:jobId/delete", &controllers.JobController{},"*:LoadDeleteJob")
	beegae.Router("/isJobNameUsed/:jobName", &controllers.JobController{},"*:CheckJobName")
	beegae.Router("/isJobNumberUsed/:jobNumber", &controllers.JobController{},"*:CheckJobNumber")

	beegae.Router("/:companyTeamName/task/add", &controllers.TaskController{},"*:AddNewTask")
	beegae.Router("/:companyTeamName/task", &controllers.TaskController{},"*:LoadTaskDetail")
	beegae.Router("/:companyTeamName/task/:taskId/edit", &controllers.TaskController{},"*:LoadEditTask")
	beegae.Router("/:companyTeamName/task/:taskId/delete", &controllers.TaskController{},"*:LoadDeleteTask")


	//View sections
	beegae.Router("/:companyTeamName/customer/:customerId/job", &controllers.JobController{},"*:LoadJobDetail")
	beegae.Router("/:companyTeamName/job/:jobId/task", &controllers.TaskController{},"*:LoadTaskDetail")

	//Login Bypass
	beegae.Router("/bypass", &controllers.ByPassController{},"*:ByPass")

}
