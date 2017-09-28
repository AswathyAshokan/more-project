package models
import(
	"golang.org/x/net/context"
	"log"
	"github.com/kjk/betterguid"
	"strings"
	"time"
	"app/passporte/helpers"
	"reflect"
)


type FitToWork struct {
	FitToWorkName string
	FitToWork	map[string]TaskFitToWorks
	Settings	FitToWorkSettings
}
type TaskFitToWorks struct {
	Description    string
	Status         string
	DateOfCreation int64

}

type FitToWorkSettings struct{
	Status	string
}
func(m *FitToWork) AddFitToWorkToDb(ctx context.Context,instructionSlice []string ,companyTeamName string) (bool){
	log.Println("adddddddddddddd")
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	fitToWorkValue,err := db.Child("FitToWork/"+companyTeamName).Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	instructionMap := make(map[string]TaskFitToWorks)
	InstructionForFitToWork := TaskFitToWorks{}
	fitToWorkValueString := strings.Split(fitToWorkValue.String(),"/")
	fitToWorkUniqueID := fitToWorkValueString[len(fitToWorkValueString)-2]
	if instructionSlice[0] !=""{
		for i := 0; i < len(instructionSlice); i++ {
			InstructionForFitToWork.Description=instructionSlice[i]
			InstructionForFitToWork.DateOfCreation=(time.Now().UnixNano() / 1000000)
			InstructionForFitToWork.Status =helpers.StatusActive
			id := betterguid.New()
			instructionMap[id] = InstructionForFitToWork
			err = db.Child("/FitToWork/"+companyTeamName+"/"+ fitToWorkUniqueID +"/Instructions/").Set(instructionMap)
			if err != nil {
				log.Println(err)
				return false
			}
		}
	}

	return  true
}

func GetAllFitToWorkDetails(ctx context.Context) (bool,map[string]FitToWork){
	fitToWorkValue := map[string]FitToWork{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("FitToWork/").Value(&fitToWorkValue)
	if err != nil {
		log.Fatal(err)
		return false, fitToWorkValue
	}
	return true, fitToWorkValue
}
func GetSelectedCompanyName(ctx context.Context, fitToWorkId string)(map[string]FitToWork){
	fitToWork :=map[string]FitToWork{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("FitToWork/"+ fitToWorkId).Value(&fitToWork)
	if err != nil {
		log.Fatal(err)
	}
	return fitToWork

}
func GetAllInstructionsOfFitToWorkById(ctx context.Context,companyTeamName string, fitToWorkId string)(map[string]TaskFitToWorks)  {
	getInstructions:=map[string]TaskFitToWorks{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("FitToWork/"+companyTeamName+"/"+ fitToWorkId +"/Instructions").Value(&getInstructions)
	if err != nil {
		log.Fatal(err)
	}
	return getInstructions


}
func(m *FitToWork) UpdateFitToWorkToDb(ctx context.Context,instructionSlice []string ,companyTeamName string,fitToWorkId string ) (bool) {
	log.Println("cp6")
	fitToWorkDetails :=FitToWork{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/FitToWork/"+companyTeamName+"/"+fitToWorkId).Value(&fitToWorkDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	m.Settings.Status = fitToWorkDetails.Settings.Status
	err = db.Child("/FitToWork/"+companyTeamName+"/"+fitToWorkId).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	instructionMap := make(map[string]TaskFitToWorks)
	InstructionForFitToWork := TaskFitToWorks{}
	instructionMapForTask :=make(map[string]TaskFitToWork)
	InstructionForFitToWorkOnTask :=TaskFitToWork{}
	if instructionSlice[0] !="" {
		for i := 0; i < len(instructionSlice); i++ {
			InstructionForFitToWork.Description = instructionSlice[i]
			InstructionForFitToWork.DateOfCreation = (time.Now().UnixNano() / 1000000)
			InstructionForFitToWork.Status = helpers.StatusActive
			id := betterguid.New()
			InstructionForFitToWorkOnTask.Status = helpers.StatusActive
			InstructionForFitToWorkOnTask.DateOfCreation = (time.Now().UnixNano() / 1000000)
			InstructionForFitToWorkOnTask.Description = instructionSlice[i]
			instructionMap[id] = InstructionForFitToWork
			instructionMapForTask[id] =InstructionForFitToWorkOnTask
			err = db.Child("/FitToWork/" + companyTeamName + "/" + fitToWorkId + "/Instructions/").Set(instructionMap)
			if err != nil {
				log.Println(err)
				return false
			}
		}
	}
	fitToWorkUpdate :=FitToWorkForTask{}
	taskValue := map[string]Tasks{}
	fitToWorkUpdate.Settings.Status =m.Settings.Status
	fitToWorkUpdate.Info.TaskFitToWorkName =m.FitToWorkName
	fitToWorkUpdate.Info.FitToWorkId =fitToWorkId
	fitToWorkUpdate.FitToWorkInstruction =instructionMapForTask
	err = db.Child("Tasks").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&taskValue)
	dataValueOfFitToWorkForTask := reflect.ValueOf(taskValue)
	for _, taskKeys:=range dataValueOfFitToWorkForTask.MapKeys(){
		if taskValue[taskKeys.String()].Settings.Status ==helpers.StatusActive&& taskValue[taskKeys.String()].FitToWork.Info.FitToWorkId ==fitToWorkId{
			err = db.Child("/Tasks/"+taskKeys.String()+"/FitToWork").Set(fitToWorkUpdate)
		}
	}

	return true
}
func GetEachFitToWorkByCompanyId(ctx context.Context, fitToWorkId string,companyTeamName string)(FitToWork){
	log.Println("cp2",fitToWorkId)
	fitWork :=FitToWork{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("FitToWork/"+companyTeamName+"/"+ fitToWorkId).Value(&fitWork)
	log.Println("fit workssssssssssss",fitWork)
	if err != nil {
		log.Fatal(err)
	}
	return fitWork

}
func GetAllInstructionsFromFitToWork(ctx context.Context,fitToWorkId string,companyTeamName string)(map[string]TaskFitToWorks){
	log.Println("cp3")
	instructionOfFitWork :=map[string]TaskFitToWorks{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("FitToWork/"+companyTeamName+"/"+fitToWorkId+"/Instructions").Value(&instructionOfFitWork)
	if err != nil {
		log.Fatal(err)
	}
	return instructionOfFitWork

}
func DeleteFitToWorkById(ctx context.Context,fitToWorkId string,companyTeamName string)(bool)  {
	updatefitToWorkStatus := FitToWorkSettings{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}

	updatefitToWorkStatus.Status = helpers.UserStatusDeleted
	err = db.Child("FitToWork/"+companyTeamName+"/"+fitToWorkId+"/Settings").Update(&updatefitToWorkStatus)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true
}
func CheckFitWorkNameIsUsed(ctx context.Context, fitWorkName string,companyTeamName string)bool{
	fitToWork :=map[string]FitToWork{}
	fitWork :=map[string]FitToWork{}
	fullFitToWork :=map[string]FitToWork{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("FitToWork/").Value(&fitToWork)
	if err!=nil{
		log.Println("Error:",err)
	}
	FitToWorkDetails := reflect.ValueOf(fitToWork)
	for _, fitKey:=range FitToWorkDetails.MapKeys() {
		if fitKey.String() ==companyTeamName {
			log.Println("inside if")
			err = dB.Child("FitToWork/"+ companyTeamName).Value(&fullFitToWork)
		}

	}
	dataValueOfFitToWork := reflect.ValueOf(fullFitToWork)
	for _, fitKeys:=range dataValueOfFitToWork.MapKeys(){
		if fullFitToWork[fitKeys.String()].FitToWorkName == fitWorkName{
			log.Println("gggggg")
			return true
			break
		}
		log.Println("dfsgdfgdggfsgfdgdgd",fitWork)
	}
	return false
}
func (m *FitToWork) IsfitToWorkUsedForTask( ctx context.Context, fitToWorkId string,companyTeamName string)(bool, map[string]FitToWork)  {
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	fitToWorkValue := map[string]FitToWork{}
	err = dB.Child("FitToWork/"+ companyTeamName).Value(&fitToWorkValue)
	log.Println("job inside task",fitToWorkValue)

	return true,fitToWorkValue
}



func (m *Tasks) IsfitToWorkContainForTask( ctx context.Context, fitToWorkName string,companyTeamName string)(bool)  {
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	taskValue := map[string]Tasks{}
	err = dB.Child("Tasks").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&taskValue)
	dataValueOfFitToWorkForTask := reflect.ValueOf(taskValue)
	for _, taskKeys:=range dataValueOfFitToWorkForTask.MapKeys(){
		if taskValue[taskKeys.String()].FitToWork.Info.TaskFitToWorkName == fitToWorkName&& taskValue[taskKeys.String()].Settings.Status == helpers.StatusActive{
			return  true
			break
		}

	}

	return false
}