package models
import (
	"golang.org/x/net/context"
	"log"
	//"encoding/json"
	//"golang.org/x/crypto/bcrypt"
	//"strings"
	//"app/passporte/helpers"

)
type LeaveInfo struct {

	StartDate    	 	int64
	EndDate    	 	int64
	NumberOfDays    	int64
	Reason			string
}
type LeaveSettings struct {
	DateOfCreation 		int64

}
type LeaveRequests   struct {
	Info     	LeaveInfo
	Settings 	LeaveSettings
	Company 	map[string]CompanyLeave

}
type CompanyLeave struct {
	CompanyName	string
	Status		string
}
func (m *LeaveRequests)GetAllLeaveRequest(ctx context.Context,userKeySlice []string)(bool,map[string]LeaveRequests) {
	leaveDetail :=  map[string]LeaveRequests{}
	//company :=Company{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	//for i := 0; i < len(userKeySlice); i++ {
		err = dB.Child("/LeaveRequests/").Value(&leaveDetail)
	//}
	//contactStatus := "Active";
	//err = dB.Child("Contacts").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).OrderBy("Settings/Status").EqualTo(contactStatus).Value(&leaveDetail)
	if err != nil {
		log.Fatal(err)
		return false,leaveDetail
	}
	return true,leaveDetail

}
func (m *LeaveRequests)GetAllLeaveRequestById(ctx context.Context,userKey string,companyId string)(bool,map[string]LeaveRequests,CompanyUsers,map[string]CompanyInvitations) {
	leaveDetailOfUser := map[string]LeaveRequests{}
	company :=CompanyUsers{}
	companyInvitation :=map[string]CompanyInvitations{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("Connection error:", err)
	}
	err = dB.Child("/LeaveRequests/"+userKey).Value(&leaveDetailOfUser)
	err = dB.Child("/Company/"+companyId+"/Users/"+userKey).Value(&company)
	err = dB.Child("/Company/"+companyId+"/Invitation").Value(&companyInvitation)
	if err != nil {
		log.Fatal(err)
		return false,leaveDetailOfUser,company,companyInvitation
	}
	return true,leaveDetailOfUser,company,companyInvitation
}
func (m *LeaveRequests) AcceptLeaveRequestById( ctx context.Context,leaveId string,userId string,companyTeamName string,companyName string)(bool)  {
	leaveDetailOfUser :=LeaveInfo{}
	//leaveDetailMap :=map[string]CompanyLeave{}
	leaveDetailStruct :=CompanyLeave{}
	leaveDetail :=LeaveRequests{}
	leaveSettings :=LeaveSettings{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/LeaveRequests/"+userId+"/"+leaveId).Value(&leaveDetail)
	leaveDetailOfUser.EndDate =leaveDetail.Info.EndDate
	leaveDetailOfUser.NumberOfDays =leaveDetail.Info.NumberOfDays
	leaveDetailOfUser.Reason =leaveDetail.Info.Reason
	leaveDetailOfUser.StartDate =leaveDetail.Info.StartDate
	leaveSettings.DateOfCreation =leaveDetail.Settings.DateOfCreation
	leaveDetailStruct.CompanyName =companyTeamName
	leaveDetailStruct.Status ="Accepted"
	//leaveDetailOfUser.Settings.Status ="Accepted"
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId+"/Info").Set(&leaveDetailOfUser)
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId+"/Settings").Set(&leaveSettings)
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId+"/Company/"+companyTeamName).Set(&leaveDetailStruct)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true

}
func (m *LeaveRequests) RejectLeaveRequestById( ctx context.Context,leaveId string,userId string,companyTeamName string,companyName string)(bool)  {
	leaveDetailOfUser :=LeaveInfo{}
	leaveDetail :=LeaveRequests{}
	leaveDetailStruct :=CompanyLeave{}
	leaveSettings :=LeaveSettings{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/LeaveRequests/"+userId+"/"+leaveId).Value(&leaveDetail)
	leaveDetailOfUser.EndDate =leaveDetail.Info.EndDate
	leaveDetailOfUser.NumberOfDays =leaveDetail.Info.NumberOfDays
	leaveDetailOfUser.Reason =leaveDetail.Info.Reason
	leaveDetailOfUser.StartDate =leaveDetail.Info.StartDate
	leaveSettings.DateOfCreation =leaveDetail.Settings.DateOfCreation
	leaveDetailStruct.Status ="Rejected"
	leaveDetailStruct.CompanyName =companyTeamName
	//leaveDetailOfUser.Settings.Status ="Rejected"
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId+"/Info").Set(&leaveDetailOfUser)
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId+"/Settings").Set(&leaveSettings)
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId+"/Company/"+companyTeamName).Set(&leaveDetailStruct)

	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true

}