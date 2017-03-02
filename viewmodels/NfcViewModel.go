package viewmodels

type NfcViewModel struct {
	Values		[][]string
	Keys 		[]string
	CompanyTeamName string
	CompanyPlan	string

}

type EditNfcViewModel struct {
	PageType	string
	NfcId		string
	CustomerName	string
	Site      	string
	Location 	string
	NFCNumber	string
	CompanyTeamName string
	CompanyPlan	string

}