package gofastbill

import (
	"errors"
	//	"net/http"
)

const (
	fastbill_plus      string = "https://my.fastbill.com/api/1.0/api.php"
	fastbill_automatic string = "https://automatic.fastbill.com/api/1.0/api.php"
)

type Initialization struct {
	Email      string `json:"email"`
	Apikey     string `json:"apikey"`
	Serviceurl string `json:"serviceurl"`
}

func Init(email string, apikey string, service bool) (*Initialization, error) {

	var init Initialization

	if email == "" || apikey == "" {
		return nil, errors.New("Invalid host or auth data to connect")
	}

	init.Email = email
	init.Apikey = apikey

	if !service {
		init.Serviceurl = fastbill_plus
	} else {
		init.Serviceurl = fastbill_automatic
	}

	return &init, nil

}
