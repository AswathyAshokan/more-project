package controllers


type DuressController struct{
	BaseController
}
func (c *DuressController) Duress(){
	c.Layout = "layout/layout.html"
	c.TplName = "template/duress-mode.html"
}
