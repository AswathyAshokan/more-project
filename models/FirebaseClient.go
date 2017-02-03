/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"github.com/astaxie/beegae"
	"gopkg.in/zabawaba99/firego.v1"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/jwt"
)


var firebaseConfig *jwt.Config
var firebaseServer string

func GetFirebaseClient(ctx context.Context, path string)(*firego.Firebase, error){
	if firebaseConfig==nil {
		jsonKey, err := ioutil.ReadFile("conf/serviceAccountCredentials.json") // or path to whatever name you downloaded the JWT to
		if err != nil {
			return nil, err
		}
		firebaseConfig, err = google.JWTConfigFromJSON(jsonKey, "https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/firebase.database")
		if err != nil {
			return nil, err
		}
		firebaseServer=beegae.AppConfig.String("FirebaseServer")
	}
	client := firebaseConfig.Client(ctx)
	f := firego.New( firebaseServer +  path, client)
	return  f, nil
}
