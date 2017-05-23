package models

import(
	"golang.org/x/net/context"
	"log"
	"time"
	"github.com/kjk/betterguid"
	"strings"
	"app/passporte/helpers"
	"reflect"
)
type ConsentReceipts struct {
	Info     	ConsentData
	Settings 	ConsentSettings
	Instructions	map[string]ConsentReceiptInstructions
	Members	 	map[string]ConsentMembers

}
type ConsentReceiptInstructions struct {
	Description    string
	//Status         string
	DateOfCreation int64

}

type ConsentMembers struct {
	MemberName	string
	Status          string
}
type ConsentData struct {
	ReceiptName  	string
	CompanyTeamName	string
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
			//InstructionForConsent.Status = helpers.StatusPending
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
	err = dB.Child("ConsentReceipts").OrderBy("Info/CompanyTeamName").EqualTo(CompanyTeamName).Value(&consentValue)
	if err != nil {
		log.Fatal(err)
		return false, consentValue
	}
	return true, consentValue
}

func GetSelectedUsersName(ctx context.Context,consentId string)(ConsentReceipts){
	consent :=ConsentReceipts{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/ConsentReceipts/"+consentId).Value(&consent)
	if err != nil {
		log.Fatal(err)
		return consent
	}
	return consent

}

func(m *ConsentReceipts) UpdateConsentDetails(ctx context.Context,consentId string,instructionSlice []string) (bool) {
	ConsentStatusDetails :=ConsentSettings{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}

	err = db.Child("/ConsentReceipts/"+consentId+"/Settings").Value(&ConsentStatusDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}

	m.Settings.DateOfCreation = ConsentStatusDetails.DateOfCreation
	m.Settings.Status = ConsentStatusDetails.Status
	err = db.Child("/ConsentReceipts/"+ consentId).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	instructionMap := make(map[string]ConsentReceiptInstructions)
	InstructionForConsent :=ConsentReceiptInstructions{}
	for i := 0; i < len(instructionSlice); i++ {
		InstructionForConsent.Description =instructionSlice[i]
		InstructionForConsent.DateOfCreation =time.Now().Unix()
		id := betterguid.New()
		instructionMap[id] = InstructionForConsent
		err = db.Child("/ConsentReceipts/"+consentId+"/Instructions/").Set(instructionMap)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	return true
}

func DeleteConsentReceiptById(ctx context.Context,consentId string)(bool)  {
	var keySlice []string
	ConsentStatusDetails :=ConsentSettings{}
	instructions := map[string]ConsentMembers{}
	updateInstructionStatus :=ConsentMembers{}
	updateConsentStatus := ConsentSettings{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/ConsentReceipts/"+consentId+"/Settings").Value(&ConsentStatusDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	updateConsentStatus.DateOfCreation = ConsentStatusDetails.DateOfCreation
	updateConsentStatus.Status = helpers.UserStatusDeleted
	err = db.Child("ConsentReceipts/"+consentId+"/Settings").Update(&updateConsentStatus)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	err = db.Child("/ConsentReceipts/"+consentId+"/Members").Value(&instructions)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	dataValue := reflect.ValueOf(instructions)

	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	log.Println("key",keySlice)
	for _,k := range keySlice {
		updateInstructionStatus.MemberName =instructions[k].MemberName
		updateInstructionStatus.Status = helpers.UserStatusDeleted
		err = db.Child("ConsentReceipts/"+consentId+"/Members/"+k).Update(&updateInstructionStatus)
		if err != nil {
			log.Fatal(err)
			return  false
		}
	}

	return true
}

