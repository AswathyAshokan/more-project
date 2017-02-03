
/* Author :Aswathy Ashok */

package controllers

import (

	//"github.com/astaxie/beegae"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"time"

	"reflect"

)

type ProjectController struct {
	BaseController
}

func (c *ProjectController)LoadProject() {
	r := c.Ctx.Request
	if r.Method == "POST" {

		project:=models.Project{}
		project.CustomerName= c.GetString("customerName")
		project.ProjectName= c.GetString("projectName")
		project.ProjectNumber = c.GetString("projectNumber")
		project.NumberOfTask = c.GetString("projectNumber")

		ce := appengine.NewContext(r)
		log.Infof(ce, "requested struct: %+v", project)
		project.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		project.Status = "Open"
		project.AddToDB(c.AppEngineCtx)

	}else {

		c.Layout = "layout/layout.html"
		c.TplName = "template/add-project.html"


	}



}
func (c *ProjectController)LoadProjectDetail() {
	project := models.Project{}
	dbStatus, projects := project.RetrieveFromDB(c.AppEngineCtx)
	viewModel := viewmodels.ProjectViewModel{}

	switch dbStatus {

	case true:
		r := c.Ctx.Request
		ce := appengine.NewContext(r)
		log.Infof(ce, "%s\n", projects)
		//var valueSlice []models.User
		dataValue := reflect.ValueOf(projects)
		var keySlice []string
		var valueSlice []models.Project
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		// To perform the opertion you want
		for _, k := range keySlice {
			valueSlice = append(valueSlice, projects[k])
			viewModel.Project = append(viewModel.Project, projects[k])
			viewModel.Key=keySlice

		}
		log.Infof(ce,"Key:", keySlice, "Value:", valueSlice)
		//log.Infof(ce,"Value: ", valueSlice)
		//log.Infof(ce,"Value: ", valueSlice)
		//mvVar := map["Name"].(string)
		//m := f.(map[string]interface{}
		//viewModel.Name = contact[result[i]].Name
		//viewModel.PhoneNumber = contact["PhoneNumber"]
		//viewModel.Email = contact["Email"]
		//viewModel.Address = contact["Address"]
		//viewModel.State = contact["State"]
		//viewModel.ZipCode = contact["ZipCode"]
		log.Infof(ce, "typeeee",(reflect.TypeOf(viewModel)))
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/project-details.html"

	case false:

	}


}
func (c *ProjectController)LoadDeleteProject() {

	r := c.Ctx.Request
	ce := appengine.NewContext(r)
	id :=c.Ctx.Input.Param(":key")
	log.Infof(ce,"idddddddddd",id)
	project := models.Project{}
	dbStatus := project.DeleteFromDB(c.AppEngineCtx, id )

	switch dbStatus {

	case true:
		c.Redirect("/projectdetail", 302)
	case false :
	}


}