package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"reflect"
	"app/passporte/helpers"
	"log"
	"strconv"
	"bytes"
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
		/*var tempGeneralLogSlice [][]string
		var lastLoginArray []string
		var lastLogoutArray []string
		logStatus,generalLogData := models.GetGeneralLogDataByUserId(c.AppEngineCtx)
		switch logStatus {
		case true:
			log.Println("generalLogData",generalLogData)
			dataValue := reflect.ValueOf(generalLogData)
			for _, key := range dataValue.MapKeys() {
				LogData := models.GetSpecificLogValues(c.AppEngineCtx,key.String())
				if LogData.LastLogin.
				lastLattitude := strconv.FormatFloat(LogData..Latitude, 'f', 6, 64)

				*//*logDataValue := reflect.ValueOf(LogData)
				for _, logKey := range logDataValue.MapKeys() {
					log.Println("logKey",logKey.String())
					var generalTempSlice   []string
					loginType := generalLogData[logKey.String()].Type
					logTime := generalLogData[logKey.String()].LogTime
					if loginType == helpers.LoginStatus{

						log.Println("logKeygggggg",logKey.String())
						unixLoginTimeArray = append(unixLoginTimeArray,logKey.String())
						UnixLoginDate := time.Unix(logTime, 0)
						unixLoginTimeArray =append(unixLoginTimeArray,UnixLoginDate.String())



					} else {

						log.Println("logKey",logKey)
						uniLogOutTimeArray = append(uniLogOutTimeArray,logKey.String())
						UnixLogOutTime := time.Unix(logTime, 0)
						uniLogOutTimeArray = append(uniLogOutTimeArray,UnixLogOutTime.String())


					}
					viewModel.GeneralLogin = append(viewModel.GeneralLogin,unixLoginTimeArray)
					viewModel.GeneralLogin = append(viewModel.GeneralLogin,uniLogOutTimeArray)
					unixLoginTimeArray =unixLoginTimeArray[:0]
					uniLogOutTimeArray =uniLogOutTimeArray[:0]


					latitudeInString :=strconv.FormatFloat(LogData[logKey.String()].Latitude, 'f', 6, 64)
					longitudeInString :=strconv.FormatFloat(LogData[logKey.String()].Longitude, 'f', 6, 64)
					generalTempSlice = append(generalTempSlice, latitudeInString)
					generalTempSlice = append(generalTempSlice,longitudeInString)
					generalTempSlice = append(generalTempSlice,LogData[logKey.String()].UserID)
					generalTempSlice = append(generalTempSlice,LogData[logKey.String()].LogDescription)
					generalTempSlice = append(generalTempSlice,LogData[logKey.String()].Type)
					generalTempSlice = append(generalTempSlice,LogData[logKey.String()].UserName)
					viewModel.GeneralLogValues = append(viewModel.GeneralLogValues,generalTempSlice)
					generalTempSlice = generalTempSlice[:0]


				}*//*
			}


			log.Println("unixLoginTimeArray",unixLoginTimeArray)

			unixLoginTimeArray = unixLoginTimeArray[:0]
			uniLogOutTimeArray = uniLogOutTimeArray[:0]


			log.Println("loginTime", unixLoginTimeArray)
		case false:

		}*/
		viewModel.Keys = keySlice
	case false :
		log.Println(helpers.ServerConnectionError)
	}









	/*switch logStatus {
	case true:
		log.Println("generalLogData",generalLogData)
		dataValue := reflect.ValueOf(generalLogData)
		for _, key := range dataValue.MapKeys() {
			loginType := generalLogData[key.String()].Type
			logTime := generalLogData[key.String()].LogTime
			if loginType == helpers.LoginStatus{
				unixLoginTimeArray = append(unixLoginTimeArray,generalLogData[key.String()].UserID)
				UnixLoginDate := time.Unix(logTime, 0)
				unixLoginTimeArray =append(unixLoginTimeArray,UnixLoginDate.String())

			} else {
				uniLogOutTimeArray = append(uniLogOutTimeArray,generalLogData[key.String()].UserID)
				UnixLogOutTime := time.Unix(logTime, 0)
				uniLogOutTimeArray = append(uniLogOutTimeArray,UnixLogOutTime.String())
			}

			var generalTempSlice   []string
			latitudeInString :=strconv.FormatFloat(generalLogData[key.String()].Latitude, 'f', 6, 64)
			longitudeInString :=strconv.FormatFloat(generalLogData[key.String()].Longitude, 'f', 6, 64)
			generalTempSlice = append(generalTempSlice, latitudeInString)
			generalTempSlice = append(generalTempSlice,longitudeInString)
			generalTempSlice = append(generalTempSlice,generalLogData[key.String()].UserName)
			generalTempSlice = append(generalTempSlice,generalLogData[key.String()].Type)
			viewModel.GeneralLogValues = append(viewModel.GeneralLogValues,generalTempSlice)
			log.Println("uuuuuuuuuuuu",generalTempSlice)
			generalTempSlice = generalTempSlice[:0]
		}
		viewModel.GeneralLogin = append(viewModel.GeneralLogin,unixLoginTimeArray)
		viewModel.GeneralLogin = append(viewModel.GeneralLogin,uniLogOutTimeArray)
		log.Println("viewModel",viewModel)


		log.Println("loginTime", unixLoginTimeArray)
	case false:

	}*/


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
