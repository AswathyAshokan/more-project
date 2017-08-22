package models

import(
	"golang.org/x/net/context"
	"log"
	"github.com/kjk/betterguid"
	"strings"
	"reflect"
	"app/passporte/helpers"

)
type ConsentReceipts struct {
	Info         ConsentData
	Settings     ConsentSettings
	Instructions ConsentInstructions


}
type ConsentInstructions struct {
	Description    string
	Users	 	map[string]ConsentMembers


}

type ConsentMembers struct {
	FullName  		string
	UserResponse    	string
}
type ConsentData struct {
	ReceiptName  	string
	CompanyName	string
}

type ConsentSettings struct {
	DateOfCreation	int64
	Status 		string
}


func(m *ConsentReceipts) AddConsentToDb(ctx context.Context,instructionSlice []string ,companyTeamName string,tempUserId []string) (bool){
	addConsentToUsers := ConsentReceiptDetails{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}


	consentValue,err := db.Child("ConsentReceipts/"+companyTeamName).Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	instructionMap := make(map[string]ConsentInstructions)
	InstructionForConsent := ConsentInstructions{}
	consentValueString := strings.Split(consentValue.String(),"/")
	consentUniqueID := consentValueString[len(consentValueString)-2]
	if instructionSlice[0] !=""{
		for i := 0; i < len(instructionSlice); i++ {
			InstructionForConsent.Description =instructionSlice[i]
			InstructionForConsent.Users= m.Instructions.Users
			id := betterguid.New()
			instructionMap[id] = InstructionForConsent

			err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentUniqueID+"/Instructions/").Set(instructionMap)
			if err != nil {
				log.Println(err)
				return false
			}
		}
	}
	var keySlice []string
	dataValue := reflect.ValueOf(m.Instructions.Users)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {
		addConsentToUsers.CompanyId = companyTeamName
		addConsentToUsers.UserResponse = m.Instructions.Users[k].UserResponse
		for i := 0; i < len(tempUserId); i++ {
			err = db.Child("/Users/"+tempUserId[i]+"/ConsentReceipts/"+consentUniqueID).Set(addConsentToUsers)
			if err != nil {
				log.Println(err)
				return false
			}
		}

	}
	return  true
}



func GetAllConsentReceiptDetails(ctx context.Context) (bool,map[string]ConsentReceipts){
	consentValue := map[string]ConsentReceipts{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("ConsentReceipts/").Value(&consentValue)
	if err != nil {
		log.Fatal(err)
		return false, consentValue
	}
	return true, consentValue
}

func GetDataOfConsentByConsentId(ctx context.Context,companyTeamName string)(map[string]ConsentReceipts){
	consent :=map[string]ConsentReceipts{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/ConsentReceipts/"+companyTeamName).Value(&consent)
	log.Println("consent",consent)
	if err != nil {
		log.Fatal(err)
	}
	return consent

}

func GetAllInstructionsById(ctx context.Context,companyTeamName string,consentId string)(map[string]ConsentInstructions)  {
	getInstructions:=map[string]ConsentInstructions{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions").Value(&getInstructions)
	if err != nil {
		log.Fatal(err)
	}
	return getInstructions


}

func(m *ConsentReceipts) UpdateConsentDetailsIfInstructionChanged(ctx context.Context,consentId string,instructionSlice []string,tempGroupId []string,tempGroupMembers []string,companyTeamName string ) (bool) {
	ConsentStatusDetails :=ConsentSettings{}
	addConsentToUsers := ConsentReceiptDetails{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId+"/Settings").Value(&ConsentStatusDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	m.Settings.DateOfCreation = ConsentStatusDetails.DateOfCreation
	m.Settings.Status = ConsentStatusDetails.Status
	err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	instructionMap := make(map[string]ConsentInstructions)
	InstructionForConsent := ConsentInstructions{}
	for i := 0; i < len(instructionSlice); i++ {
		InstructionForConsent.Description =instructionSlice[i]
		InstructionForConsent.Users= m.Instructions.Users
		id := betterguid.New()
		instructionMap[id] = InstructionForConsent
		err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions/").Set(instructionMap)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	addConsentToUsers.CompanyId = companyTeamName
	addConsentToUsers.UserResponse = helpers.StatusPending
	for i := 0; i < len(tempGroupId); i++ {
		err = db.Child("/Users/"+tempGroupId[i]+"/ConsentReceipts/"+consentId).Set(addConsentToUsers)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	return true
}

func DeleteConsentReceiptById(ctx context.Context,consentId string,companyTeamName string)(bool)  {
	//allUsers := map[string]Users{}
	ConsentStatusDetails :=ConsentSettings{}
	updateConsentStatus := ConsentSettings{}
	getInstructions:=map[string]ConsentInstructions{}
	//consentInUsers := map[string]ConsentReceiptDetails{}
	//updateConsentInUsers :=ConsentReceiptDetails{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId+"/Settings").Value(&ConsentStatusDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	updateConsentStatus.DateOfCreation = ConsentStatusDetails.DateOfCreation
	updateConsentStatus.Status = helpers.UserStatusDeleted
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Settings").Update(&updateConsentStatus)
	addConsentToUsers := ConsentReceiptDetails{}



	//delete consent from users
	 var userKeySlice []string
	var instructionKey []string
	addConsentToUsers.CompanyId = companyTeamName
	addConsentToUsers.UserResponse = helpers.UserStatusDeleted
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions").Value(&getInstructions)
	//err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions").Value(&consentRecepitDetails)
	dataValue := reflect.ValueOf(getInstructions)
	for _, key := range dataValue.MapKeys() {
		instructionKey = append(instructionKey, key.String())
	}
	log.Println("instruction key",instructionKey)
	for i:=0 ;i<1;i++{
		dataValue := reflect.ValueOf(getInstructions[instructionKey[i]].Users)
		for _, key := range dataValue.MapKeys() {
			userKeySlice = append(userKeySlice, key.String())
		}

	}
	log.Println("user Key ",userKeySlice)
	log.Println("consent id ",consentId)
	for i := 0; i < len(userKeySlice); i++ {
		err = db.Child("/Users/" + userKeySlice[i] + "/ConsentReceipts/" + consentId).Set(addConsentToUsers)
	}
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true
}



func IsInstructionEdited(ctx context.Context,instructionSlice []string,consentId string,companyTeamName string)(bool)  {
	log.Println("cp5")
	count :=0
	instructions :=map[string]ConsentInstructions{}
	var AllInstructions []string
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions").Value(&instructions)
	if err != nil {
		log.Println("cpp1")
		log.Fatal(err)
	}
	log.Println("instructions",instructions)
	dataValue := reflect.ValueOf(instructions)

	for _, key := range dataValue.MapKeys() {
		AllInstructions = append(AllInstructions,instructions[key.String()].Description)
		log.Println("instructions[key.String()].Description",instructions[key.String()].Description)
	}

	for i:=0;i<len(AllInstructions);i++{
		for _, v := range AllInstructions {
			if v == instructionSlice[i] {
				count = count+1
			}
		}
		if count == len(AllInstructions){
			return true
		} else {
			return false
		}
		log.Println("time",count)
	}

	return false
}


func GetEachConsentByCompanyId(ctx context.Context,consentId string,companyTeamName string)(ConsentReceipts){
	log.Println("cp2")
	consent :=ConsentReceipts{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId).Value(&consent)
	if err != nil {
		log.Fatal(err)
	}
	return consent

}

func GetAllInstructionsFromConsent(ctx context.Context,consentId string,companyTeamName string)(map[string]ConsentInstructions){
	log.Println("cp3")
	instructionOfConsent :=map[string]ConsentInstructions{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions").Value(&instructionOfConsent)
	if err != nil {
		log.Fatal(err)
	}
	return instructionOfConsent

}


func GetAllUsersFromInstructions(ctx context.Context,consentId string,companyTeamName string)(ConsentMembers){
	usersOfConsent :=ConsentMembers{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions"+"/Users").Value(&usersOfConsent)
	if err != nil {
		log.Fatal(err)
	}
	return usersOfConsent

}

func(m *ConsentReceipts) UpdateConsentDataIfInstructionNotChanged(ctx context.Context,consentId string,instructionSlice []string,tempGroupId []string,tempGroupMembers []string,companyTeamName string ) (bool) {
	ConsentStatusDetails :=ConsentSettings{}
	addConsentToUsers := ConsentReceiptDetails{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId+"/Settings").Value(&ConsentStatusDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	log.Println("ConsentStatusDetails",ConsentStatusDetails)
	m.Settings.DateOfCreation = ConsentStatusDetails.DateOfCreation
	m.Settings.Status = ConsentStatusDetails.Status
	err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	instructionMap := make(map[string]ConsentInstructions)
	InstructionForConsent := ConsentInstructions{}
	for i := 0; i < len(instructionSlice); i++ {
		InstructionForConsent.Description =instructionSlice[i]
		InstructionForConsent.Users= m.Instructions.Users
		id := betterguid.New()
		instructionMap[id] = InstructionForConsent
		err = db.Child("/ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions/").Set(instructionMap)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	var keySlice []string
	dataValue := reflect.ValueOf(m.Instructions.Users)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {
		addConsentToUsers.CompanyId = companyTeamName
		addConsentToUsers.UserResponse = m.Instructions.Users[k].UserResponse
		for i := 0; i < len(tempGroupId); i++ {
			err = db.Child("/Users/"+tempGroupId[i]+"/ConsentReceipts/"+consentId).Set(addConsentToUsers)
			if err != nil {
				log.Println(err)
				return false
			}
		}
	}
	return true
}


func GetSelectedUsersName(ctx context.Context,consentId string)(map[string]ConsentReceipts){
	consent :=map[string]ConsentReceipts{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+consentId).Value(&consent)
	if err != nil {
		log.Fatal(err)
	}
	return consent

}


