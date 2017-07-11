package controllers
import (

	"app/passporte/models"
	//"app/passporte/viewmodels"

	"reflect"
	"app/passporte/helpers"
	"log"
	"time"
	"strconv"

	//"fmt"
	"github.com/kelvins/sunrisesunset"

)
type TimeSheetController struct {
	BaseController
}

func (c *TimeSheetController)LoadTimeSheetDetails() {
	//r := c.Ctx.Request
	//w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	var keySlice []string
	var keySliceForActiveTask []string
	var keySliceForActiveTaskUsers []string
	var keyForLog []string
	//var tempValueSlice []string
	logDetails := models.WorkLog{}
	task := models.Tasks{}

		dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(tasks)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for _, k := range keySlice {
				if tasks[k].Settings.TaskStatus == helpers.StatusCompleted {
					keySliceForActiveTask = append(keySliceForActiveTask, k)
					userValue := reflect.ValueOf(tasks[k].UsersAndGroups.User)
					for _, key := range userValue.MapKeys() {
						if tasks[k].UsersAndGroups.User[key.String()].Status == helpers.StatusActive {
							keySliceForActiveTaskUsers = append(keySliceForActiveTaskUsers, key.String())
						}

					}
				}
				dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)

				switch dbStatus {
				case true:
					//var userName string
					logValue := reflect.ValueOf(logUserDetail)
					for _, key := range logValue.MapKeys() {
						keyForLog = append(keyForLog, key.String())
					}
					log.Println("logggg",keyForLog)
					startTime := time.Unix(logUserDetail[keyForLog[0]].LogTime, 0)
					startTimeOfLog := startTime.String()[11:16]
					log.Println("log time",startTimeOfLog)
					utcOffset := -3.0
					date      := time.Date(2017, 3, 23, 0, 0, 0, 0, time.UTC)
					lat, _ := strconv.ParseFloat(tasks[k].Location.Latitude, 64)
					long, _ := strconv.ParseFloat(tasks[k].Location.Longitude, 64)

					// Calculate the sunrise and sunset times
					sunrise, sunset, err := sunrisesunset.GetSunriseSunset(lat, long, utcOffset, date)

					// If no error has occurred, print the results
					if err == nil {
						log.Println("Sunrise:", sunrise.Format("15:04:05")) // Sunrise: 06:11:44
						log.Println("Sunset:", sunset.Format("15:04:05")) // Sunset: 18:14:27
					} else {
						log.Println(err)
					}
					utc := time.Now().UTC()
					log.Println(utc)
					local := utc
					log.Println("location",tasks[k].Info.TaskLocation)
					location, err := time.LoadLocation("Asia/Delhi")
					if err == nil {
						local = local.In(location)
					}
					log.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))

					for i := 0; i < len(keySliceForActiveTaskUsers); i++ {
						for _, k := range keyForLog {

							if logUserDetail[k].UserID == keySliceForActiveTaskUsers[i] {
								startDateOfLog := time.Unix(logUserDetail[k].LogTime, 0).Format("2006/01/02")
								startDateOfTask :=time.Unix(tasks[k].Info.StartDate, 0).Format("2006/01/02")
								log.Println(startDateOfLog)
								log.Println(startDateOfTask)

								loc, _ := time.LoadLocation(tasks[k].Info.TaskLocation)
								now := time.Now().In(loc)
								log.Println("utc time",now)
								if startDateOfLog ==startDateOfTask{
									log.Println("start time",startTimeOfLog)
								}

							}
						}
					}
				//tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].UserName)
				}
			}
		}
		c.TplName = "template/time-sheet.html"

	}

