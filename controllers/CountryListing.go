//created by farsana
// for displaying countrycodes

package controllers

import (
	"app/passporte/helpers"
	"log"
)
type CountryController struct {
	BaseController
}

//to Display Plan Details
func (c *CountryController) ListCountries() {
	log.Println("gggggg")
	countryFileLocation := "./datafiles/data/cities.csv"
	codeFileLocation := "./datafiles/data/CountryCodes.json"
	helpers.ReadTextFile(countryFileLocation, codeFileLocation, c.AppEngineCtx)
}
