package controllers

import (
	"github.com/NaySoftware/go-fcm"
	"fmt"
	"log"
)

type PushNotificationController struct {
	BaseController
}
const (
	serverKey = "AAAALbek1F0:APA91bEQrk375cmVRnGATV3B4PzYGYfcOmdPRSC7HTpbJPBKkkwCPaKjeKBhc75AcIsjMmysR9fH5YWYWvibQWujJERes7URVp2lMUVS9wSd8l3ikeGs_94YMfvUBYZIqlo7k-KYkNF0 "
)
// Add new groups to database
func (c *PushNotificationController) CreateNotification() {

	var NP fcm.NotificationPayload
	NP.Title="hello"
	NP.Body="world"
	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}
	ids := []string{
		"token1",
	}
	xds := []string{
		"token5",
		"token6",
		"token7",
	}
	d := fcm.NewFcmClient(serverKey)
	d.NewFcmRegIdsMsg(ids, data)
	d.AppendDevices(xds)
	d.SetNotificationPayload(&NP)
	log.Println("push",d)
	status, err := d.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}


