package controllers
import (

	"app/passporte/models"
	"app/passporte/viewmodels"
	"time"
	"reflect"
	"app/passporte/helpers"
	"log"
	"strconv"
	"strings"
	"fmt"
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
	var duration []string
	dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)

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
			logTimeUser := logTimeNew.String()[11:16]
			logTimeInString := strconv.FormatInt(logUserDetail[key.String()].LogTime, 10)
			logTime, err := strconv.ParseInt(logTimeInString, 10, 64)
			if err != nil {
				log.Println(err)

			}
			timeStamp := time.Unix(logTime, 0)
			hr, min, _ := timeStamp.Clock()
			logTimeInMinutes :=hr*60+min
			duration = strings.Split(logUserDetail[key.String()].Duration, ":")
			log.Println(duration[0])
			log.Println(duration[1])
			durationInMinutesFirst, _ := strconv.Atoi(duration[0])
			durationInMinutesSecond, _ := strconv.Atoi(duration[1])
			durationInMinutes := durationInMinutesFirst*60+durationInMinutesSecond
			logBetween := logTimeInMinutes - durationInMinutes
			log.Println(logBetween)
			logHour := logBetween /60
			logMinutes :=logBetween % 60
			var logHourInString =string(logHour)
			var logMinutesInString =string(logMinutes)
			var prependLogHours =""
			var prependLogMinutes =""
			if len(logHourInString) ==1 {

				prependLogHours = fmt.Sprintf("%02d", logHour)
			} else {
				prependLogHours = string(logHour)
			}
			if len(logMinutesInString) ==1 {

				prependLogMinutes = fmt.Sprintf("%02d", logMinutes)
			} else {
				prependLogMinutes = string(logMinutes)
			}
			logTimeForUser := prependLogHours +":"+ prependLogMinutes
			tempValueSlice = append(tempValueSlice, logTimeForUser + "  to " + logTimeUser)
			latitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Latitude, 'f', 6, 64)
			longitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Longitude, 'f', 6, 64)
			tempValueSlice = append(tempValueSlice, latitudeInString + "," + longitudeInString)
			logDate := time.Unix(logUserDetail[key.String()].LogTime, 0).Format("01/02/2006")
			tempValueSlice = append(tempValueSlice,logDate)
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		viewModel.Keys =keySlice
	case false :
		log.Println(helpers.ServerConnectionError)
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
