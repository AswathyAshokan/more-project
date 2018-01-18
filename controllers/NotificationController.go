package controllers
import (

	//"app/passporte/viewmodels"
	//
	"app/passporte/models"
	//"encoding/json"
	//"log"
	"log"
	"gopkg.in/sendgrid/sendgrid-go.v2"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"


)
type NotificationController struct {
	BaseController
}

//to Display Plan Details
func (c *NotificationController) NotificationUpdate() {
	log.Println("notificationnnnnnnnnnnnnnnnnnnnnnnnnn")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	UpdateIdArray := c.GetStrings("DeletedId")
	expiryId := c.GetStrings("DeletedExpiryId")
	userId := c.GetStrings("DeletedUserId")
	log.Println("uuuitititit",expiryId)
	log.Println("userid",userId)
	w := c.Ctx.ResponseWriter
	dbStatus:= models.UpdateAllNotifications(c.AppEngineCtx,companyTeamName,UpdateIdArray,expiryId,userId)
	switch dbStatus {
	case true:
	//w.Write([]byte("true"))

	case false:
		w.Write([]byte("false"))

	}
}

func (c *NotificationController) NotificationDelete() {
	log.Println("deleteeeeeeee")
	UpdateIdArray := c.GetStrings("DeletedId")
	expiryId := c.GetStrings("DeletedExpiryId")
	userId := c.GetStrings("DeletedUserId")
	//log.Println("deletedArray ##%uyyuu",UpdateIdArray)
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	w := c.Ctx.ResponseWriter
	dbStatus:= models.DeleteAllNotifications(c.AppEngineCtx,companyTeamName,UpdateIdArray,expiryId,userId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))

	case false:
		w.Write([]byte("false"))

	}
}

func (c *NotificationController) EmailToUser() {
	log.Println("iam here for send a email")
	w := c.Ctx.ResponseWriter
	r :=c.Ctx.Request
	emailId := c.Ctx.Input.Param(":emailId")
	log.Println("userrrrid",emailId)
	//employeeName :=c.Ctx.Input.Param(":employeeName")
	notificationString := c.Ctx.Input.Param(":notificationString")
	notificationMessage := c.GetStrings("notification")
	log.Println("notiiiiiiiiiiiiiiiiiiiiii",notificationMessage)
	//location := c.Ctx.Input.Param(":location")
	//taskName := c.Ctx.Input.Param(":taskName")
	//delayTime := c.Ctx.Input.Param(":delayTime")
	//log.Println("emailId",emailId,employeeName,notificationString,location,taskName,delayTime)
	//body :="Dear member, we received a request for password change .this is your automatic genereted key "
	////+"Go to site to set your new password. The key will be active for 10 minutes"
	//
	////"Regards,"+
	////"The Passporte team"
	//from := "passportetest@gmail.com"
	//to := "aswathyashok85@gmail.com"
	//subject := "Subject: Passporte - Forgot Password\n"
	//mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	//message := []byte(subject + mime + "\n" + body)
	//if err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", "passportetest@gmail.com", "passporte123", "smtp.gmail.com"), from, []string{to}, []byte(message)); err != nil {
	//	log.Println(err)
	//}

	key := "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"
	sg := sendgrid.NewSendGridClientWithApiKey(key)

	// must change the net/http client to not use default transport
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	sg.Client = client // <-- now using urlfetch, "overriding" default

	message := sendgrid.NewMail()
	if notificationString =="After"{
		message.SetHTML(notificationMessage[0])
	} else {
		message.SetHTML(notificationMessage[0])

	}
	message.AddTo(emailId)
	message.SetFrom("passportetest@gmail.com")
	message.SetSubject("Work delay ")
	if e := sg.Send(message); e == nil {
		log.Println("lllllll")
	} else {

		log.Println("error",e)
	}
	w.Write([]byte("true"))
}