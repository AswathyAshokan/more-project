package controllers

import (
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"fmt"
	"app/passporte/viewmodels"
)
type PaymentController struct {
	BaseController
}
// PayPal variables -- change it according to your requirements.
// NOTE : I decided to use individual variables here so that I can include
//        explanations. It would be better and "cleaner" to put these variables
//        into array(such as url.Values{} https://golang.org/pkg/net/url/#Values ) and
//        loop the array for name and values in the form hidden fields below.

var currency_code = "USD"
var business = "aswathyashok85-buyer@gmail.com"                                 //PayPal account to receive money
var image_url = "https://d1ohg4ss876yi2.cloudfront.net/logo35x35.png" // image on top of PayPal

// change socketloop.com:8080 to your domain name
// REMEMBER : localhost won't work and IPN simulator only deal with port 80 or 443
var cancel_return = "http://localhost:8080/paymentcancelreturn"
var return_url = "http://localhost:8080/paymentSuccessPage"// return is Golang's keyword
var notify_url = "http://localhost:8080/ipn"            // <--- important for IPN to work!

// just an example for custom field, could be username, etc. Use custom field
// for extra verification purpose or to mark PAID status in
// in database, etc.

var custom = "donation"

// See
// https://developer.paypal.com/docs/classic/paypal-payments-standard/integration-guide/Appx_websitestandard_htmlvariables/
// for the meaning of rm and _xclick

var rm = "2" // rm 2 equal Return method = POST
var cmd = "_xclick"
var item_name = "Donation for SocketLoop"
var quantity = "1"
var amount = "5" // keeping it simple for this tutorial. You should accept the amount
// from a form instead of hard coding it here.

// uncomment to switch to real PayPal instead of sandbox
//var paypal_url = "https://www.paypal.com/cgi-bin/webscr"
var paypal_url = "https://www.sandbox.paypal.com/cgi-bin/webscr"
func (c *PaymentController)Home() {
	log.Println("iam in payment page ")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	sessionValues, _ := SessionForPlan(w,r)
	companyTeamName := sessionValues.CompanyTeamName
	log.Println("company name",companyTeamName)
	companyPlan := c.Ctx.Input.Param(":companyPlan")
	log.Println("company plan",companyPlan)
	amount := c.Ctx.Input.Param(":paymentPrice")
	noOfUsers :=c.Ctx.Input.Param(":NumberOfUsers")
	sessionForPayment := SessionForPayment{}
	sessionForPayment.NumberOfUsers=noOfUsers
	SetSessionForPayment(w, sessionForPayment)

	//sessionValues.NumberOfUsers = noOfUsers
	companyName :=sessionValues.CompanyName
	html := "<html><body><h1>You will be directed to PayPal now to pay USD " + amount + " to SocketLoop!</h1>"
	html = html + "<form action=' " + paypal_url + "' method='post'>"
	// now add the PayPal variables to be posted
	// a cleaner way to create an array and use for loop
	html = html + "<input type='hidden' name='currency_code' value='" + currency_code + "'>"
	html = html + "<input type='hidden' name='business' value='" + business + "'>"
	html = html + "<input type='hidden' name='image_url' value='" + image_url + "'>"
	html = html + "<input type='hidden' name='cancel_return' value='" + cancel_return + "'>"
	html = html + "<input type='hidden' name='notify_url' value='" + notify_url + "'>"
	html = html + "<input type='hidden' name='return' value='" + return_url +"?numberofusers="+noOfUsers + "'>" //use return instead of return_url
	html = html + "<input type='hidden' name='company Name' value='" + companyName + "'>"
	html = html + "<input type='hidden' name='rm' value='" + rm + "'>"
	html = html + "<input type='hidden' name='cmd' value='" + cmd + "'>"
	html = html + "<input type='hidden' name='Selected Plan' value='" + companyPlan + "'>"
	//html = html + "<input type='hidden' name='quantity' value='" + quantity + "'>"
	html = html + "<input type='hidden' name='amount' value='" + amount + "'>"

	html = html + " <input type='submit' value='Proceed to PayPal'></form></body></html>"

	w.Write([]byte(fmt.Sprintf(html)))

}
func (c *PaymentController)PaymentSuccess() {
	//w := c.Ctx.ResponseWriter
	//// This is where you would probably want to thank the user for their order
	//// or what have you.  The order information at this point is in POST
	//// variables.  However, you don't want to "process" the order until you
	//// get validation from the IPN.  That's where you would have the code to
	//// email an admin, update the database with payment status, activate a
	//// membership, etc.
	//
	//html := "<html><body><h1>Thank you! Payment accepted!</h1></body></html>"
	//w.Write([]byte(fmt.Sprintf(html)))
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	sessionValues, _ := SessionForPlan(w,r)
	companyTeamName := sessionValues.CompanyTeamName
	log.Println("company name",companyTeamName)
	//NumberOfUsers :=sessionValues.NumberOfUsers
	//log.Println("number of users",NumberOfUsers)
	viewModel :=viewmodels.PaymentViewModel{}
	viewModel.CompanyTeamName=companyTeamName
	//viewModel.NumberOfUsers =NumberOfUsers
	c.Data["vm"] = viewModel
	c.TplName = "template/paymentSucessOfWeb.html"
	}








func (c *PaymentController)PurchaseSuccess() {
	w := c.Ctx.ResponseWriter
	// This is where you would probably want to thank the user for their order
	// or what have you.  The order information at this point is in POST
	// variables.  However, you don't want to "process" the order until you
	// get validation from the IPN.  That's where you would have the code to
	// email an admin, update the database with payment status, activate a
	// membership, etc.

	html := "<html><body><h1>Thank you! Payment accepted!</h1></body></html>"
	w.Write([]byte(fmt.Sprintf(html)))
}

func (c *PaymentController)PaymentCancelReturn() {
	w := c.Ctx.ResponseWriter
	html := "<html><body><h1>Oh ok. Payment cancelled!</h1></body></html>"
	w.Write([]byte(fmt.Sprintf(html)))
}
func (c *PaymentController)IPN() {
	r := c.Ctx.Request


	// Payment has been received and IPN is verified.  This is where you
	// update your database to activate or process the order, or setup
	// the database with the user's order details, email an administrator,
	// etc. You can access a slew of information via the IPN data from r.Form

	// Check the paypal documentation for specifics on what information
	// is available in the IPN POST variables.  Basically, all the POST vars
	// which paypal sends, which we send back for validation.

	// For this tutorial, we'll just print out all the IPN data.

	log.Println("IPN received from PayPal")

	err := r.ParseForm() // need this to get PayPal's HTTP POST of IPN data

	if err != nil {
		log.Println(err)
		return
	}

	if r.Method == "POST" {

		var postStr string = paypal_url + "&cmd=_notify-validate&"

		for k, v := range r.Form {
			log.Println("key :", k)
			log.Println("value :", strings.Join(v, ""))

			// NOTE : Store the IPN data k,v into a slice. It will be useful for database entry later.

			postStr = postStr + k + "=" + url.QueryEscape(strings.Join(v, "")) + "&"
		}

		// To verify the message from PayPal, we must send
		// back the contents in the exact order they were received and precede it with
		// the command _notify-validate

		// PayPal will then send one single-word message, either VERIFIED,
		// if the message is valid, or INVALID if the messages is not valid.

		// See more at
		// https://developer.paypal.com/webapps/developer/docs/classic/ipn/integration-guide/IPNIntro/

		// post data back to PayPal
		client := &http.Client{}
		req, err := http.NewRequest("POST", postStr, nil)

		if err != nil {
			log.Println(err)
			return
		}

		req.Header.Add("Content-Type: ", "application/x-www-form-urlencoded")

		// fmt.Println(req)

		resp, err := client.Do(req)

		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Response : ")
		log.Println(resp)
		log.Println("Status :")
		log.Println(resp.Status)


		// convert response to string
		respStr, _ := ioutil.ReadAll(resp.Body)

		//fmt.Println("Response String : ", string(respStr))

		verified, err := regexp.MatchString("VERIFIED", string(respStr))

		if err != nil {
			log.Println(err)
			return
		}

		if verified {
			log.Println("IPN verified")
			log.Println("TODO : Email receipt, increase credit, etc")
		} else {
			log.Println("IPN validation failed!")
			log.Println("Do not send the stuff out yet!")
		}

	}

}
func (c *PaymentController)AppPurchaseSuccess() {
	c.TplName = "template/paymentSucessApp.html"
}