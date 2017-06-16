package helpers

import (
	"log"
	"os"
	"bufio"
	"strings"
	//"app/passporte/models"

	"gopkg.in/zabawaba99/firego.v1"

	"reflect"
	"encoding/json"
	"io/ioutil"
	"golang.org/x/oauth2/jwt"
	"golang.org/x/oauth2/google"
	"github.com/astaxie/beegae"

	"golang.org/x/net/context"
)

type Country struct {
	CountryName	string
	DialCode	string
	States   	map[string]string
}

type Code struct {
	Name		string	`json:"name"`
	DialCode 	string	`json:"dial_code"`
	CountryCode  	string	`json:"code"`
}



func ReadTextFile(countryFileLocation, codeFileLocation string, ctx  context.Context) {
	countries := make(map[string]Country)
	countryFileData, err := os.Open(countryFileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer countryFileData.Close()
	scanner := bufio.NewScanner(countryFileData)

	codeFileData, err := ioutil.ReadFile(codeFileLocation)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	var codeSlice []Code

	err = json.Unmarshal(codeFileData, &codeSlice)
	if err != nil {
		log.Println(err)
	}


	var lines []string
	for  scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	for  i := 0; i < len(lines); i++ {
		country := Country{}
		//lines = append(lines, scanner.Text())
		lineSlice := strings.Split(lines[i], ",")
		if len(countries) == 0{
			country.CountryName = lineSlice[5]

			/*------------states----------------------*/
			state := make(map[string]string)
			for j := 0; j < len(lines); j++ {
				stateLineSlice := strings.Split(lines[j], ",")
				if stateLineSlice[2] == lineSlice[2] {
					if len(state) == 0 {
						state[stateLineSlice[1]] = stateLineSlice[4]
					} else {
						stateValue := reflect.ValueOf(state)
						stateFlag := 0
						for _, key := range stateValue.MapKeys() {
							if key.String() == stateLineSlice[1] {
								stateFlag++
							}
						}
						if stateFlag == 0 {
							state[stateLineSlice[1]] = stateLineSlice[4]
						}
					}
				}
			}
			country.States = state
			/*------------states ends-----------------*/

			/*------------Dial Code-------------------*/
			for k := 0; k < len(codeSlice); k++ {
				if codeSlice[k].CountryCode == lineSlice[2] {
					country.DialCode = codeSlice[k].DialCode
				}
			}
			/*------------Dial Code Ends--------------*/

			countries[lineSlice[2]] = country
		} else {
			flag := 0
			// dataValue := reflect.ValueOf(countries)
			reflectedCountries := reflect.ValueOf(countries)
			for _, key := range reflectedCountries.MapKeys() {
				if key.String() == lineSlice[2] {
					flag++
				}
			}
			if flag == 0 {
				country.CountryName = lineSlice[5]

				/*------------states----------------------*/
				state := make(map[string]string)
				for j := 0; j < len(lines); j++ {
					stateLineSlice := strings.Split(lines[j], ",")
					if stateLineSlice[2] == lineSlice[2] {
						if len(state) == 0 {
							state[stateLineSlice[1]] = stateLineSlice[4]
						} else {
							stateValue := reflect.ValueOf(state)
							stateFlag := 0
							for _, key := range stateValue.MapKeys() {
								if key.String() == stateLineSlice[1] {
									stateFlag++
								}
							}
							if stateFlag == 0 {
								state[stateLineSlice[1]] = stateLineSlice[4]
							}
						}
					}
				}
				country.States = state

				/*------------states ends-----------------*/

				/*------------Dial Code-------------------*/
				for k := 0; k < len(codeSlice); k++ {
					if codeSlice[k].CountryCode == lineSlice[2] {
						country.DialCode = codeSlice[k].DialCode
					}
				}
				/*------------Dial Code Ends--------------*/

				countries[lineSlice[2]] = country
			}
		}
	}

	//return countries
	db,err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	countryValues := reflect.ValueOf(countries)
	for _, key := range countryValues.MapKeys() {
		err = db.Child("/CountryDetails/" + key.String()).Set(countries[key.String()])
		if err != nil {
			log.Println(err)
		}
	}
}

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


