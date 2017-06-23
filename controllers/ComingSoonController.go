package controllers

import "log"

type ComingSoonController struct {
	BaseController
}

//to Display Plan Details
func (c *ComingSoonController) LoadComingSoonController() {
	log.Println("gggggggggggggggg")
	c.TplName = "template/pendingWorks.html"
}
