
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
	w :=c.Ctx.ResponseWriter
	if r.Method == "POST" {

		project:=models.Project{}
		project.CustomerName= c.GetString("customerName")
		project.ProjectName= c.GetString("projectName")
		project.ProjectNumber = c.GetString("projectNumber")
		project.NumberOfTask = c.GetString("numberOfTask")
		context := appengine.NewContext(r)
		log.Infof(context, "requested struct: %+v", project)
		project.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		project.Status = "Open"
		dbStatus :=project.AddProjectToDB(c.AppEngineCtx)
		switch dbStatus {

		case true:
			w.Write([]byte("true"))

		case false:
			w.Write([]byte("false"))
		}

	}else {

		c.Layout = "layout/layout.html"
		c.TplName = "template/add-project.html"

	}



}
func (c *ProjectController)LoadProjectDetail() {
	project := models.Project{}
	dbStatus, projects := project.RetrieveProjectFromDB(c.AppEngineCtx)
	viewModel := viewmodels.ProjectViewModel{}

	switch dbStatus {

	case true:
		r := c.Ctx.Request
		context := appengine.NewContext(r)
		log.Infof(context, "%s\n", projects)
		//var valueSlice []models.User
		dataValue := reflect.ValueOf(projects)
		var keySlice []string
		var valueSlice []models.Project
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		for _, k := range keySlice {
			valueSlice = append(valueSlice, projects[k])
			viewModel.Project = append(viewModel.Project, projects[k])
			viewModel.Key=keySlice

		}
		log.Infof(context,"Key:", keySlice, "Value:", valueSlice)
		log.Infof(context, "typeeee",(reflect.TypeOf(viewModel)))
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/project-details.html"

	case false:
	}


}
func (c *ProjectController)LoadDeleteProject() {

	r := c.Ctx.Request
	context := appengine.NewContext(r)
	projectId :=c.Ctx.Input.Param(":projectId")
	log.Infof(context,"idddddddddd", projectId)
	project := models.Project{}
	dbStatus := project.DeleteProjectFromDB(c.AppEngineCtx, projectId)

	switch dbStatus {

	case true:
		c.Redirect("/project", 302)
	case false :
	}


}