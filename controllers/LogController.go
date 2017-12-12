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
	log.Println("inside it")
	viewModel := viewmodels.WorkLogViewModel{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	logDetails :=models.WorkLog{}
	var userId []string
	dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(logUserDetail)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())

		}
		for _, key := range dataValue.MapKeys() {
			if logUserDetail[key.String()].Category =="TaskLocation"{
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
				tempValueSlice = append(tempValueSlice,tempTaskNames)
				latitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Latitude, 'f', 6, 64)
				longitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Longitude, 'f', 6, 64)
				tempValueSlice = append(tempValueSlice, latitudeInString)
				tempValueSlice = append(tempValueSlice,longitudeInString)
				viewModel.Values = append(viewModel.Values, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
				userId = append(userId,logUserDetail[key.String()].UserID)

			} else {
				//var buffer bytes.Buffer
				var tempValueSlice []string
				tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].UserName)
				tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].Type)
				tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].Duration)
				logTimeNew :=strconv.FormatInt(int64(logUserDetail[key.String()].LogTime), 10)
				tempValueSlice = append(tempValueSlice, logTimeNew)
				latitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Latitude, 'f', 6, 64)
				longitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Longitude, 'f', 6, 64)

				tempValueSlice = append(tempValueSlice, latitudeInString)
				tempValueSlice = append(tempValueSlice,longitudeInString)
				viewModel.WorkLocationValues = append(viewModel.WorkLocationValues, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]


			}
			userId = append(userId,logUserDetail[key.String()].UserID)

		}
		companyUser :=models.GetCompanyUsers(c.AppEngineCtx,companyTeamName)
		var generalKeySlice []string
		logStatus,generalLogData := models.GetGeneralLogDataByUserId(c.AppEngineCtx)
		switch logStatus {
		case true:
			dataValue := reflect.ValueOf(generalLogData)
			for _, key := range dataValue.MapKeys() {
				log.Println("key ?????",key.String())
				//userId = append(userId,key.String())

				for k :=0;k<len(companyUser);k++{
					if companyUser[k] == key.String(){
						generalKeySlice = append(generalKeySlice,key.String())


						LogData := models.GetSpecificLogValues(c.AppEngineCtx,key.String())
						if LogData.UserName !="" {
							var tempGeneralLogSlice []string
							tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.UserName)
							tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.Type)
							tempGeneralLogSlice = append(tempGeneralLogSlice,strconv.FormatInt(int64(LogData.LogTime), 10))
							tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.Duration)
							lastLoginLattitude := strconv.FormatFloat(LogData.Latitude, 'f', 6, 64)
							lastLoginLongitude := strconv.FormatFloat(LogData.Longitude, 'f', 6, 64)
							tempGeneralLogSlice = append(tempGeneralLogSlice,lastLoginLattitude)
							tempGeneralLogSlice = append(tempGeneralLogSlice,lastLoginLongitude)
							tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.UserID)
							tempGeneralLogSlice = append(tempGeneralLogSlice,LogData.LogDescription)
							viewModel.GeneralLogValues = append(viewModel.GeneralLogValues,tempGeneralLogSlice)
						}
						LogoutData := models.GetSpecificLogoutValues(c.AppEngineCtx,key.String())
						if LogoutData.UserName !=""{
							var tempGenerealLogoutSlice []string
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,LogoutData.UserName)
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,LogoutData.Type)
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,strconv.FormatInt(int64(LogoutData.LogTime), 10))
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,LogoutData.Duration)
							lastLogoutLattitude := strconv.FormatFloat(LogoutData.Latitude, 'f', 6, 64)
							lastLogoutLongitude := strconv.FormatFloat(LogoutData.Longitude, 'f', 6, 64)
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,lastLogoutLattitude)
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,lastLogoutLongitude)
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,LogoutData.UserID)
							tempGenerealLogoutSlice = append(tempGenerealLogoutSlice,LogoutData.LogDescription)
							viewModel.GeneralLogValues = append(viewModel.GeneralLogValues,tempGenerealLogoutSlice)
						}
						userTrackLog := models.GetUserTrackLogValues(c.AppEngineCtx,key.String())

						trackDataValue := reflect.ValueOf(userTrackLog)
						for _, trackKey := range trackDataValue.MapKeys() {
							eachUserTrackData := models.GetSpecificUserTrackDetails(c.AppEngineCtx,key.String(),trackKey.String())
							if eachUserTrackData.UserName !=""{
								var tempUserTrackSlice []string
								tempUserTrackSlice = append(tempUserTrackSlice,eachUserTrackData.UserName)
								tempUserTrackSlice = append(tempUserTrackSlice,eachUserTrackData.Type)
								tempUserTrackSlice = append(tempUserTrackSlice,strconv.FormatInt(int64(eachUserTrackData.LogTime), 10))
								tempUserTrackSlice = append(tempUserTrackSlice,"kkkkkkk")
								trackLatitude := strconv.FormatFloat(eachUserTrackData.Latitude, 'f', 6, 64)
								trackLongitude := strconv.FormatFloat(eachUserTrackData.Longitude, 'f', 6, 64)
								tempUserTrackSlice = append(tempUserTrackSlice,trackLatitude)
								tempUserTrackSlice = append(tempUserTrackSlice,trackLongitude)
								tempUserTrackSlice = append(tempUserTrackSlice,eachUserTrackData.UserID)
								tempUserTrackSlice = append(tempUserTrackSlice,eachUserTrackData.LogDescription)
								viewModel.GeneralLogValues = append(viewModel.GeneralLogValues,tempUserTrackSlice)
							}
						}



					}



				}
			}
			viewModel.GeneralKey = generalKeySlice
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
	log.Println("viewModel.GeneralLogValues !!!!!!!!!!!!!!!!!!!!!!!!",viewModel.GeneralLogValues)
	c.Data["vm"] = viewModel
	c.TplName = "template/log.html"
}



func (c *LogController)LoadActivityLogDetails() {
	log.Println("haiiiiiiii    iammmm hereee................................")
	//c.TplName = "template/log.html"


}
