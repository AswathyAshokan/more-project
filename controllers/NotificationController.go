package controllers
import (

	//"app/passporte/viewmodels"
	//
	"app/passporte/models"
	//"encoding/json"
	//"log"
	"log"
)
type NotificationController struct {
	BaseController
}

//to Display Plan Details
func (c *NotificationController) NotificationUpdate() {
	log.Println("notificationnnnnnnnnnnnnnnnnnnnnnnnnn")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	w := c.Ctx.ResponseWriter
	dbStatus:= models.UpdateAllNotifications(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))

	case false:
		w.Write([]byte("false"))

	}
}

func (c *NotificationController) NotificationDelete() {
	log.Println("notificationnnnnnnnnnnnnnnnnnnnnnnnnn")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	w := c.Ctx.ResponseWriter
	dbStatus:= models.DeleteAllNotifications(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))

	case false:
		w.Write([]byte("false"))

	}
}