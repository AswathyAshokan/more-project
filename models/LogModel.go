package models

import (
	"golang.org/x/net/context"
	"log"
	//"app/passporte/helpers"
)

type WorkLog struct {
	Duration 	string
	Latitude	float64
	LogDescription	string
	LogTime		int64
	Longitude	float64
	Type		string
	UserID		string
	UserName 	string

}
func (m *WorkLog)GetLogDetailOfUser(ctx context.Context,companyTeamName string)(bool,map[string]WorkLog) {
	workDetail := map[string]WorkLog{}
	dB, err := GetFirebaseClient(ctx,"")
	//contactStatus := "Active";
	log.Println("model",companyTeamName)
	err = dB.Child("WorkLog/"+companyTeamName).Value(&workDetail)
	log.Println("workDetail",workDetail)
	if err != nil {
		log.Fatal(err)
		return false, workDetail
	}
	return true, workDetail

}