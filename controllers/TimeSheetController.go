package controllers
import (

	//"app/passporte/models"
	//"app/passporte/viewmodels"
	//"time"
	//"reflect"
	//"app/passporte/helpers"
	//"log"
	//"strconv"
	//"strings"
	//"fmt"
)
type TimeSheetController struct {
	BaseController
}
func (c *TimeSheetController)LoadTimeSheetDetails() {
	c.Layout = "layout/layout.html"
	c.TplName = "template/time-sheet.html"

}

