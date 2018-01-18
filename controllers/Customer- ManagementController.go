
//created by farsana
//for
package controllers

import (
	"reflect"
	"app/passporte/helpers"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
	"time"
)

type CustomerManagementController struct {
	BaseController
}

func (c *CustomerManagementController) CustomerManagement() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter

	storedSession := ReadSessionForSuperAdmin(w,r)
	log.Println("read session:",storedSession)
	customerManagementViewModel := viewmodels.CustomerManagement{}
	dbStatus,allCompanyData:= models.GetAllRegisteredCompanyDetails(c.AppEngineCtx)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allCompanyData)
		var keySlice []string
		var adminKeyFromCompany []string
		for _, k := range dataValue.MapKeys() {
			var tempValueSlice []string


			if allCompanyData[k.String()].Settings.Status == helpers.StatusActive{
				keySlice = append(keySlice, k.String())
				tempValueSlice = append(tempValueSlice, allCompanyData[k.String()].Info.CompanyName)
				//tempValueSlice = append(tempValueSlice, allCompanyData[k].Info.Address)
				dataValue := reflect.ValueOf(allCompanyData[k.String()].Admins)

				for _, key := range dataValue.MapKeys() {
					adminKeyFromCompany = append(adminKeyFromCompany, key.String())
				}
				adminStatus,adminDetails := models.GetAdminDetailsById(c.AppEngineCtx, adminKeyFromCompany)
				switch adminStatus {
				case true:
					tempValueSlice = append(tempValueSlice,adminDetails.Info.FirstName)
					tempValueSlice = append(tempValueSlice,adminDetails.Info.Email)
					log.Println("phone number",adminDetails.Info.PhoneNo)
					tempValueSlice = append(tempValueSlice,adminDetails.Info.PhoneNo)
				case false:
					log.Println(helpers.ServerConnectionError)
				}
				//tempTym :=strconv.FormatInt(allCompanyData[k].Settings.DateOfCreation,10)
				//i, _ := strconv.ParseInt(tempTym, 10, 64)
				//log.Println("hhh",i)
				startDate := time.Unix(allCompanyData[k.String()].Settings.DateOfCreation, 0).Format("2006/01/02")
				tempValueSlice = append(tempValueSlice, startDate)
				tempValueSlice = append(tempValueSlice,allCompanyData[k.String()].Plan)
				customerManagementViewModel.Values = append(customerManagementViewModel.Values,tempValueSlice)
				tempValueSlice = tempValueSlice[:0]

			}
			customerManagementViewModel.Keys = keySlice
			c.Data["vm"] = customerManagementViewModel
			c.Layout = "layout/layout-superadmin.html"
			c.TplName = "template/customer-management.html"
			}

	case false:
		log.Println(helpers.ServerConnectionError)
	}
}


/*To delete selected record from database*/

func (c *CustomerManagementController)LoadDeleteCustomerManagement() {
	w := c.Ctx.ResponseWriter
	log.Println("delete inside")
	customerManagementId :=c.Ctx.Input.Param(":customermanagementid")
	log.Println("customer id",customerManagementId)
	dbStatus:= models.DeleteCustomerManagementData(c.AppEngineCtx,customerManagementId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}


}





