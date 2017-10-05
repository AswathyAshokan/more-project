package models



type Expirations struct {
	Info 		ExpirationInfo
	Settings 	ExpirationSettings
	Company         map[string]CompanyData

}

type CompanyData struct {
	CompanyName  		string
	CompanyStatus  		string
	NotificationShedules	map[string]ExpiryNotification

}
type ExpiryNotification struct {
	Category		string
	ExpiryId		string
	IsDeleted 		bool
	IsRead			bool
	IsViewed		bool
	LocalDate		string
	NotificationDate	int64

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
