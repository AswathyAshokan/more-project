package models
import (
	"golang.org/x/net/context"
	"log"
	//"strings"
	"reflect"
	//
	//"app/passporte/helpers"
	"strconv"
	//"app/myFirstBeego/db"
	"app/passporte/helpers"
)
func(m *Users) GetAllUsersDetails(ctx context.Context,companyId string) (bool,[][]string,[]string,[][]string,[][]string){
	userDetails := map[string]Users{}
	IndiviualUserDetails :=map[string]UsersCompany{}
	var UserArrayDup [][]string
	var KeyDup []string
	var UserArrayForExpand [][]string
	var UserNextOfKin      [][]string

	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("Users").Value(&userDetails)
	userDetailsForProfile := reflect.ValueOf(userDetails)
	for _, UserKey := range userDetailsForProfile.MapKeys() {
		var UserArray []string
		var UserArrayExpandDup []string
		var NextOfKin       []string
		//err = db.Child("Users/"+UserKey.String()).Value(&IndiviualuserDetails)
		err = db.Child("Users/"+UserKey.String()+"/Company").Value(&IndiviualUserDetails)
		companyDetails := reflect.ValueOf(IndiviualUserDetails)
		for _, CompanyKey := range companyDetails.MapKeys() {
			if CompanyKey.String()==companyId && userDetails[UserKey.String()].Company[companyId].Status==helpers.StatusActive{
				UserArray=append(UserArray,"")
				UserArray=append(UserArray,userDetails[UserKey.String()].Info.FullName)
				UserArray=append(UserArray,userDetails[UserKey.String()].Info.Email)
				UserArray=append(UserArray,userDetails[UserKey.String()].Company[companyId].UserType)
				UserArrayDup =append(UserArrayDup,UserArray)
				UserArrayExpandDup=append(UserArrayExpandDup,userDetails[UserKey.String()].Info.Address)
				UserArrayExpandDup=append(UserArrayExpandDup,userDetails[UserKey.String()].Info.City)
				UserArrayExpandDup=append(UserArrayExpandDup,userDetails[UserKey.String()].Info.State)
				UserArrayExpandDup=append(UserArrayExpandDup,userDetails[UserKey.String()].Info.Country)
				UserArrayExpandDup=append(UserArrayExpandDup,userDetails[UserKey.String()].Info.ZipCode)
				phone := strconv.FormatInt(userDetails[UserKey.String()].Info.Phone, 10)
				UserArrayExpandDup=append(UserArrayExpandDup,phone)
				dateOfBirth := strconv.FormatInt(userDetails[UserKey.String()].Info.DateOfBirth, 10)
				UserArrayExpandDup=append(UserArrayExpandDup,dateOfBirth)
				UserArrayExpandDup=append(UserArrayExpandDup,userDetails[UserKey.String()].Settings.ThumbProfilePicture)
				UserArrayExpandDup=append(UserArrayExpandDup,UserKey.String())
				UserArrayForExpand =append(UserArrayForExpand,UserArrayExpandDup)
				NextOfKin=append(NextOfKin,userDetails[UserKey.String()].NextOfKin.KinEmail)
				NextOfKin=append(NextOfKin,userDetails[UserKey.String()].NextOfKin.KinName)

				NextOfKin=append(NextOfKin,userDetails[UserKey.String()].NextOfKin.KinPhone)
				NextOfKin=append(NextOfKin,userDetails[UserKey.String()].NextOfKin.Relation)
				NextOfKin=append(NextOfKin,UserKey.String())
				UserNextOfKin=append(UserNextOfKin,NextOfKin)
				KeyDup =append(KeyDup,UserKey.String())

			}
		}
	}
	return true,UserArrayDup,KeyDup,UserArrayForExpand,UserNextOfKin
}
