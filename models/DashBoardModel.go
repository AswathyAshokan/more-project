package models
import (
	"golang.org/x/net/context"
	"log"
	//"app/passporte/helpers"
)
func (m *TaskIdInfo)RetrieveTaskFromCompany(ctx context.Context,companyId string)(bool,map[string]TaskIdInfo) {
	companyTaskDetail := map[string]TaskIdInfo{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Company/"+companyId+"/Tasks/").Value(&companyTaskDetail)
	if err != nil {
		log.Fatal(err)
		return false, companyTaskDetail
	}

	return true, companyTaskDetail

}
func (m *TaskSetting)UpdateTaskStatus(ctx context.Context,taskId string,taskStatus string,completedPercentage float32,pendingPercentage float32)(bool) {

	taskSettings :=TaskSetting{}
	taskUpdateSettings :=TaskSetting{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Tasks/"+taskId+"/Settings/").Value(&taskSettings)
	taskUpdateSettings.TaskStatus =taskStatus
	taskUpdateSettings.DateOfCreation =taskSettings.DateOfCreation
	taskUpdateSettings.Status =taskSettings.Status
	taskUpdateSettings.FitToWorkDisplayStatus =taskSettings.FitToWorkDisplayStatus
	taskUpdateSettings.CompletedPercentage =completedPercentage
	taskUpdateSettings.PendingPercentage =pendingPercentage
	err = dB.Child("/Tasks/"+taskId+"/Settings/").Set(taskUpdateSettings)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}