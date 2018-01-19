package controllers


type DuressController struct{
	BaseController
}

//loading duress page
func (c *DuressController) Duress(){
	c.TplName = "template/duress.html"
}
