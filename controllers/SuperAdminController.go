package controllers


type SuperAdminController struct {
	BaseController
}

// add new customer to database
func (c *SuperAdminController) AddSuperAdmin() {
	/*r := c.Ctx.Request
	w := c.Ctx.ResponseWriter*//*
	superAdmin := models.SuperAdmins{}
	superAdmin.Info.Email = "super@admin.com"
	superAdmin.Info.FirstName = "Super"
	superAdmin.Info.LastName = "Admin"
	superAdmin.Info.Password = "JDJhJDEwJGNlQjlqVHovcndwUVYzWWFLRjdXRi5GY1ZQMGQxZXF3SnMvMlBUUWpnOHMyTUhrWDlVdEJt"
	superAdmin.Info.PhoneNo = "1343325464"
	superAdmin.Settings.DateOfCreation = 1488522350
	superAdmin.Settings.Status = "Active"
	dbStatus := superAdmin.AddSuperAdminToDb(c.AppEngineCtx)
	switch dbStatus {
	case true:

	case false:

	}*/

}