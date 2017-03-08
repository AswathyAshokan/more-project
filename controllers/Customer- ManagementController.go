package controllers



type CustomerManagementController struct {
	BaseController
}

func (c *CustomerManagementController) CustomerManagement() {
	/*r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("The userDetails stored in session:",storedSession)
	customerViewModel := viewmodels.Customer{}
	allCustomer,dbStatus:= models.GetAllCustomerDetails(c.AppEngineCtx,companyTeamName)
	log.Println("view",allCustomer)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allCustomer)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		log.Println("key",keySlice)
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.CustomerName)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Address)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.State)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.ZipCode)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Email)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Phone)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.ContactPerson)
			//tempValueSlice = append(tempValueSlice,allCustomer[k].Info.)
			customerViewModel.Values=append(customerViewModel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		customerViewModel.Keys = keySlice
		customerViewModel.CompanyTeamName = storedSession.CompanyTeamName
		customerViewModel.CompanyPlan = storedSession.CompanyPlan
		c.TplName = "template/customer-management.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}*/
}

