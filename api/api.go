package api

import (
	"gopkg.in/resty.v1"
	"obas/config"
)

//const BASE_URL string = "https://ict.cput.ac.za/obasapi"
const BASE_URL string = "http://155.238.32.64:9000"

func Rest() *resty.Request {
	return resty.R().SetAuthToken("").
		SetHeader("Accept", "bursary/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "bursary/json")
}

var JSON = config.ConfigWithCustomTimeFormat
