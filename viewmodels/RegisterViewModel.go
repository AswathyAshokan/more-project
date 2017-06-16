package viewmodels
type EditProfileViewModel struct {
	FirstName	string
	LastName	string
	PhoneNo		string
	Email		string
	CompanyTeamName	string
	CompanyPlan	string
	AdminFirstName	string
	AdminLastName	string
	ProfilePicture	string
}

type ForgotPassword struct {
	VerificationKey string
}

type DisplayCountryDetails struct {
	CountryName     []string
	CountryAllData	[][]string
	DialCode    	[]string
	Key 		[]string

}
