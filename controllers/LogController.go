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
	var userId []string
	dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)
	log.Println("log data",logUserDetail)
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
				//taskData := taskName+"("+JobName+")"
				tempValueSlice = append(tempValueSlice,tempTaskNames)
				latitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Latitude, 'f', 6, 64)
				longitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Longitude, 'f', 6, 64)
				tempValueSlice = append(tempValueSlice, latitudeInString)
				tempValueSlice = append(tempValueSlice,longitudeInString)
				/*tempValueSlice = append(tempValueSlice,logUserDetail[key.String()].Category)
				tempValueSlice = append(tempValueSlice,logUserDetail[key.String()].WorkId)*/
				/*logDate := time.Unix(logUserDetail[key.String()].LogTime, 0).Format("01/02/2006")
				tempValueSlice = append(tempValueSlice,logDate)*/
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
				//taskId := logUserDetail[key.String()].TaskID
				/*taskName,JobName := models.GetTaskDataById(c.AppEngineCtx, taskId)
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
				log.Println("task name",tempTaskNames)
				tempValueSlice = append(tempValueSlice,tempTaskNames)*/
				latitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Latitude, 'f', 6, 64)
				longitudeInString :=strconv.FormatFloat(logUserDetail[key.String()].Longitude, 'f', 6, 64)

				tempValueSlice = append(tempValueSlice, latitudeInString)
				tempValueSlice = append(tempValueSlice,longitudeInString)
				/*tempValueSlice = append(tempValueSlice,logUserDetail[key.String()].Category)
				tempValueSlice = append(tempValueSlice,logUserDetail[key.String()].WorkId)*/
				/*logDate := time.Unix(logUserDetail[key.String()].LogTime, 0).Format("01/02/2006")
				tempValueSlice = append(tempValueSlice,logDate)*/
				viewModel.WorkLocationValues = append(viewModel.WorkLocationValues, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
				userId = append(userId,logUserDetail[key.String()].UserID)

			}




		}
		var tempGeneralLogSlice []string
		var tempGenerealLogoutSlice []string
		var generalKeySlice []string
		logStatus,generalLogData := models.GetGeneralLogDataByUserId(c.AppEngineCtx)
		switch logStatus {
		case true:
			dataValue := reflect.ValueOf(generalLogData)
			for _, key := range dataValue.MapKeys() {
				generalKeySlice = append(generalKeySlice,key.String())
				var tempArray []string
				for i:=0;i<len(userId);i++{

					exists := false
					for v := 0; v < i; v++ {
						if userId[v] == userId[i] {
							exists = true
							break
						}
					}
					// If no previous element exists, append this one.
					if !exists {
						tempArray = append(tempArray, userId[i])
					}
				}
				for k :=0;k<len(tempArray);k++{
					if tempArray[k] == key.String(){
						LogData := models.GetSpecificLogValues(c.AppEngineCtx,key.String())
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
						tempGeneralLogSlice = tempGeneralLogSlice[:0]

						LogoutData := models.GetSpecificLogoutValues(c.AppEngineCtx,key.String())
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
						tempGenerealLogoutSlice = tempGenerealLogoutSlice[:0]



					}

				}
			}

			/*for _, key := range dataValue.MapKeys() {


			}*/
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
	c.Data["vm"] = viewModel
	c.TplName = "template/log.html"
}



func (c *LogController)LoadActivityLogDetails() {
	log.Println("haiiiiiiii    iammmm hereee................................")
	//c.TplName = "template/log.html"


}
