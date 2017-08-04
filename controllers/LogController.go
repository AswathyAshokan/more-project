package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"time"
	"reflect"
	"app/passporte/helpers"
	"log"
	"strconv"
)


type LogController struct {
	BaseController
}
func (c *LogController)LoadLogDetails() {
	viewModel := viewmodels.WorkLogViewModel{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	logDetails :=models.WorkLog{}
	//var duration []string
	dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)
	var userId string
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(logUserDetail)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())

		}
		for _, key := range dataValue.MapKeys() {

			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].UserName)
			tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].Type)
			tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].Duration)
			logTimeNew := time.Unix(logUserDetail[key.String()].LogTime, 0)
			tempValueSlice = append(tempValueSlice, logTimeNew.String())
			latitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Latitude, 'f', 6, 64)
			longitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Longitude, 'f', 6, 64)
			tempValueSlice = append(tempValueSlice, latitudeInString)
			tempValueSlice = append(tempValueSlice,longitudeInString)
			logDate := time.Unix(logUserDetail[key.String()].LogTime, 0).Format("01/02/2006")
			tempValueSlice = append(tempValueSlice,logDate)
			taskId := logUserDetail[key.String()].TaskID
			taskName,JobName := models.GetTaskDataById(c.AppEngineCtx, taskId)
			taskData := taskName+(JobName)
			userId = logUserDetail[key.String()].UserID
			tempValueSlice = append(tempValueSlice,taskData)
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		viewModel.Keys = keySlice
	case false :
		log.Println(helpers.ServerConnectionError)
	}
	userId = "LmPgZSOM6OXtkbU9VNLdaPnLd042"
	logStatus,generalLogData := models.GetGeneralLogDataByUserId(c.AppEngineCtx,userId)
	switch logStatus {

	case true:
		log.Println("generalLogData",generalLogData)
	case false:

	}


	viewModel.CompanyTeamName =companyTeamName
	viewModel.CompanyPlan = storedSession.CompanyPlan
	viewModel.AdminLastName =storedSession.AdminLastName
	viewModel.AdminFirstName =storedSession.AdminFirstName
	viewModel.ProfilePicture =storedSession.ProfilePicture
	log.Println("company plan",viewModel.CompanyPlan)
	log.Println("admin last",viewModel.AdminLastName)
	log.Println("admin first",viewModel.AdminFirstName)
	log.Println("profile ",viewModel.ProfilePicture)
	c.Data["vm"] = viewModel
	c.TplName = "template/log.html"
}



func (c *LogController)LoadActivityLogDetails() {
	log.Println("haiiiiiiii    iammmm hereee................................")
	//c.TplName = "template/log.html"


}
