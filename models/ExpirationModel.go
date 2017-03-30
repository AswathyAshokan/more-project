package models



type Expirations struct {
	Info 		ExpirationInfo
	Settings 	ExpirationSettings

}


type ExpirationInfo struct {
	AlertType 		string
	Description 		string
	DocumentId		string
	ExpirationDate		int64
	Mode  			string
	NotificationType	string
	Type 			string

}


type ExpirationSettings struct {
	DateOfCreation	int64
	Status 		string
}
