package controllers

import (
	"reflect"
	"app/passporte/helpers"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
)

type CustomerManagementController struct {
	BaseController
}

func (c *CustomerManagementController) CustomerManagement() {
	customerManagementViewModel := viewmodels.Customer{}
	registeredCompany := models.Company{}
	dbStatus,allCompanyData:= registeredCompany.GetAllRegisteredCompanyDetails(c.AppEngineCtx)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allCompanyData)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		log.Println("key",keySlice)
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, allCompanyData[k].Info.CompanyName)
			tempValueSlice = append(tempValueSlice, allCompanyData[k].Info.Address)
			tempValueSlice = append(tempValueSlice, allCompanyData[k].Info.ZipCode)
			/*tempValueSlice = append(tempValueSlice, allCompanyData[k].Settings.DateOfCreation)*/
			/*log.Println("temp:",allCompanyData[k].Admins)
			dataValue := reflect.ValueOf(allCompanyData[k].Admins)
			var adminKey []string
			for _, key := range dataValue.MapKeys() {
				adminKey = append(adminKey, key.String())
			}
			var tempAdminDetais []string
			for _,i := range adminKey{

				tempAdminDetais = append(tempAdminDetais,allCompanyData[k].Admins[i].FirstName)

			}
			log.Println("tempkey:",tempAdminDetais)

*/
			log.Println("temp:",tempValueSlice)
			customerManagementViewModel.Values = append(customerManagementViewModel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]

		}

		customerManagementViewModel.Keys = keySlice
		c.Data["vm"] = customerManagementViewModel
		c.TplName = "template/customer-management.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

