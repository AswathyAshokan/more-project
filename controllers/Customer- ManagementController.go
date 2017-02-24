package controllers

type CustomerManagementController struct {
	BaseController
}

func (c *CustomerManagementController) CustomerManagement() {

	c.TplName = "template/customer-management.html"
}
