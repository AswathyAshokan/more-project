package models

import(
	"golang.org/x/net/context"
	"log"
)
type ConsentReceipts struct {
	Info     ConsentData
	Settings ConsentSettings
	Members	 	map[string]ConsentMembers

}

type ConsentMembers struct {
	MemberName	string
}
type ConsentData struct {
	ReceiptName  	string
	Instructions 	string
}

type ConsentSettings struct {
	DateOfCreation	int64
	Status 		string
}


func(m *ConsentReceipts) AddConsentToDb(ctx context.Context) (bool){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("ConsentReceipts").Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}

