package models

type Users struct {
	Expiry			[]Expiry
	Info			UserInfo
	NextOfKin		NextOfKin
	Settings		UserSettings
	SocialNetworks		UserSocialNetworks
}

type Expiry struct {
	Alert			string
	DateOfExpiry		string
	Description		string
	DocumentLocation	string
	Mode			string
	Notify			string
	Type			string
}

type UserInfo struct {
	Address			string
	Email			string
	FullName		string
	Phone			string
	UserName		string
}

type NextOfKin struct {
	KinEmail		string
	KinName			string
	KinPhone		string
	Relation		string
}

type UserSettings struct {
	DateOfCreation		string
	Status			string
}

type UserSocialNetworks struct {
	FacebookId		string
	SkypeId			string
}