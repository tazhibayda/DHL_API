package main

import (
	"fmt"
	"github.com/tazhibayda/DHL_API/dhl"
)

const (
	Account  = ""
	Password = ""
	ApiURL   = "https://api-mock.dhl.com/mydhlapi/"
)

func main() {

	post := dhl.NewDHL(Account, Password, ApiURL)

	fmt.Println(post)

}
