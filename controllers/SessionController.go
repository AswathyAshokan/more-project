/* Author: Sarath
Date: 02/02/17 */
package controllers

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"log"
)

type SessionController struct{
	BaseController
}

type SessionValues struct{
	AdminId		string
	AdminFirstName	string
	AdminLastName	string
	AdminEmail	string
	CompanyId	string
	CompanyName	string
	CompanyTeamName	string
	CompanyPlan	string
}

var cookieToken = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(w http.ResponseWriter, sessionValues SessionValues){

	value := make(map[string]string)
	value["adminId"] = sessionValues.AdminId
	value["adminEmail"] = sessionValues.AdminEmail
	value["adminFirstName"] = sessionValues.AdminFirstName
	value["adminLastName"] = sessionValues.AdminLastName
	value["companyId"] = sessionValues.CompanyId
	value["companyName"] = sessionValues.CompanyName
	value["companyTeamName"] = sessionValues.CompanyTeamName
	value["companyPlan"] = sessionValues.CompanyPlan

	if encoded, err := cookieToken.Encode("session",value);err == nil{
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w,cookie)
		log.Println("Session is Set!")
	}
}

func ReadSession (w http.ResponseWriter, r *http.Request, companyTeamName string) (SessionValues) {
	sessionValues := SessionValues{}
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = cookieToken.Decode("session", cookie.Value, &value); err == nil {
			sessionValues.CompanyTeamName = value["companyTeamName"]
			if sessionValues.CompanyTeamName == companyTeamName {
				sessionValues.AdminId = value["adminId"]
				sessionValues.AdminEmail = value["adminEmail"]
				sessionValues.AdminFirstName = value["adminFirstName"]
				sessionValues.AdminLastName = value["adminLastName"]
				sessionValues.CompanyId = value["companyId"]
				sessionValues.CompanyName = value["companyName"]
				sessionValues.CompanyPlan = value["companyPlan"]

			} else {
				ClearSession(w)
				http.Redirect(w, r, "/login", 302)
				log.Println("Access Denied! You are not logged in!")
			}

		} else {
			http.Redirect(w, r, "/login", 302)
			log.Println("Access Denied! You are not logged in!")
		}
	} else {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
		log.Println("Access Denied! You are not logged in!")
	}
	return sessionValues
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	log.Println("Logged out Successfully!")
	log.Println("The value in session after Logout:", cookie.Value)

}
