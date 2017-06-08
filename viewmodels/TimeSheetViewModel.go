package viewmodels

import (
	"github.com/tbalthazar/onesignal-go"
	"net/http"
)


type TimeSheetViewModel struct {
	Nfs	*onesignal.PlayerListResponse
	Res	*http.Response
	Error	error


}

