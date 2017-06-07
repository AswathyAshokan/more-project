package controllers

import (
	"github.com/tbalthazar/onesignal-go"
	"log"
	"encoding/json"
)
type ServerNotificationController struct {
	BaseController
}
func (c *ServerNotificationController)ServerNotificationDetails() {

	w := c.Ctx.ResponseWriter
	client := onesignal.NewClient(nil)
	//client.UserKey = "YourOneSignalUserKey"
	client.AppKey = "c4df0d1f-ac5e-4757-9446-2e3c75dbbebb"
	//app, res, err := client.Apps.Get("YourAppID")
	//
	////create the players
	//playerRequest := &onesignal.PlayerRequest{
	//	AppID:        "c4df0d1f-ac5e-4757-9446-2e3c75dbbebb",
	//	DeviceType:   1,
	//	Identifier:   "fakeidentifier2",
	//	Language:     "fake-language",
	//	Timezone:     -28800,
	//	GameVersion:  "1.0",
	//	DeviceOS:     "iOS",
	//	DeviceModel:  "iPhone5,2",
	//	AdID:         "fake-ad-id2",
	//	SDK:          "fake-sdk",
	//	SessionCount: 1,
	//	Tags: map[string]string{
	//		"a":   "1",
	//		"foo": "bar",
	//	},
	//	AmountSpent: 1.99,
	//	CreatedAt:   1395096859,
	//	Playtime:    12,
	//	BadgeCount:  1,
	//	LastActive:  1395096859,
	//	TestType:    1,
	//}
	//createRes, res, err := client.Players.Create(playerRequest)
	////create a new session for players
	//opt := &onesignal.PlayerOnSessionOptions{
	//	Identifier:  "FakeIdentifier",
	//	Language:    "en",
	//	Timezone:    -28800,
	//	GameVersion: "1.0",
	//	DeviceOS:    "7.0.4",
	//	AdID:        "fake-ad-id",
	//	SDK:         "fake-sdk",
	//	Tags: map[string]string{
	//		"a":   "1",
	//		"foo": "bar",
	//	},
	//}
	//
	////create a notification
	//successRes, res, err := client.Players.OnSession("196354561117-s7760vuj5f9k3mei3pvs6qv0b6gc1oma.apps.googleusercontent.com", opt)
	//
	//playerID := "aPlayerID"
	//notificationReq := &onesignal.NotificationRequest{
	//	AppID:            1:196354561117:android:e66ed333a8ea889d,
	//	Contents:         map[string]string{"en": "English message"},
	//	IsIOS:            true,
	//	IncludePlayerIDs: []string{playerID},
	//}
	//createRes, res, err := client.Notifications.Create(notificationReq)


	listOpt := &onesignal.PlayerListOptions{
		AppID:  client.AppKey,
		Limit:  10,
		Offset: 0,
	}
	listRes, res, err := client.Players.List(listOpt)
	if err != nil {
		log.Println("Error on APP connection: ",err)
	}
	log.Println("ListRes: ", listRes)
	log.Println("Res: ", res)
	slices := []interface{}{listRes, res,err}
	sliceToClient, _ := json.Marshal(slices)
	w.Write(sliceToClient)
	c.Layout = "layout/layout.html"
	c.TplName = "template/time-sheet.html"


}
