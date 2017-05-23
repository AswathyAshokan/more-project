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
	Status         		string
}
type LeaveRequests   struct {
	Info     	LeaveInfo
	Settings 	LeaveSettings

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
func (m *LeaveRequests) AcceptLeaveRequestById( ctx context.Context,leaveId string,userId string)(bool)  {
	leaveDetailOfUser :=LeaveRequests{}
	leaveDetail :=LeaveRequests{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/LeaveRequests/"+userId+"/"+leaveId).Value(&leaveDetail)
	leaveDetailOfUser.Info.EndDate =leaveDetail.Info.EndDate
	leaveDetailOfUser.Info.NumberOfDays =leaveDetail.Info.NumberOfDays
	leaveDetailOfUser.Info.Reason =leaveDetail.Info.Reason
	leaveDetailOfUser.Info.StartDate =leaveDetail.Info.StartDate
	leaveDetailOfUser.Settings.DateOfCreation =leaveDetail.Settings.DateOfCreation
	leaveDetailOfUser.Settings.Status ="Accepted"
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId).Update(&leaveDetailOfUser)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true

}
func (m *LeaveRequests) RejectLeaveRequestById( ctx context.Context,leaveId string,userId string)(bool)  {
	leaveDetailOfUser :=LeaveRequests{}
	leaveDetail :=LeaveRequests{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/LeaveRequests/"+userId+"/"+leaveId).Value(&leaveDetail)
	leaveDetailOfUser.Info.EndDate =leaveDetail.Info.EndDate
	leaveDetailOfUser.Info.NumberOfDays =leaveDetail.Info.NumberOfDays
	leaveDetailOfUser.Info.Reason =leaveDetail.Info.Reason
	leaveDetailOfUser.Info.StartDate =leaveDetail.Info.StartDate
	leaveDetailOfUser.Settings.DateOfCreation =leaveDetail.Settings.DateOfCreation
	leaveDetailOfUser.Settings.Status ="Rejected"
	err = dB.Child("/LeaveRequests/"+ userId+"/"+leaveId).Update(&leaveDetailOfUser)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true

}