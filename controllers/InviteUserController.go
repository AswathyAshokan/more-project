/*Created By Farsana*/
package controllers

import (
	"app/passporte/models"
	"time"
	"reflect"
	"app/passporte/viewmodels"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	//"time"
	//"reflect"
	//"app/passporte/helper"

	"net/http"
)

type InviteUserController struct {
	BaseController
}

func (c *InviteUserController) AddUser() {
	user := models.InviteUser{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {

		user.FirstName = c.GetString("firstname")
		user.LastName = c.GetString("lastname")
		user.EmailId = c.GetString("emailid")
		user.UserType = c.GetString("usertype")
		user.DateOfCreation =(time.Now().UnixNano() / 1000000)
		user.Status = "inactive"
		dbStatus := user.AdduserToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:


		}
	} else {
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-invite-user.html"
	}
}

func (c *InviteUserController) UserDetails() {

	r := c.Ctx.Request
	exam := appengine.NewContext(r)

	user := models.InviteUser{}
	result := user.DisplayUser(c.AppEngineCtx)
	dataValue := reflect.ValueOf(result)
	var valueSlice []models.InviteUser
	viewmodel := viewmodels.UserViewModel{}
	var keySlice []string
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())//to get keys
		valueSlice = append(valueSlice, result[key.String()])//to get values
		viewmodel.Users = append(viewmodel.Users, result[key.String()])


	}
	viewmodel.Key=keySlice
	log.Infof(exam, "key of",viewmodel)
	c.Data["vm"] = viewmodel
	c.Layout = "layout/layout.html"
	c.TplName = "template/invite-user-details.html"
}

//delete each users




func (c *InviteUserController) UserDelete() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	key:=c.Ctx.Input.Param(":Key")
	exam := appengine.NewContext(r)
	user := models.InviteUser{}
	result :=user.DeleteUser(c.AppEngineCtx,key)
	switch result {
	case true:
		http.Redirect(w, r, "/user-details", 301)
	case false:
		log.Infof(exam,"failed")

	}
	//log.Infof(exam, "vvvvv: %v", user)


}

//edit profile of each users

func (c *InviteUserController) UserEdit() {
	user := models.InviteUser{}
	r := c.Ctx.Request
	key:=c.Ctx.Input.Param(":Key")
	exam := appengine.NewContext(r)
	result,DbStatus :=user.EditUser(c.AppEngineCtx,key)
	switch DbStatus {
	case true:
		viewmodel := viewmodels.UserViewModel{}
		viewmodel.FirstName = result.FirstName
		viewmodel.LastName = result.LastName
		viewmodel.EmailId = result.EmailId
		viewmodel.UserType = result.UserType
		c.Data["vm"] = viewmodel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-user.html"
	case false:
		log.Infof(exam,"failed")

	}
	log.Infof(exam,"jhfjsgjgj: %+v",result)
}

//view the user

func (c *InviteUserController) UserView() {
	r := c.Ctx.Request
	//var Key int
	key:=c.Ctx.Input.Param(":Key")
	exam := appengine.NewContext(r)
	log.Infof(exam, "iddddddddd: %v", key)

}


