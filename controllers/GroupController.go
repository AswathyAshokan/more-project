/*Created By Farsana*/
package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	//"google.golang.org/appengine/log"
	//"google.golang.org/appengine"
	"reflect"
	"net/http"
	"log"
)

type GroupController struct {
	BaseController
}
func (c *GroupController) AddGroup() {
	r := c.Ctx.Request
	if r.Method == "POST" {
		user := models.InviteUser{}
		result :=user.DropDown(c.AppEngineCtx)
		log.Print("ffffff",result)
		//dataValue := reflect.ValueOf(result)
		//var valueSlice []models.User
		///var keySlice []string
		//for _, key := range dataValue.MapKeys() {
			//keySlice = append(keySlice, key.String())//to get keys
			//valueSlice = append(valueSlice, result[key.String()])//to get values
			//viewmodel. = append(viewmodel.Groups, result[key.String()])


		//viewmodel.GroupMembers= result.GroupMembers


		group := models.Group{}
		group.GroupName = c.GetString("groupname")
		//group.GroupMembers = c.GetString("addgroup")
		group.AddgroupToDb(c.AppEngineCtx)
	} else {
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-group.html"
	}

}

func (c *GroupController) GroupDetails() {
	//r := c.Ctx.Request
	group := models.Group{}
	result := group.DisplayGroup(c.AppEngineCtx)
	dataValue := reflect.ValueOf(result)
	var valueSlice []models.Group
	viewmodel := viewmodels.Group{}
	var keySlice []string
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())//to get keys
		valueSlice = append(valueSlice, result[key.String()])//to get values
		viewmodel.Groups = append(viewmodel.Groups, result[key.String()])

	}
	viewmodel.Key=keySlice
	c.Data["vm"] = viewmodel
	c.Layout = "layout/layout.html"
	c.TplName = "template/group-details.html"
}

func (c *GroupController) GroupDelete() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	key:=c.Ctx.Input.Param(":Key")
	group := models.Group{}
	result :=group.DeleteGroup(c.AppEngineCtx,key)
	switch result {
	case true:
		http.Redirect(w, r, "/group-details", 301)
	case false:
		log.Println("false")
	}
	//log.Infof(exam, "vvvvv: %v", user)


}

