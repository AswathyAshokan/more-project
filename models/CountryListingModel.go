package models


type Country struct {
	CountryName	string
	DialCode	string
	//States   	map[string]string
}

type CountryData struct {
	CountryName	string
	DialCode	string
	States   	map[string]string
}
