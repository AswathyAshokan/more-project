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
	ProfilePicture	string
	PaymentStatus   string

}

// for Super Admin
type SessionForAdminValues  struct {
	SuperAdminId		string
	SuperAdminFullName	string
	SuperAdminEmail		string
}
type SessionForPayment struct{
	NumberOfUsers		string
}

var cookieToken = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))


//function for set session values
func SetSession(w http.ResponseWriter, sessionValues SessionValues){

	value := make(map[string]string)
	value["adminId"] = sessionValues.AdminId
	value["adminEmail"] = sessionValues.AdminEmail
	value["adminFirstName"] = sessionValues.AdminFirstName
	value["adminLastName"] = sessionValues.AdminLastName
	value["profilePicture"] = sessionValues.ProfilePicture
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

//function to read the session values
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
				sessionValues.ProfilePicture =value["profilePicture"]
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


//function to clear all session values
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

		//session for Plan
/*-----------------------------------------------------------------------------------------*/

func SessionForPlan(w http.ResponseWriter, r *http.Request) (SessionValues, bool) {
	sessionValues := SessionValues{}
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = cookieToken.Decode("session", cookie.Value, &value); err == nil {
			sessionValues.CompanyTeamName = value["companyTeamName"]
			sessionValues.AdminId = value["adminId"]
			sessionValues.AdminEmail = value["adminEmail"]
			sessionValues.AdminFirstName = value["adminFirstName"]
			sessionValues.AdminLastName = value["adminLastName"]
			sessionValues.ProfilePicture =value["profilePicture"]
			sessionValues.CompanyId = value["companyId"]
			sessionValues.CompanyName = value["companyName"]
			sessionValues.CompanyPlan = value["companyPlan"]



		} else {
			return sessionValues, false
		}
	} else {
		return sessionValues, false
	}
	return sessionValues, true
}

			// session for SuperAdmin :
/*-------------------------------------------------------------------------------------*/
//session for payment


var cookieTokenForPayment = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSessionForPayment(w http.ResponseWriter,sessionForPayment SessionForPayment )  {
	valueOfPayment := make(map[string]string)
	valueOfPayment["noofUsers"] = sessionForPayment.NumberOfUsers
	if encoded, err := cookieTokenForPayment.Encode("sessionForPayment",valueOfPayment);err == nil{
		cookie := &http.Cookie{
			Name:  "sessionForPayment",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w,cookie)
		log.Println("Session is Set!")
	}
}
//read payment Session

func ReadSessionForPayment (w http.ResponseWriter, r *http.Request) (SessionForPayment) {
	sessionValues := SessionForPayment{}
	if cookie, err := r.Cookie("sessionForPayment"); err == nil {
		value := make(map[string]string)
		if err = cookieTokenForPayment.Decode("sessionForPayment", cookie.Value, &value); err == nil {
			sessionValues.NumberOfUsers = value["noofUsers"]
			log.Println("sessionValues")
		} else {
			ClearSessionForPayment(w)
			http.Redirect(w, r, "/login", 302)
			log.Println("Access Denied! You are not logged in!")
		}
	} else {
		http.Redirect(w, r, "/", 302)
		log.Println("Access Denied! You are not logged in!")
	}
	return sessionValues
}


//function to clear session of payments
func ClearSessionForPayment(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessionForPayment",
		Value:  "",
		//Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	log.Println("Logged out Successfully!")
	log.Println("The value in session after Logout:", cookie.Value)

}








var cookieTokenForSuperAdmin = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))


func SetSessionForSuperAdmin(w http.ResponseWriter, sessionValueForSuperAdmin SessionForAdminValues){
	valueOfSuperAdmin := make(map[string]string)
	valueOfSuperAdmin["superAdminEmail"] = sessionValueForSuperAdmin.SuperAdminEmail
	valueOfSuperAdmin["superAdminId"] = sessionValueForSuperAdmin.SuperAdminId
	valueOfSuperAdmin["superAdminName"] = sessionValueForSuperAdmin.SuperAdminFullName
	if encoded, err := cookieTokenForSuperAdmin.Encode("sessionForSuperAdmin",valueOfSuperAdmin);err == nil{
		cookie := &http.Cookie{
			Name:  "sessionForSuperAdmin",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w,cookie)
		log.Println("Session is Set!")
	}
}



func ReadSessionForSuperAdmin (w http.ResponseWriter, r *http.Request) (SessionForAdminValues) {
	sessionValues := SessionForAdminValues{}
	if cookie, err := r.Cookie("sessionForSuperAdmin"); err == nil {
		value := make(map[string]string)
		if err = cookieTokenForSuperAdmin.Decode("sessionForSuperAdmin", cookie.Value, &value); err == nil {
			sessionValues.SuperAdminFullName = value["superAdminName"]
			sessionValues.SuperAdminId = value["superAdminId"]
			sessionValues.SuperAdminEmail = value["superAdminEmail"]
			log.Println("sessionValues")
		} else {
			ClearSessionForSuperAdmin(w)
				http.Redirect(w, r, "/login", 302)
				log.Println("Access Denied! You are not logged in!")
			}
	} else {
		http.Redirect(w, r, "/login", 302)
		log.Println("Access Denied! You are not logged in!")
	}
	return sessionValues
}


func ClearSessionForSuperAdmin(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessionForSuperAdmin",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	log.Println("Logged out Successfully!")
	log.Println("The value in session after Logout:", cookie.Value)

}

/*_____________________________________________________________________________________________*/