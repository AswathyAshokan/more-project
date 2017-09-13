package controllers


type DuressController struct{
	BaseController
}
func (c *DuressController) Duress(){
	c.TplName = "template/duress.html"
}
