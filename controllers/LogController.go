package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"reflect"
	"app/passporte/helpers"
	"log"
	"strconv"
	"bytes"
	//"time"
	//"time"
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
	//var userId string
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(logUserDetail)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())

		}
		for _, key := range dataValue.MapKeys() {
			var buffer bytes.Buffer
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].UserName)
			tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].Type)
			tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].Duration)
			logTimeNew :=strconv.FormatInt(int64(logUserDetail[key.String()].LogTime), 10)
			tempValueSlice = append(tempValueSlice, logTimeNew)
			taskId := logUserDetail[key.String()].TaskID
			taskName,JobName := models.GetTaskDataById(c.AppEngineCtx, taskId)
			tempTaskNames := ""
			if len(JobName) != 0 {
				buffer.WriteString(taskName)
				buffer.WriteString("(")
				buffer.WriteString(JobName)
				buffer.WriteString(")")
				tempTaskNames = buffer.String()
				buffer.Reset()
			} else {
				buffer.WriteString(taskName)
				tempTaskNames = buffer.String()
				buffer.Reset()
			}
			//taskData := taskName+"("+JobName+")"
			tempValueSlice = append(tempValueSlice,tempTaskNames)
			latitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Latitude, 'f', 6, 64)
			longitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Longitude, 'f', 6, 64)
			tempValueSlice = append(tempValueSlice, latitudeInString)
			tempValueSlice = append(tempValueSlice,longitudeInString)
			/*logDate := time.Unix(logUserDetail[key.String()].LogTime, 0).Format("01/02/2006")
			tempValueSlice = append(tempValueSlice,logDate)*/
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]



		}
		var tempGeneralLogSlice []string
		var tempGenerealLogoutSlice []string
		/*var lastLoginArray []string
		var lastLogoutArray []string*/
		logStatus,generalLogData := models.GetGeneralLogDataByUserId(c.AppEngineCtx)
		switch logStatus {
		case true:
			log.Println("generalLogData",generalLogData)
			dataValue := reflect.ValueOf(generalLogData)
			for _, key := range dataValue.MapKeys() {
				LogData := models.GetSpecificLogValues(c.AppEngineCtx,key.String())
				log.Println("LogData",LogData)
				lastLoginLattitude := strconv.FormatFloat(LogData.Latitude, 'f', 6, 64)
				lastLoginLongitude := strconv.FormatFloat(LogData.Longitude, 'f', 6, 64)
				tempGeneralLogSlice = append(tempGeneralLogSlice,lastLoginLattitude)
				tempGeneralLogSlice = append(tempGeneralLogSlice,lastLoginLongitude)
				tempGeneralLogSlice = append(tempGeneralLogSlice,strconv.FormatInt(int64(LogData.LogTime), 10))
				tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.Type)
				tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.UserID)
				tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.UserName)
				tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.LogDescription)
				viewModel.GeneralLogValues = append(viewModel.GeneralLogValues,tempGeneralLogSlice)
				tempGeneralLogSlice = tempGeneralLogSlice[:0]

			}

			for _, key := range dataValue.MapKeys() {
				LogoutData := models.GetSpecificLogoutValues(c.AppEngineCtx,key.String())
				log.Println("LogData",LogoutData)
				lastLogoutLattitude := strconv.FormatFloat(LogoutData.Latitude, 'f', 6, 64)
				lastLogoutLongitude := strconv.FormatFloat(LogoutData.Longitude, 'f', 6, 64)
				tempGeneralLogSlice = append(tempGenerealLogoutSlice,lastLogoutLattitude)
				tempGeneralLogSlice = append(tempGenerealLogoutSlice,lastLogoutLongitude)
				tempGeneralLogSlice = append(tempGenerealLogoutSlice,strconv.FormatInt(int64(LogoutData.LogTime), 10))
				tempGeneralLogSlice = append(tempGenerealLogoutSlice,LogoutData.Type)
				tempGeneralLogSlice = append(tempGenerealLogoutSlice,LogoutData.UserID)
				tempGeneralLogSlice = append(tempGenerealLogoutSlice,LogoutData.UserName)
				tempGeneralLogSlice = append(tempGenerealLogoutSlice,LogoutData.LogDescription)
				viewModel.GeneralLogoutValues = append(viewModel.GeneralLogoutValues,tempGenerealLogoutSlice)
				tempGenerealLogoutSlice = tempGenerealLogoutSlice[:0]

			}

		log.Println("viewModel.GeneralLogoutValues",viewModel.GeneralLogoutValues)
			log.Println("viewModel.GeneralLogValues ",viewModel.GeneralLogValues )
		case false:

		}
		viewModel.Keys = keySlice
	case false :
		log.Println(helpers.ServerConnectionError)
	}




	viewModel.CompanyTeamName =companyTeamName
	viewModel.CompanyPlan = storedSession.CompanyPlan
	viewModel.AdminLastName =storedSession.AdminLastName
	viewModel.AdminFirstName =storedSession.AdminFirstName
	viewModel.ProfilePicture =storedSession.ProfilePicture
	c.Data["vm"] = viewModel
	c.TplName = "template/log.html"
}



func (c *LogController)LoadActivityLogDetails() {
	log.Println("haiiiiiiii    iammmm hereee................................")
	//c.TplName = "template/log.html"


}
