
/* Author :Aswathy Ashok */

package controllers

import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
	"time"
	"fmt"
	"reflect"
	"app/passporte/helpers"
)

type ContactUserController struct {
	BaseController
}

/* Add contact detail to DB*/
func (c *ContactUserController)AddNewContact() {
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	viewModel := viewmodels.ContactUserViewModel{}
	if r.Method == "POST" {
		user:=models.ContactUser{}
		user.Info.Name= c.GetString("name")
		user.Info.State = c.GetString("state")
		user.Info.ZipCode = c.GetString("zipcode")
		user.Info.Email = c.GetString("emailAddress")
		user.Info.PhoneNumber= c.GetString("phoneNumber")
		user.Info.Address = c.GetString("address")
		user.Info.Country = c.GetString("country")
		tempCustomerName := c.GetStrings("customerName")
		tempCustomerId := c.GetStrings("customerId")
		user.Settings.DateOfCreation =time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println(reflect.TypeOf(user.Settings.DateOfCreation))
		user.Settings.Status = helpers.StatusActive
		user.Info.CompanyTeamName = storedSession.CompanyTeamName
		customerMap := make(map[string]models.CustomerDetails)
		customerDetail :=models.CustomerDetails{}
		for i := 0; i < len(tempCustomerId); i++ {
			customerDetail.CustomerName = tempCustomerName[i]
			customerDetail.Status =helpers.StatusActive
			customerMap[tempCustomerId[i]] = customerDetail
		}
		user.Customer = customerMap
		dbStatus := user.AddContactToDB(c.AppEngineCtx)
		switch dbStatus {
		case true:

			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}else {

		customers ,dbStatus:=models.GetAllCustomerDetails(c.AppEngineCtx,companyTeamName)
		var keySlice []string
		dataValue := reflect.ValueOf(customers)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(customers)
			for _, k := range dataValue.MapKeys() {
				if customers[k.String()].Settings.Status =="Active"{
					viewModel.CustomerNameArray  = append(viewModel.CustomerNameArray, customers[k.String()].Info.CustomerName)
					viewModel.CustomerKeys=append(viewModel.CustomerKeys, k.String())
				}

			}
			log.Println("customer name array",viewModel.CustomerNameArray)
		case false:
			log.Println(helpers.ServerConnectionError)
		}

		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.CompanyPlan = storedSession.CompanyPlan
		viewModel.AdminLastName =storedSession.AdminLastName
		viewModel.AdminFirstName =storedSession.AdminFirstName
		viewModel.ProfilePicture =storedSession.ProfilePicture
		viewModel.PageType = helpers.SelectPageForAdd
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-contacts.html"
	}


}

/*Display all contact detail*/
func (c *ContactUserController)DisplayContactDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	log.Println("company...",companyTeamName)
	storedSession := ReadSession(w, r, companyTeamName)
	dbStatus, contact := models.GetAllContact(c.AppEngineCtx,companyTeamName)
	log.Println(contact)
	viewModel := viewmodels.ContactUserViewModel{}

	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(contact)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			if contact[k].Settings.Status == helpers.StatusActive {
				tempValueSlice = append(tempValueSlice, contact[k].Info.Name)
				tempValueSlice = append(tempValueSlice, contact[k].Info.Address)
				tempValueSlice = append(tempValueSlice,contact[k].Info.Country)
				tempValueSlice = append(tempValueSlice, contact[k].Info.State)
				tempValueSlice = append(tempValueSlice, contact[k].Info.ZipCode)
				tempValueSlice = append(tempValueSlice, contact[k].Info.Email)
				tempValueSlice = append(tempValueSlice, contact[k].Info.PhoneNumber)
				viewModel.Values = append(viewModel.Values, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}
		}
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.AdminFirstName =storedSession.AdminFirstName
		viewModel.AdminLastName =storedSession.AdminLastName
		viewModel.CompanyPlan = storedSession.CompanyPlan
		viewModel.ProfilePicture =storedSession.ProfilePicture
		log.Println("dhdghgdfh",viewModel.ProfilePicture)
		viewModel.Keys = keySlice
		viewModel.PageType=helpers.SelectPageForAdd
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/contacts-details.html"

	case false:

		log.Println(helpers.ServerConnectionError)
	}
}

/*Function for delete contact from DB*/
func (c *ContactUserController)LoadDeleteContact() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	contactId :=c.Ctx.Input.Param(":contactId")
	user := models.TasksContact{}
	dbStatus,contactDetail := user.IsContactUsedForTask(c.AppEngineCtx, contactId)
	log.Println("status",dbStatus)
	log.Println(contactDetail)
	var condition string
	switch dbStatus {
	case true:
		log.Println("true")
		if len(contactDetail) !=0{
			dataValue := reflect.ValueOf(contactDetail)
			for _, key := range dataValue.MapKeys() {
				if contactDetail[key.String()].TaskContactStatus ==helpers.StatusActive{
					log.Println("insideeee fgjgfjh")
					condition ="true"

					break
				}else{
					log.Println("false")

				}
			}
			if condition == "true"{

				w.Write([]byte("true"))
			}else {
				w.Write([]byte("false"))
			}
		}else{
			w.Write([]byte("false"))
		}


	case false :
		w.Write([]byte("false"))
	}



}
/*To perform edit operation*/
func (c *ContactUserController)LoadEditContact() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		contactId := c.Ctx.Input.Param(":contactId")
		user:=models.ContactUser{}
		user.Info.Name= c.GetString("name")
		user.Info.State = c.GetString("state")
		user.Info.ZipCode = c.GetString("zipcode")
		user.Info.Email = c.GetString("emailAddress")
		user.Info.PhoneNumber= c.GetString("phoneNumber")
		user.Info.Country = c.GetString("country")
		user.Info.Address = c.GetString("address")
		tempCustomerName := c.GetStrings("customerName")
		tempCustomerId := c.GetStrings("customerId")
		customerMap := make(map[string]models.CustomerDetails)
		customerDetail :=models.CustomerDetails{}
		for i := 0; i < len(tempCustomerId); i++ {
			customerDetail.CustomerName = tempCustomerName[i]
			customerDetail.Status =helpers.StatusActive
			customerMap[tempCustomerId[i]] = customerDetail
		}
		user.Customer = customerMap
		user.Settings.DateOfCreation =time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println(reflect.TypeOf(user.Settings.DateOfCreation))
		user.Settings.Status = helpers.StatusActive
		user.Info.CompanyTeamName = storedSession.CompanyTeamName
		dbStatus := user.UpdateContactToDB(c.AppEngineCtx,contactId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		contact :=models.ContactUser{}

		contactId := c.Ctx.Input.Param(":contactId")
		viewModel := viewmodels.ContactUserViewModel{}
		customers ,dbStatusCustomer:=models.GetAllCustomerDetails(c.AppEngineCtx,companyTeamName)
		var keySlice []string
		var activeCustomer []string
		dataValue := reflect.ValueOf(customers)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		switch dbStatusCustomer {
		case true:
			dataValue := reflect.ValueOf(customers)
			for _, k := range dataValue.MapKeys() {
				if customers[k.String()].Settings.Status =="Active"{
					viewModel.CustomerNameArray  = append(viewModel.CustomerNameArray, customers[k.String()].Info.CustomerName)
					viewModel.CustomerKeys=append(viewModel.CustomerKeys, k.String())
					activeCustomer =append(activeCustomer, k.String())
				}

			}
			log.Println("customer name array",viewModel.CustomerNameArray)
		case false:
			log.Println(helpers.ServerConnectionError)
		}

		dbStatus,contact := contact.RetrieveContactIdFromDB(c.AppEngineCtx, contactId)
		switch dbStatus {
		case true:
			viewModel.PageType = helpers.SelectPageForEdit
			dataValue := reflect.ValueOf(contact.Customer)
			for _, customerKey := range dataValue.MapKeys() {
				for i:=0;i<len(activeCustomer);i++{
					if activeCustomer[i] ==customerKey.String(){
						viewModel.EditCustomerKey = append(viewModel.EditCustomerKey, customerKey.String())

					}
				}

			}
			log.Println("page type",viewModel.PageType);
			viewModel.Name=contact.Info.Name
			viewModel.Address =contact.Info.Address
			viewModel.State =contact.Info.State
			viewModel.ZipCode =contact.Info.ZipCode
			viewModel.Email =contact.Info.Email
			viewModel.PhoneNumber =contact.Info.PhoneNumber
			viewModel.Country = contact.Info.Country
			viewModel.ContactId=contactId
			viewModel.CompanyTeamName = storedSession.CompanyTeamName
			viewModel.CompanyPlan = storedSession.CompanyPlan
			viewModel.AdminFirstName =storedSession.AdminFirstName
			viewModel.AdminLastName = storedSession.AdminLastName
			viewModel.ProfilePicture =storedSession.ProfilePicture
			c.Data["vm"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-contacts.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}
//deletion function
func (c *ContactUserController) DeleteContactIfNotInTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	contactId := c.Ctx.Input.Param(":contactId")
	user :=models.ContactUser{}
	log.Println("inside deletion of cotact")
	contact :=models.TasksContact{}
	var TaskSlice []string
	dbStatus,contactDetails := contact.IsContactUsedForTask(c.AppEngineCtx, contactId)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(contactDetails)
		for _, key := range dataValue.MapKeys() {
			TaskSlice = append(TaskSlice, key.String())
		}
		dbStatus := user.DeleteContactFromDB(c.AppEngineCtx, contactId,TaskSlice)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false :
			w.Write([]byte("false"))
		}
	}
}



func (c *ContactUserController) RemoveContactFromTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	contactId := c.Ctx.Input.Param(":contactId")
	log.Println("hiiii")
	user :=models.ContactUser{}
	dbStatus := user.DeleteContactFromDBForNonTask(c.AppEngineCtx, contactId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}
	}
func (c *ContactUserController)CheckPhoneNumber(){
	w := c.Ctx.ResponseWriter
	phoneNumber := c.GetString("phoneNumber")
	log.Println("phone number",phoneNumber)
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	isPhoneNumberUsed := models.CheckPhoneNumberIsUsed(c.AppEngineCtx, phoneNumber,companyTeamName)
	switch isPhoneNumberUsed{
	case true:
		w.Write([]byte("false"))
	case false:
		w.Write([]byte("true"))
	}
}
func (c *ContactUserController)CheckEmailAddress(){
	w := c.Ctx.ResponseWriter
	emailAddress := c.GetString("emailAddress")
	log.Println("email",emailAddress)
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	isEmailAddressUsed := models.CheckEmailAddressIsUsed(c.AppEngineCtx, emailAddress,companyTeamName)
	switch isEmailAddressUsed{
	case true:
		w.Write([]byte("false"))
	case false:
		w.Write([]byte("true"))
	}
}