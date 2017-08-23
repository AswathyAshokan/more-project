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
	beegae.Router("/:companyTeamName/editProfile", &controllers.RegisterController{}, "*:EditProfile")
	beegae.Router("/:companyTeamName/changePassword",&controllers.RegisterController{},"*:ChangeAdminsPassword")
	beegae.Router("/:companyTeamName/isOldAdminPasswordCorrect/:oldPassword", &controllers.RegisterController{},"*:OldAdminPasswordCheck")
	beegae.Router("/forgotPassword", &controllers.RegisterController{}, "*:ForgotPassword")
	beegae.Router("/forgot-password/email-checking", &controllers.RegisterController{}, "*:CheckingEmailId")
	beegae.Router("/forgot-password/passwordReset", &controllers.RegisterController{}, "*:ResetPassword")
	beegae.Router("/register/getstate",&controllers.RegisterController{},"*:GetStates")

	beegae.Router("/isEmailUsed/:emailId",&controllers.RegisterController{},"*:CheckEmail")
	beegae.Router("/:companyTeamName/nfc", &controllers.NfcController{},"*:NFCDetails")
	beegae.Router("/:companyTeamName/nfc/add",&controllers.NfcController{},"*:AddNFC")
	//beegae.Router("/datatable",&controllers.NfcController{},"*:Datatable")
	beegae.Router("/:companyTeamName/nfc/:nfcId/delete",&controllers.NfcController{},"*:DeleteNFC")
	beegae.Router("/:companyTeamName/nfc/:nfcId/edit",&controllers.NfcController{},"*:EditNFC")
	beegae.Router("/logout", &controllers.LoginController{},"*:Logout")
	beegae.Router("/duress", &controllers.DuressController{},"*:Duress")

	//Farsana
	beegae.Router("/:companyTeamName/customer/add", &controllers.CustomerController{}, "*:AddCustomer")
	beegae.Router("/:companyTeamName/customer", &controllers.CustomerController{}, "*:CustomerDetails")
	beegae.Router("/:companyTeamName/customer/:customerid/delete", &controllers.CustomerController{}, "*:LoadDeleteCustomer")
	beegae.Router("/:companyTeamName/customer/:customerid/edit", &controllers.CustomerController{}, "*:EditCustomer")
	beegae.Router("/iscustomernameused/:customername/:type/:oldName", &controllers.CustomerController{}, "*:CustomerNameCheck")
	beegae.Router("/:companyTeamName/customer/:customerid/RemoveTask", &controllers.CustomerController{}, "*:RemoveCustomerFromTask")
	beegae.Router("/:companyTeamName/customer/:customerid/deletionOfCustomer", &controllers.CustomerController{}, "*:DeleteCustomerIfNotInTask")
	beegae.Router("/:companyTeamName/customer/:customerid/deletionOfCustomerIfUsedForJob", &controllers.CustomerController{}, "*:DeleteCustomerIfUsedForJob")
	beegae.Router("/:companyTeamName/customer/:customerid/deletionOfCustomerFromJob", &controllers.CustomerController{}, "*:DeleteCustomerFromJob")




	beegae.Router("/:companyTeamName/group/add", &controllers.GroupController{}, "*:AddGroup")
	beegae.Router("/:companyTeamName/group", &controllers.GroupController{}, "*:GroupDetails")
	beegae.Router("/:companyTeamName/group/:groupId/delete", &controllers.GroupController{}, "*:LoadDeleteGroup")
	beegae.Router("/:companyTeamName/group/:groupId/edit", &controllers.GroupController{}, "*:EditGroup")
	beegae.Router("/isgroupnameused/:groupName/:type/:oldName", &controllers.GroupController{}, "*:GroupNameCheck")
	beegae.Router("/:companyTeamName/group/:groupId/RemoveTask", &controllers.GroupController{}, "*:RemoveGroupFromTask")
	beegae.Router("/:companyTeamName/group/:groupId/deletionOfGroup", &controllers.GroupController{}, "*:DeleteGroupIfNotInTask")

	beegae.Router("/:companyTeamName/invite/add", &controllers.InviteUserController{}, "*:AddInvitation")
	beegae.Router("/:companyTeamName/invite", &controllers.InviteUserController{}, "*:InvitationDetails")
	beegae.Router("/:companyTeamName/invite/:inviteuserid/delete", &controllers.InviteUserController{}, "*:DeleteInvitation")
	beegae.Router("/:companyTeamName/invite/:inviteuserid/edit", &controllers.InviteUserController{}, "*:EditInvitation")
	beegae.Router("/:companyTeamName/invite/:inviteuserid/RemoveTask", &controllers.InviteUserController{}, "*:RemoveUserFromTask")
	beegae.Router("/:companyTeamName/invite/:numberOfUsers/AddExtraUserByUpgradePlan", &controllers.InviteUserController{}, "*:AddInvitationByUpgradationOfPlan")

	beegae.Router("/:companyTeamName/invite/:inviteuserid/deletionOfUser", &controllers.InviteUserController{}, "*:DeleteUserIfNotInTask")
	/*beegae.Router("/:companyTeamName/invite/:inviteuserid/ChangeInactiveToactive", &controllers.InviteUserController{}, "*:UpdateStatusWhenResponsePending")*/
	beegae.Router("/plan", &controllers.PlanController{}, "*:PlanDetails")
	beegae.Router("/plan/update", &controllers.PlanController{}, "*:PlanUpdate")

	/*beegae.Router("/superadmin", &controllers.SuperAdminController{}, "*:AddSuperAdmin")*/
	beegae.Router("/logoutForSuperAdmin", &controllers.LoginController{},"*:LogoutForSuperAdmin")
	beegae.Router("/customer-management", &controllers.CustomerManagementController{}, "*:CustomerManagement")
	beegae.Router("/customer-management/:customermanagementid/delete", &controllers.CustomerManagementController{}, "*:LoadDeleteCustomerManagement")

	beegae.Router("/accounts",&controllers.AccountsController{},"*:SuperAdminsAccount")
	beegae.Router("/changePassword",&controllers.AccountsController{},"*:ChangeSuperAdminsPassword")
	beegae.Router("/isOldPasswordCorrect/:oldPassword", &controllers.AccountsController{},"*:OldPasswordCheck")
	beegae.Router("/:companyTeamName/shareddocuments/:inviteuserid",&controllers.SharedDocumentController{},"*:LoadSharedDocuments")

	beegae.Router("/:companyTeamName/consent/add", &controllers.ConsentReceiptController{}, "*:AddConsentReceipt")
	beegae.Router("/:companyTeamName/consent", &controllers.ConsentReceiptController{}, "*:LoadConsentReceipt")
	beegae.Router("/:companyTeamName/consent/:consentId/edit", &controllers.ConsentReceiptController{}, "*:EditConsentReceipt")
	beegae.Router("/:companyTeamName/consent/:consentId/delete", &controllers.ConsentReceiptController{}, "*:DeleteConsentReceipt")
	beegae.Router("/:companyTeamName/worklocation/add", &controllers.WorkLocationcontroller{}, "*:AddWorkLocation")
	beegae.Router("/:companyTeamName/worklocation/:worklocationid/edit", &controllers.WorkLocationcontroller{}, "*:EditWorkLocation")
	beegae.Router("/:companyTeamName/worklocation", &controllers.WorkLocationcontroller{}, "*:LoadWorkLocation")

	beegae.Router("/listCountry",&controllers.CountryController{},"*:ListCountries")


	//Aswathy
	beegae.Router("/:companyTeamName/contact/add", &controllers.ContactUserController{},"*:AddNewContact")
	beegae.Router("/:companyTeamName/contact", &controllers.ContactUserController{},"*:DisplayContactDetails")
	beegae.Router("/:companyTeamName/contact/:contactId/edit", &controllers.ContactUserController{},"*:LoadEditContact")
	beegae.Router("/:companyTeamName/contact/:contactId/delete", &controllers.ContactUserController{},"*:LoadDeleteContact")
	beegae.Router("/:companyTeamName/contact/:contactId/RemoveTask", &controllers.ContactUserController{}, "*:RemoveContactFromTask")
	beegae.Router("/:companyTeamName/contact/:contactId/deletionOfContact", &controllers.ContactUserController{}, "*:DeleteContactIfNotInTask")

	beegae.Router("/isPhoneNumberUsed/:phoneNumber", &controllers.ContactUserController{},"*:CheckPhoneNumberAdd")
	beegae.Router("/isemailAddressUsed/:emailAddress", &controllers.ContactUserController{},"*:CheckEmailAddressAdd")
	beegae.Router("/isPhoneNumberUsed/:phoneNumber/:type/:oldNumber", &controllers.ContactUserController{},"*:CheckPhoneNumber")
	beegae.Router("/isemailAddressUsed/:emailAddress/:type/:oldEmail", &controllers.ContactUserController{},"*:CheckEmailAddress")



	beegae.Router("/:companyTeamName/job/add", &controllers.JobController{},"*:AddNewJob")
	beegae.Router("/:companyTeamName/job", &controllers.JobController{},"*:LoadJobDetail")
	beegae.Router("/:companyTeamName/job/:jobId/edit", &controllers.JobController{},"*:LoadEditJob")
	beegae.Router("/:companyTeamName/job/:jobId/delete", &controllers.JobController{},"*:LoadDeleteJob")
	beegae.Router("/isJobNameUsed/:jobName", &controllers.JobController{},"*:CheckJobName")
	beegae.Router("/isJobNumberUsed/:jobNumber", &controllers.JobController{},"*:CheckJobNumber")
	beegae.Router("/:companyTeamName/job/:jobId/RemoveTask", &controllers.JobController{}, "*:RemoveJobFromTask")
	beegae.Router("/:companyTeamName/job/:jobId/deletionOfJob", &controllers.JobController{}, "*:DeleteJobIfNotInTask")

	beegae.Router("/:companyTeamName/task/add", &controllers.TaskController{},"*:AddNewTask")
	beegae.Router("/:companyTeamName/task", &controllers.TaskController{},"*:LoadTaskDetail")
	beegae.Router("/:companyTeamName/task/:taskId/edit", &controllers.TaskController{},"*:LoadEditTask")
	beegae.Router("/:companyTeamName/task/:taskId/delete", &controllers.TaskController{},"*:LoadDeleteTask")
	beegae.Router("/:companyTeamName/task/:taskId/taskStatus", &controllers.TaskController{},"*:LoadTaskStatus")
	beegae.Router("/:companyTeamName/task/add/:jobName/:customerName", &controllers.TaskController{},"*:AddNewTask")


	beegae.Router("/:companyTeamName/leave", &controllers.LeaveController{},"*:LoadUserLeave")
	beegae.Router("/:companyTeamName/leave/:leaveKey/:userKey/accept", &controllers.LeaveController{},"*:LoadAcceptUserLeave")
	beegae.Router("/:companyTeamName/leave/:leaveKey/:userKey/reject", &controllers.LeaveController{},"*:LoadRejectUserLeave")
	beegae.Router("/:companyTeamName/fitToWork/add", &controllers.FitToWorkController{},"*:AddNewFitToWork")
	beegae.Router("/:companyTeamName/fitToWork", &controllers.FitToWorkController{}, "*:LoadFitToWork")
	beegae.Router("/:companyTeamName/fitToWork/:fitToWorkId/edit", &controllers.FitToWorkController{}, "*:EditFitToWork")
	beegae.Router("/:companyTeamName/fitToWork/:fitToWorkId/delete", &controllers.FitToWorkController{}, "*:DeleteFitToWork")
	beegae.Router("/:companyTeamName/isFitToWorkNameUsed/:fitWorkName", &controllers.FitToWorkController{},"*:CheckFitToWork")
	beegae.Router("/:companyTeamName/fitToWork/:fitToWorkId/deletionOfFitToWorkIfUsedForTask", &controllers.FitToWorkController{},"*:DeleteFitToWorkInTask")


	//View sections
	beegae.Router("/:companyTeamName/customer/:customerId/job", &controllers.JobController{},"*:LoadJobDetail")
	beegae.Router("/:companyTeamName/job/:jobId/task", &controllers.TaskController{},"*:LoadTaskDetail")

	//Login Bypass
	beegae.Router("/bypass", &controllers.ByPassController{},"*:ByPass")


	//Paypal sections'

	beegae.Router("/:companyTeamName/:companyPlan/payment",&controllers.PaymentController{},"*:Home")
	beegae.Router("/:companyTeamName/:companyPlan/paymentcancelreturn",&controllers.PaymentController{},"*:PaymentCancelReturn")
	beegae.Router("/:companyTeamName/:companyPlan/paymentsuccess" ,&controllers.PaymentController{},"*:PaymentSuccess" )
	beegae.Router("/:companyTeamName/:companyPlan/ipn" ,&controllers.PaymentController{},"*:IPN")

	//user log detail section


	beegae.Router("/:companyTeamName/workLog",&controllers.LogController{},"*:LoadLogDetails")
	beegae.Router("/:companyTeamName/activityworkLog",&controllers.LogController{},"*:LoadActivityLogDetails")

	//dash board section
	beegae.Router("/:companyTeamName/dashBoard", &controllers.DashBoardController{},"*:LoadDashBoard")
	beegae.Router("/:companyTeamName/timeSheet", &controllers.TimeSheetController{},"*:LoadTimeSheetDetails")
	beegae.Router("/push", &controllers.ServerNotificationController{},"*:ServerNotificationDetails")


	beegae.Router("/:companyTeamName/PendingWorks", &controllers.ComingSoonController{},"*:LoadComingSoonController")

	beegae.Router("/goroutine", &controllers.PushNotificationController{},"*:CreateNotification")



}
