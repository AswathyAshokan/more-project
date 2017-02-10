/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
	"log"
	"time"
)

type NFC struct {
	CustomerName	string
	Site      	string
	Location 	string
	NFCNumber	string
}


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

func (m *NFC)GetNFCDetails(ctx context.Context)map[string]NFC{
	nfcDetail := map[string]NFC{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connection!")
	}
	err = dB.Child("NFCTag").Value(&nfcDetail)
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

func DeleteNFC(ctx  context.Context, key string)bool{
	log.Println("model:DeleteNFC()")
	log.Println(key)
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connection!")
	}
	err = dB.Child("/NFCTag/"+key).Remove()
	if err!=nil{
		log.Println("Removing Child Failed!")
		return false
	}else{
		log.Println("Child Removed Successfully!")
		return true
	}

}