package models
import (
	"golang.org/x/net/context"
	"log"

)
type CompanyFileUpload struct {
	Info     	DocumentInfo
	Settings 	DocumentSettings
}
type DocumentInfo struct {
	FolderName	string
	FileName	string
	DocumentUrl	string

}
type DocumentSettings struct {
	DateOfCreation		int64
	Status			string

}

func (m *CompanyFileUpload)AddCompanyDocument(ctx context.Context,companyId string) (bool) {
	log.Println("values in m:",m)
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	if (len(m.Info.DocumentUrl) !=0 &&len(m.Info.FolderName)!=0){
		_, err = dB.Child("CompanyDocument/"+companyId).Push(m)
	}


	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}


func GetAllCompanyDocument(ctx context.Context,companyId string) (bool,map[string]CompanyFileUpload) {
	dB, err := GetFirebaseClient(ctx,"")
	CompanyFileUpload :=map[string]CompanyFileUpload{}
	if err!=nil{
		log.Println("Connection error:",err)
	}
	 err = dB.Child("CompanyDocument/"+companyId).Value(&CompanyFileUpload)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false,CompanyFileUpload
	}
	return true,CompanyFileUpload
}



func GetCompanyDocumentById(ctx context.Context,companyId string,documentId string) (bool,CompanyFileUpload) {
	dB, err := GetFirebaseClient(ctx,"")
	CompanyFileUpload :=CompanyFileUpload{}
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("CompanyDocument/"+companyId+"/"+documentId).Value(&CompanyFileUpload)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false,CompanyFileUpload
	}
	return true,CompanyFileUpload
}



func (m *CompanyFileUpload)EditCompanyIdWithoutChange(ctx context.Context,companyId string,documentId string) (bool) {
	log.Println("values in m:",m)
	CompanyFileUpload :=CompanyFileUpload{}
	//UpdatedFileUpload :=CompanyFileUpload{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("CompanyDocument/"+companyId+"/"+documentId).Value(&CompanyFileUpload)
	if (m.Info.DocumentUrl ==""){
		m.Info.DocumentUrl=CompanyFileUpload.Info.DocumentUrl
	}else{
		m.Info.DocumentUrl=m.Info.DocumentUrl
	}

	m.Info.FileName =m.Info.FileName
	m.Info.FolderName =m.Info.FolderName
	m.Settings.DateOfCreation =CompanyFileUpload.Settings.DateOfCreation
	m.Settings.Status =CompanyFileUpload.Settings.Status
	err = dB.Child("CompanyDocument/"+companyId+"/"+documentId).Set(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}


func DeleteCompanyDocument(ctx context.Context,companyId string,documentId string) (bool) {
	log.Println("document id",documentId)
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("CompanyDocument/"+companyId+"/"+documentId).Remove()
	if err!=nil{
		log.Println("Deletion error error:",err)
		return false
	}
	return true
}