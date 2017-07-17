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
	if instructionSlice[0] !="" {
		for i := 0; i < len(instructionSlice); i++ {
			InstructionForFitToWork.Description = instructionSlice[i]
			InstructionForFitToWork.DateOfCreation = (time.Now().UnixNano() / 1000000)
			InstructionForFitToWork.Status = helpers.StatusActive
			id := betterguid.New()
			instructionMap[id] = InstructionForFitToWork
			err = db.Child("/FitToWork/" + companyTeamName + "/" + fitToWorkId + "/Instructions/").Set(instructionMap)
			if err != nil {
				log.Println(err)
				return false
			}
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
	log.Println("inside checking value",fitToWork)
	FitToWorkDetails := reflect.ValueOf(fitToWork)
	for _, fitKey:=range FitToWorkDetails.MapKeys() {
		log.Println("vddfgsf",companyTeamName)
		log.Println("hhhhh",fitKey.String())
		if fitKey.String() ==companyTeamName {
			log.Println("inside if")
			err = dB.Child("FitToWork/"+ companyTeamName).Value(&fullFitToWork)
			log.Println("inside fgfdgdg",fullFitToWork)

		}

	}
	dataValueOfFitToWork := reflect.ValueOf(fullFitToWork)
	for _, fitKeys:=range dataValueOfFitToWork.MapKeys(){
		log.Println("idddd",fitKeys.String())
		log.Println("gggg",fitWorkName)
		log.Println("mmmm",fullFitToWork[fitKeys.String()].FitToWorkName)
		if fullFitToWork[fitKeys.String()].FitToWorkName == fitWorkName{
			log.Println("gggggg")
			return true
			break
		}
		log.Println("dfsgdfgdggfsgfdgdgd",fitWork)
	}
	return false
}
func (m *FitToWork) IsfitToWorkUsedForTask( ctx context.Context, fitToWorkId string,companyTeamName string)(bool,map[string]FitToWork)  {
	fitDetail := map[string]FitToWork{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("FitToWork/"+companyTeamName+"/"+fitToWorkId).Value(&fitDetail)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false,fitDetail
	}
	log.Println(fitDetail)
	log.Println("job inside task",fitDetail)

	return true,fitDetail
}