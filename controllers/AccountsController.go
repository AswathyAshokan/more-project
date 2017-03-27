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

/*
func (c *AccountsController) SuperAdminsAccount() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	superAdminViewModel := viewmodels.SuperAdmin{}
	storedSession := ReadSessionForSuperAdmin(w,r)
	log.Println("read session:",storedSession)
	dbStatus,superAdminAllDetails := models.GetAllSuperAdminsDetails(c.AppEngineCtx)
	switch dbStatus {
	case true :dataValue := reflect.ValueOf(superAdminAllDetails)
		for _, key := range dataValue.MapKeys() {
			superAdminViewModel.FirstName = superAdminAllDetails[key.String()].Info.FirstName
			superAdminViewModel.LastName = superAdminAllDetails[key.String()].Info.LastName
			superAdminViewModel.Email = superAdminAllDetails[key.String()].Info.Email
			superAdminViewModel.PhoneNo = superAdminAllDetails[key.String()].Info.PhoneNo


		}


		log.Println("true",superAdminViewModel)
	case false:
		 log.Println("false...",superAdminViewModel)

	}
	c.Data["vm"] = superAdminViewModel
	c.Layout = "layout/layout-superadmin.html"
	c.TplName = "template/accounts.html"
}
*/


func (c *AccountsController) SuperAdminsAccount() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSessionForSuperAdmin(w,r)
	superAdmin := models.SuperAdmins{}
	superAdminViewModel := viewmodels.SuperAdmin{}
	if r.Method == "POST" {
		superAdminId := storedSession.SuperAdminId
		superAdmin.Info.FirstName = c.GetString("superAdminName")
		superAdmin.Info.Email = c.GetString("superadminEmail")
		superAdmin.Info.PhoneNo = c.GetString("superAdminPhone")
		log.Println("from controller :", c.GetString("superAdminName"))
		dbStatus := superAdmin.EditSuperAdminDetails(c.AppEngineCtx, superAdminId)

		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		dbStatus,superAdminAllDetails := models.GetAllSuperAdminsDetails(c.AppEngineCtx)
		switch dbStatus {
		case true :dataValue := reflect.ValueOf(superAdminAllDetails)
			for _, key := range dataValue.MapKeys() {
				superAdminViewModel.FirstName = superAdminAllDetails[key.String()].Info.FirstName
				superAdminViewModel.LastName = superAdminAllDetails[key.String()].Info.LastName
				superAdminViewModel.Email = superAdminAllDetails[key.String()].Info.Email
				superAdminViewModel.PhoneNo = superAdminAllDetails[key.String()].Info.PhoneNo


			}
			log.Println("true",superAdminAllDetails)
		case false:
			log.Println("false...",superAdminAllDetails)

		}
		c.Data["vm"] = superAdminViewModel
		c.Layout = "layout/layout-superadmin.html"
		c.TplName = "template/accounts.html"
	}
}

func (c *AccountsController) ChangeSuperAdminsPassword() {
	log.Println("cp13232222")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSessionForSuperAdmin(w,r)
	superAdmin := models.SuperAdmins{}
	if r.Method == "POST" {
		log.Println("cp2")
		superAdminId := storedSession.SuperAdminId
		enteredOldPassword := (c.GetString("oldPassword"))
		passwordStatus := models.IsEnteredPasswordCorrect(c.AppEngineCtx, superAdminId,enteredOldPassword)
		switch passwordStatus{
		case true :
			log.Println("cp8")
			superAdmin.Info.Password= []byte(c.GetString("confirmpassword"))
			dbStatus := superAdmin.EditSuperAdminPassword(c.AppEngineCtx, superAdminId)
			switch dbStatus {

			case true:
				log.Println("cp15")
				w.Write([]byte("true"))
			case false:
				log.Println("cp16")
				w.Write([]byte("false"))
			}
		case false :
			log.Println("cp17")

		}
	}
}


