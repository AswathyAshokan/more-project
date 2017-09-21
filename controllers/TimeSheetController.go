package controllers
import (

	"app/passporte/models"
	"app/passporte/viewmodels"
	"reflect"
	//"app/passporte/helpers"
	"strconv"
	"strings"

	"log"
)
type TimeSheetController struct {
	BaseController
}

func (c *TimeSheetController)LoadTimeSheetDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	var keySlice []string
	viewModel := viewmodels.TimeSheetViewModel{}
	timeSheet :=models.TimeSheet{}
	dbStatus, timeSheetDetails := timeSheet.RetrieveTimeSheetFromDB(c.AppEngineCtx,companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(timeSheetDetails)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for i:=0;i<len(keySlice);i++{
				dbStatus, timeSheetUserDetails := timeSheet.RetrieveTimeSheetUserFromDB(c.AppEngineCtx,companyTeamName,keySlice[i])
				switch dbStatus {
				case true:
					dataValue := reflect.ValueOf(timeSheetUserDetails)
					for _, key := range dataValue.MapKeys() {
						var logDetailsSlice []string
						var workDetailsSlice []string
						log.Println("inside",timeSheetUserDetails[key.String()].TaskId)
						var dailyWorkStartTimeSlice  []string
						var dailyWorkEndTimeSlice   []string
						if len(timeSheetUserDetails[key.String()].TaskId)!=0{
							log.Println("insideaaaa")
							var dailyTaskStartTimeSlice []string
							var dailyTaskEndTimeSlice    []string

							logDetailsSlice=append(logDetailsSlice,timeSheetUserDetails[key.String()].UserName)
							logDetailsSlice=append(logDetailsSlice,timeSheetUserDetails[key.String()].TaskName)

							logDetailsSlice=append(logDetailsSlice,"1")
							logDetailsSlice=append(logDetailsSlice,"0")
							if len(timeSheetUserDetails[key.String()].TaskStartTime)!=0{
								dataValue := reflect.ValueOf(timeSheetUserDetails[key.String()].TaskStartTime)
								for _, taskStartTimeKey := range dataValue.MapKeys() {
									StartTime := strconv.FormatInt(timeSheetUserDetails[key.String()].TaskStartTime[taskStartTimeKey.String()].Time, 10)
									dailyTaskStartTimeSlice=append(dailyTaskStartTimeSlice,StartTime)

								}

								logDetailsSlice=append(logDetailsSlice,dailyTaskStartTimeSlice[0])
							}else{
								logDetailsSlice=append(logDetailsSlice,"")
							}
							taskDateFrom := strconv.FormatInt(timeSheetUserDetails[key.String()].TaskDateFrom, 10)
							logDetailsSlice =append(logDetailsSlice,taskDateFrom)
							DailyStartTime := strconv.FormatInt(timeSheetUserDetails[key.String()].DailyStartTime, 10)
							logDetailsSlice =append(logDetailsSlice,DailyStartTime)
							if len(timeSheetUserDetails[key.String()].TaskEndTime) !=0{
								log.Println(("log end timeee"))
								dataValueOfEndTime := reflect.ValueOf(timeSheetUserDetails[key.String()].TaskEndTime)
								for _, taskEndTimeKey := range dataValueOfEndTime.MapKeys() {
									EndTime := strconv.FormatInt(timeSheetUserDetails[key.String()].TaskEndTime[taskEndTimeKey.String()].Time, 10)
									dailyTaskEndTimeSlice=append(dailyTaskEndTimeSlice,EndTime)

								}
								lengthOfSlice :=len(dailyTaskEndTimeSlice)
								logDetailsSlice=append(logDetailsSlice,dailyTaskEndTimeSlice[lengthOfSlice-1])
							}else{
								logDetailsSlice=append(logDetailsSlice,"")
							}

							TaskDateTo := strconv.FormatInt(timeSheetUserDetails[key.String()].TaskDateTo, 10)
							logDetailsSlice =append(logDetailsSlice,TaskDateTo)
							DailyEndTime := strconv.FormatInt(timeSheetUserDetails[key.String()].DailyEndTime, 10)
							logDetailsSlice =append(logDetailsSlice,DailyEndTime)
							logDetailsSlice =append(logDetailsSlice,keySlice[i])
							logDetailsSlice =append(logDetailsSlice,timeSheetUserDetails[key.String()].TaskId)
							logDetailsSlice =append(logDetailsSlice,timeSheetUserDetails[key.String()].Date)

						}

						viewModel.TaskTimeSheetDetail =append(viewModel.TaskTimeSheetDetail,logDetailsSlice)

						workDetailsSlice=append(workDetailsSlice,timeSheetUserDetails[key.String()].UserName)
						result := strings.Split(timeSheetUserDetails[key.String()].WorkLocation, ",")
						workDetailsSlice=append(workDetailsSlice,result[0]+","+result[1])
						workDetailsSlice=append(workDetailsSlice,"1")
						workDetailsSlice=append(workDetailsSlice,"0")
						log.Println("in1",workDetailsSlice)
						if len(timeSheetUserDetails[key.String()].WorkStartTime) !=0{
							dataValue := reflect.ValueOf(timeSheetUserDetails[key.String()].WorkStartTime)
							for _, workStartTimeKey := range dataValue.MapKeys() {
								StartTime := strconv.FormatInt(timeSheetUserDetails[key.String()].WorkStartTime[workStartTimeKey.String()].Time, 10)
								dailyWorkStartTimeSlice=append(dailyWorkStartTimeSlice,StartTime)

							}
							workDetailsSlice=append(workDetailsSlice,dailyWorkStartTimeSlice[0])

						}else{
							workDetailsSlice=append(workDetailsSlice,"")
						}
						DailyStartTime := strconv.FormatInt(timeSheetUserDetails[key.String()].DailyStartTime, 10)
						workDetailsSlice =append(workDetailsSlice,DailyStartTime)
						if len(timeSheetUserDetails[key.String()].WorkEndTime)!=0{
							dataValueOfWorkEndTime := reflect.ValueOf(timeSheetUserDetails[key.String()].WorkEndTime)
							for _, workEndTimeKey := range dataValueOfWorkEndTime.MapKeys() {
								EndTime := strconv.FormatInt(timeSheetUserDetails[key.String()].WorkEndTime[workEndTimeKey.String()].Time, 10)
								dailyWorkEndTimeSlice=append(dailyWorkEndTimeSlice,EndTime)

							}
							log.Println("end time")
							lengthOfSlice :=len(dailyWorkEndTimeSlice)
							log.Println("end time length",lengthOfSlice)
							log.Println("end time slice",dailyWorkEndTimeSlice)
							log.Println("end time value",dailyWorkEndTimeSlice[lengthOfSlice-1])

							workDetailsSlice=append(workDetailsSlice,dailyWorkEndTimeSlice[lengthOfSlice-1])
						}else{
							workDetailsSlice=append(workDetailsSlice,"")

						}

						DailyEndTime := strconv.FormatInt(timeSheetUserDetails[key.String()].DailyEndTime, 10)
						workDetailsSlice =append(workDetailsSlice,DailyEndTime)
						workDetailsSlice =append(workDetailsSlice,keySlice[i])
						workDetailsSlice =append(workDetailsSlice,timeSheetUserDetails[key.String()].Date)
						log.Println("workdetails  ",workDetailsSlice)
						log.Println("length of arrayy",len(viewModel.WorkTimeSheeetDetails))

							log.Println("insideeee sucesss")
							if len(viewModel.WorkTimeSheeetDetails) ==0{
								log.Println("n111")
								viewModel.WorkTimeSheeetDetails =append(viewModel.WorkTimeSheeetDetails,workDetailsSlice)
								log.Println("urrrr under",viewModel.WorkTimeSheeetDetails)
							}else{
								var condition=""
								for i :=0;i<len(viewModel.WorkTimeSheeetDetails);i++{
									log.Println("n222")
									log.Println("i1",viewModel.WorkTimeSheeetDetails[i][8])
									log.Println("i2",workDetailsSlice[8])
									log.Println("i3",viewModel.WorkTimeSheeetDetails[i][9] )
									log.Println("i4",workDetailsSlice[9])
									if viewModel.WorkTimeSheeetDetails[i][8] ==workDetailsSlice[8] && viewModel.WorkTimeSheeetDetails[i][9] ==workDetailsSlice[9]{
										condition ="true";
										break

									}else{
										condition ="false"
									}

								}
								if condition=="false"{
									viewModel.WorkTimeSheeetDetails =append(viewModel.WorkTimeSheeetDetails,workDetailsSlice)
								}
							}
					}
				case false:
				}
			}
			log.Println("tiem sheet user details",viewModel.TaskTimeSheetDetail)
		case false:
		}
	dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
	var notificationCount=0
	switch dbStatus {
	case true:

		notificationOfUser := reflect.ValueOf(notificationValue)
		for _, notificationUserKey := range notificationOfUser.MapKeys() {
			dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
			switch dbStatus {
			case true:
				notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
				for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
					var NotificationArray []string
					if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
						notificationCount=notificationCount+1;
					}
					NotificationArray =append(NotificationArray,notificationUserKey.String())
					NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
					date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
					NotificationArray =append(NotificationArray,date)
					viewModel.NotificationArray=append(viewModel.NotificationArray,NotificationArray)

				}
			case false:
			}
		}
	case false:
	}
	viewModel.NotificationNumber=notificationCount
	viewModel.CompanyTeamName =companyTeamName
	viewModel.CompanyPlan = storedSession.CompanyPlan
	viewModel.AdminLastName =storedSession.AdminLastName
	viewModel.AdminFirstName =storedSession.AdminFirstName
	viewModel.ProfilePicture =storedSession.ProfilePicture
	c.Data["vm"] = viewModel
	c.TplName = "template/time-sheet.html"

	}

