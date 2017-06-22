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
	Info     	ConsentData
	Settings 	ConsentSettings
	Instructions	ConsentReceiptInstructions


}
type ConsentReceiptInstructions struct {
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
	instructionMap := make(map[string]ConsentReceiptInstructions)
	InstructionForConsent :=ConsentReceiptInstructions{}
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
		log.Println("from model values",addConsentToUsers)
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
	err = dB.Child("ConsentReceipts").Value(&consentValue)
	log.Println("hehehe",consentValue)
	if err != nil {
		log.Fatal(err)
		return false, consentValue
	}
	return true, consentValue
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

func GetAllInstructionsById(ctx context.Context,consentId string,instructionId string)(map[string]ConsentReceiptInstructions)  {
	log.Println("get key",consentId)
	getInstructions:=map[string]ConsentReceiptInstructions{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+consentId+"/"+instructionId+"/Instructions").Value(&getInstructions)
	log.Println("get instruction",getInstructions)
	if err != nil {
		log.Fatal(err)
	}
	return getInstructions


}

func GetAllUsersNameAndStatus(ctx context.Context,companyTeamName string,consentId string,instructionId string)(map[string]ConsentMembers){
	consent :=map[string]ConsentMembers{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions"+"/"+instructionId+"/Users").Value(&consent)
	log.Println("users",instructionId)
	if err != nil {
		log.Fatal(err)
	}
	return consent

}

func GetEachUserDetailsById(ctx context.Context,companyTeamName string,consentId string,instructionId string,specificId string) (ConsentMembers) {
	consentUsers :=ConsentMembers{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Instructions"+"/"+instructionId+"/Users/"+specificId).Value(&consentUsers)
	log.Println("users",specificId)
	if err != nil {
		log.Fatal(err)
	}
	return consentUsers
}

func GetMemberStatus(ctx context.Context,consentId string)( map[string]ConsentMembers){
	members := map[string]ConsentMembers{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/ConsentReceipts/"+consentId+"/Members").Value(&members)
	if err != nil {
		log.Fatal(err)
	}
	return members

	/*dataValue := reflect.ValueOf(members)

	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _,k := range keySlice {
		*//*err = db.Child("/Users/"+k+"/Info").Value(&userName)
		log.Println("username",userName)
		if err != nil {
			log.Fatal(err)
		}*//*
		updateMemberStatus.Status = members[k].Status
		updateMemberStatus.MemberName = members[k].MemberName
	}
	return updateMemberStatus*/
}



func(m *ConsentReceipts) UpdateConsentDetails(ctx context.Context,consentId string,instructionSlice []string,tempGroupId []string,tempGroupMembers []string) (bool) {
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

func DeleteConsentReceiptById(ctx context.Context,consentId string,companyTeamName string)(bool)  {
	allUsers := map[string]Users{}
	ConsentStatusDetails :=ConsentSettings{}
	updateConsentStatus := ConsentSettings{}
	consentInUsers := map[string]ConsentReceiptDetails{}
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
	/*err = db.Child("ConsentReceipts/"+companyTeamName+"/"+consentId+"/Settings").Update(&updateConsentStatus)
	if err != nil {
		log.Fatal(err)
		return  false
	}*/
	err = db.Child("Users").Value(&allUsers)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	dataValue := reflect.ValueOf(allUsers)

	for _, key := range dataValue.MapKeys() {
		log.Println("key",key.String())
		err = db.Child("Users/"+key.String()+"/ConsentReceipts").Value(&consentInUsers)

		//consentDataValue := reflect.ValueOf(consentInUsers)
		/*for _, k := range consentDataValue.MapKeys() {

			*//*tempUserResponse := consentInUsers[k.String()].UserResponse
			tempCompanyId :=consentInUsers[k.String()].*//*
		}
*/

	}
	log.Println("user consent",consentInUsers)
	/*for _, key := range keySlice{

	}*/
	return true
}



func IsInstructionEdited(ctx context.Context,instructionSlice []string,consentId string)(bool)  {
	count :=0
	instructions :=map[string]ConsentReceiptInstructions{}
	var keySlice []string
	var AllInstructions []string
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
		return false
	}
	err = db.Child("/ConsentReceipts/"+consentId+"/Instructions").Value(&instructions)
	if err != nil {
		log.Fatal(err)
	}
	dataValue := reflect.ValueOf(instructions)

	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _,k := range keySlice {
		 AllInstructions = append(AllInstructions,instructions[k].Description)
	}
	for i:=0;i<len(AllInstructions);i++{
		if AllInstructions[i] == instructionSlice[i]{
			log.Println("true from model")
			count = count+1
		}
	}
	if count == len(AllInstructions){
		return true
	} else {
		return false
	}

	//return false
}


