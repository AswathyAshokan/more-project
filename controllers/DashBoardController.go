package controllers


type DashBoardController struct {
	BaseController
}
func (c *DashBoardController)LoadDashBoard() {
	c.Layout = "layout/layout.html"
	c.TplName = "template/dash-board.html"

}
