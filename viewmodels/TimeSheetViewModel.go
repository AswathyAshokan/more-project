package viewmodels

import (
	"github.com/tbalthazar/onesignal-go"
	"net/http"
)


type TimeSheetViewModel struct {
	Nfs	*onesignal.PlayerListResponse
	Res	*http.Response
	Error	error
	LogArray			[][]LogDetails
	TaskDetails			[][]string
	LeaveDetails			[][]string
	CompanyTeamName			string
	LogDetailsTask          	[][]string
	Keys				[]string
	CompanyPlan			string
	AdminFirstName			string
	AdminLastName			string
	ProfilePicture			string
	UserStartTimeAndEndTime         [][]string
	TaskTimeSheetDetail		[][]string
	WorkTimeSheeetDetails		[][]string


}
type LogDetails struct {
	LogTime		int64
	Type		string
	UserID		string
	UserName 	string
	TaskID		string
	LogDescription	string

}
