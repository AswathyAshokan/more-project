package controllers

type FitToWorkController struct {
	BaseController
}
func (c *FitToWorkController)AddNewFitToWork() {

	c.Layout = "layout/layout.html"
	c.TplName = "template/add-fit-work.html"
}