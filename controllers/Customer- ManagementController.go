package controllers

import (
	"reflect"
	"app/passporte/helpers"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
	"strconv"
)

type CustomerManagementController struct {
	BaseController
}

func (c *CustomerManagementController) CustomerManagement() {
	customerManagementViewModel := viewmodels.CustomerManagement{}
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

		var adminKeyFromCompany []string
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, allCompanyData[k].Info.CompanyName)
			tempValueSlice = append(tempValueSlice, allCompanyData[k].Info.Address)
			dataValue := reflect.ValueOf(allCompanyData[k].Admins)

			for _, key := range dataValue.MapKeys() {
				adminKeyFromCompany = append(adminKeyFromCompany, key.String())
			}
			adminStatus,adminDetails := models.GetAdminDetailsById(c.AppEngineCtx, adminKeyFromCompany)
				switch adminStatus {
				case true:
					tempValueSlice = append(tempValueSlice,adminDetails.Info.FirstName)
					tempValueSlice = append(tempValueSlice,adminDetails.Info.Email)
					tempValueSlice = append(tempValueSlice,adminDetails.Info.PhoneNo)
				case false:
					log.Println("false")

				}
			tempValueSlice = append(tempValueSlice, strconv.FormatInt(allCompanyData[k].Settings.DateOfCreation,10))
			tempValueSlice = append(tempValueSlice,allCompanyData[k].Plan)
			log.Println("temo",tempValueSlice)
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


/*To delete selected record from database*/

/*
func (c *ContactUserController)LoadDeleteCustomerManagement() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	user := models.ContactUser{}
	dbStatus := user.DeleteContactFromDB(c.AppEngineCtx)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}


}
*/
