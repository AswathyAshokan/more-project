package models

import(
	"golang.org/x/net/context"
	"log"
	"time"
	"github.com/kjk/betterguid"
	"strings"
	"app/passporte/helpers"
)
type ConsentReceipts struct {
	Info     	ConsentData
	Settings 	ConsentSettings
	Instructions	map[string]ConsentReceiptInstructions
	Members	 	map[string]ConsentMembers

}
type ConsentReceiptInstructions struct {
	Description    string
	Status         string
	DateOfCreation int64

}

type ConsentMembers struct {
	MemberName	string
}
type ConsentData struct {
	ReceiptName  	string
}

type ConsentSettings struct {
	DateOfCreation	int64
	Status 		string
}


func(m *ConsentReceipts) AddConsentToDb(ctx context.Context,instructionSlice []string) (bool){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}


	consentValue,err := db.Child("ConsentReceipts").Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	instructionMap := make(map[string]ConsentReceiptInstructions)
	InstructionForConsent :=ConsentReceiptInstructions{}
	consentValueString := strings.Split(consentValue.String(),"/")
	consentUniqueID := consentValueString[len(consentValueString)-2]
	if instructionSlice[0] !=""{

		for i := 0; i < len(instructionSlice); i++ {

			InstructionForConsent.Description =instructionSlice[i]
			InstructionForConsent.DateOfCreation =time.Now().Unix()
			InstructionForConsent.Status = helpers.StatusActive
			id := betterguid.New()
			instructionMap[id] = InstructionForConsent
			err = db.Child("/ConsentReceipts/"+consentUniqueID+"/Instructions/").Set(instructionMap)
			if err != nil {
				log.Println(err)
				return false
			}

		}
	}
	return  true
}


func GetAllConsentReceiptDetails(ctx context.Context,CompanyTeamName string) (bool,map[string]ConsentReceipts){
	consentValue := map[string]ConsentReceipts{}
	dB, err := GetFirebaseClient(ctx,"")
	//taskStatus := "Active"
	err = dB.Child("Tasks").OrderBy("Info/CompanyTeamName").EqualTo(CompanyTeamName).Value(&consentValue)
	if err != nil {
		log.Fatal(err)
		return false, consentValue
	}
	return true, consentValue
}


