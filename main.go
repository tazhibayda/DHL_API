//package main
//
//import (
//	"fmt"
//	"github.com/tazhibayda/DHL_API/dhl"
//)
//
//const (
//	Account  = ""
//	Password = ""
//	ApiURL   = "https://api-mock.dhl.com/mydhlapi/"
//)
//
//func main() {
//
//	post := dhl.NewDHL(Account, Password, ApiURL)
//
//	fmt.Println(post)
//
//}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api-mock.dhl.com/mydhlapi/address-validate?type=SOME_STRING_VALUE&countryCode=SOME_STRING_VALUE"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Message-Reference", "d0e7832e-5c98-11ea-bc55-0242ac13")
	req.Header.Add("Message-Reference-Date", "Wed, 21 Oct 2015 07:28:00 GMT")
	req.Header.Add("Authorization", "Basic ZGVtby1rZXk6ZGVtby1zZWNyZXQ=")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
