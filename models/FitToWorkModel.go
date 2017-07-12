package models
import(
	"golang.org/x/net/context"
	"log"
	"github.com/kjk/betterguid"
	"strings"
	"time"
	"app/passporte/helpers"
)


type FitToWork struct {
	FitToWorkName string
	FitToWork	map[string]TaskFitToWorks
	Status		string
}
type TaskFitToWorks struct {
	Description    string
	Status         string
	DateOfCreation int64

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
	m.Status = fitToWorkDetails.Status
	err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+fitToWorkId).Update(&m)

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
	log.Println("cp2")
	fitWork :=FitToWork{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("FitToWork/"+companyTeamName+"/"+ fitToWorkId).Value(&fitWork)
	if err != nil {
		log.Fatal(err)
	}
	return fitWork

}