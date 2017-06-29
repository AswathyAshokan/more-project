package controllers

import (
	"github.com/tbalthazar/onesignal-go"
	//"google.golang.org/appengine/urlfetch"
	"log"
	"fmt"
	//"google.golang.org/appengine/urlfetch"
	//"google.golang.org/appengine/urlfetch"

)

type PushNotificationController struct {
	BaseController
}

// Add new groups to database
func (c *PushNotificationController) CreateNotification() {

	/*r := *http.Request{}
	ctx := appengine.NewContext(r)
	client1 := urlfetch.Client(ctx)
	resp, err := client1.Get("https://www.google.com/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("resp.Status: ", resp.Status)*/

	/*transport := &onesignal.Transport{
		Config:    tweetlibConfig,
		Token:     token,
		Transport: &urlfetch.Transport{Context: ctx},
	}*/

	client := onesignal.NewClient(nil)
	client.AppKey = "c4df0d1f-ac5e-4757-9446-2e3c75dbbebb"
	client.UserKey = "YTQxMGNkM2ItY2MyNS00Y2IyLWEyYjMtNGJiN2NhOGNmZDI1"
	log.Println("client: ", client)
	notifID := CreateNotifications(client)
	log.Println(notifID)
	/*client.UserKey =
	client.AppKey = "c4df0d1f-ac5e-4757-9446-2e3c75dbbebb"
	apps, res, err := client.Apps.List()
	log.Println("apps, res, err",apps, res, err)*/


}


func CreateNotifications(client *onesignal.Client) string {
	fmt.Println("### CreateNotifications ###")
	playerID := "b5e46b79-04dc-4627-bd14-61104b9c79d5" // valid
	// playerID := "83823c5f-53ce-4e35-be6a-a3f27e5d838f" // invalid
	notificationReq := &onesignal.NotificationRequest{
		AppID:            "c4df0d1f-ac5e-4757-9446-2e3c75dbbebb",
		Contents:         map[string]string{"en": "English message"},
		IsIOS:            true,
		IncludePlayerIDs: []string{playerID},
	}

	createRes, res, err := client.Notifications.Create(notificationReq)
	if err != nil {
		log.Println("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- createRes:%+v\n", createRes)
	fmt.Println()

	return createRes.ID
}
