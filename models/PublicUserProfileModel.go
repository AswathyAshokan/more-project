package models

type PublicUserProfile struct {
	Email 			string
	FullName		string
	ThumbProfilePicture	string
	Company          	map[string]CompanyDataForProfile
}

type CompanyDataForProfile struct {

	UserType	string
	Status 		string
}
