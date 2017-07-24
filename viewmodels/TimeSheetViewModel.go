package viewmodels

import (
	"github.com/tbalthazar/onesignal-go"
	"net/http"
)


type TimeSheetViewModel struct {
	Nfs	*onesignal.PlayerListResponse
	Res	*http.Response
	Error	error
	logArray		[][]LogDetails


}
type LogDetails struct {
	LogTime		int64
	Type		string
	UserID		string
	UserName 	string
	TaskID		string

}
