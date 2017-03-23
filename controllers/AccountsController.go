package controllers

import (
	"log"
	"app/passporte/models"
	"reflect"
	"app/passporte/viewmodels"
)

type AccountsController struct {
	BaseController
}

func (c *AccountsController) SuperAdminsAccount() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	superAdminViweModel := viewmodels.SuperAdmin{}
	storedSession := ReadSessionForSuperAdmin(w,r)
	log.Println("read session:",storedSession)
	dbStatus,superAdminAllDetails := models.GetAllSuperAdminsDetails(c.AppEngineCtx)
	switch dbStatus {
	case true :dataValue := reflect.ValueOf(superAdminAllDetails)
		for _, key := range dataValue.MapKeys() {
			/*tempValueSlice = append(tempValueSlice,superAdminAllDetails[key.String()].Info.Email)
			tempValueSlice = append(tempValueSlice,superAdminAllDetails[key.String()].Info.FirstName)
			tempValueSlice = append(tempValueSlice,superAdminAllDetails[key.String()].Info.LastName)
			tempValueSlice = append(tempValueSlice,superAdminAllDetails[key.String()].Info.PhoneNo)
			superAdminViwemodel.Values = append(superAdminViwemodel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]*/
			superAdminViweModel.FirstName = superAdminAllDetails[key.String()].Info.FirstName
			superAdminViweModel.LastName = superAdminAllDetails[key.String()].Info.LastName
			superAdminViweModel.Email = superAdminAllDetails[key.String()].Info.Email
			superAdminViweModel.PhoneNo = superAdminAllDetails[key.String()].Info.PhoneNo


		}


		log.Println("true",superAdminAllDetails)
	case false:
		 log.Println("false...",superAdminAllDetails)

	}
	c.Data["vm"] = superAdminViweModel
	c.Layout = "layout/layout-superadmin.html"
	c.TplName = "template/accounts.html"
}


func (c *AccountsController) EditSuperAdmin() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSessionForSuperAdmin(w, r)
	log.Println("read session:", storedSession)
	if r.Method == "POST" {
		log.Println("cp3")
		superAdmin := models.SuperAdmins{}
		/*superAdmin.Info.FirstName = c.GetString("superAdminName")
		superAdmin.Info.Email = c.GetStrings("superadminEmail")
		superAdmin.Info.PhoneNo = c.GetStrings("phone")*/
		dbStatus := superAdmin.EditSuperAdminDetails(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {

	}
}

