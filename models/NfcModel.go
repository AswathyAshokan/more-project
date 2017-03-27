/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
	"log"
	"time"
	"app/passporte/helpers"
)
//Structs for Adding NFC Tag
type NFCInfo struct{
	CustomerName	string
	Site      	string
	Location 	string
	NFCNumber	string
	CompanyTeamName	string
}

type NFCSettings struct{
	Status string
	DateOfCreation int64
}

type NFC struct {
	Info NFCInfo
	Settings NFCSettings
}

//Add new NFC Tag
func (m *NFC)AddNFC(ctx context.Context)bool{
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No Db Connection!")

	}

	_,err = dB.Child("NFCTag").Push(m)
	if err!=nil{
		log.Println("NFC Tag Insertion failed!")
		return false
	}else{
		log.Println("NFC Tag Inserted Successfully!")
		return true
	}

}

//Get existing NFC Tag Details
func (m *NFC)GetAllNFCDetails(ctx context.Context, companyTeamName string)map[string]NFC{
	nfcDetail := map[string]NFC{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connection!")
	}
	nfcStatus :="Active"
	err = dB.Child("NFCTag").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).OrderBy("Settings/Status").EqualTo(nfcStatus).Value(&nfcDetail)
	if err!=nil{
		log.Println("Retrieving value failed!")
	}
	log.Println("NFC Details are:",nfcDetail)
	log.Println("Time:",int32(time.Now().Unix()))
	return nfcDetail
}

func (m *NFC)GetNFCDetailsById(ctx context.Context, nfcId string)(bool, NFC){
	log.Println("GetNFCDetailsById()")
	log.Println("NFC ID:",nfcId)
	nfcDetails := NFC{}
	dB, err := GetFirebaseClient(ctx,"")
	if err !=nil{
		log.Println("No Db Connection!")
	}
	err = dB.Child("/NFCTag/"+nfcId).Value(&nfcDetails)
	if err!=nil{
		log.Println("Retrieving value failed!")
		return false, nfcDetails
	}else{
		log.Println("Retrieving value success!")
		log.Println(nfcDetails)
		return true, nfcDetails
	}

}

//Update Existing NFC Tag Details
func (m *NFC)UpdateNFCDetails(ctx context.Context, nfcId string)bool{
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connection!")
	}
	err = dB.Child("/NFCTag/"+nfcId+"/Info").Update(&m.Info)
	if err!=nil{
		log.Println("Update failed!",err)
		return false
	}else{
		log.Println("Update Success!")
		return true
	}

}

//Delete existing NFC Tag
func DeleteNFC(ctx  context.Context, key string)bool{
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connection!")
	}
	nfcDetail :=NFC{}
	nfcUpdate :=NFCSettings{}
	err = dB.Child("/NFCTag/"+key).Value(&nfcDetail)
	nfcUpdate.DateOfCreation =nfcDetail.Settings.DateOfCreation
	nfcUpdate.Status =helpers.StatusInActive
	err = dB.Child("/NFCTag/"+key+"/Settings").Update(&nfcUpdate)
	if err!=nil{
		log.Println("Removing Child Failed!")
		return false
	}else{
		log.Println("Child Removed Successfully!")
		return true
	}

}