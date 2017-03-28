package models



type Expirations struct {
	Info 		ExpirationInfo
	Settings 	ExpirationSettings

}


type ExpirationInfo struct {
	AlertType 		string
	Description 		string
	DocumentId		string
	ExpirationDate		string
	Mode  			string
	NotificationType	string
	Type 			string

}


type ExpirationSettings struct {
	DateOfCreation	string
	Status 		string
}
