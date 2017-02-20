/* Author: Sarath
Date: 02/02/17 */
package controllers

import (
	"github.com/gorilla/securecookie"
	"app/passporte/models"
	"net/http"
	"log"
)

type SessionController struct{
	BaseController
}

type SessionValues struct{
	Info  models.Info
	Settings models.Settings
}

var cookieToken = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(w http.ResponseWriter, adminDetails models.CompanyAdmins){
	value := map[string]string{
		"email": adminDetails.Info.Email,
		"firstName": adminDetails.Info.FirstName,
		"lastName": adminDetails.Info.LastName,
	}
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
func ReadSession (w http.ResponseWriter, r *http.Request) (SessionValues) {
	sessionValues := SessionValues{}
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieToken.Decode("session", cookie.Value, &cookieValue); err == nil {
			sessionValues.Info.Email = cookieValue["email"]
			sessionValues.Info.FirstName = cookieValue["firstName"]
			sessionValues.Info.LastName = cookieValue["lastName"]
		}
	} else {
		http.Redirect(w, r, "/", 302)
		log.Println("Access Denied! You are not logged in!")
	}
	return sessionValues
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
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
