package models


type SuperAdmins struct {
	Info 		SuperAdminData
	Settings 	SuperAdminSettings
}

type SuperAdminData struct {
	Email 		string
	FirstName	string
	LastName	string
	Password 	[]byte
	PhoneNo  	string


}
type SuperAdminSettings struct {
	DateOfCreation	int64
	Status 		string
}

/*
func(m *SuperAdmins) AddSuperAdminToDb(ctx context.Context) (bool){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("SuperAdmins").Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}
*/
