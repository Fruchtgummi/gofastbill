package gofastbill

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var debug = false

const (
	fastbill_plus      string = "https://my.fastbill.com/api/1.0/api.php"
	fastbill_automatic string = "https://automatic.fastbill.com/api/1.0/api.php"
)

type Initialization struct {
	Email      string `json:"email"`
	Apikey     string `json:"apikey"`
	Serviceurl string `json:"serviceurl"`
	Typ        string
}

type Body struct {
	FBAPI    string
	SERVICE  string
	LIMIT    string
	OFFSET   string
	FILTER   string
	DATA     string
	REQUEST  string
	RESPONSE string
	ERRORS   string
}

type REQUEST struct {
	SERVICE string `xml:"SERVICE,omitempty" json:"SERVICE,omitempty"`
	FILTER  FILTER `xml:"FILTER,omitempty" json:"FILTER,omitempty"`
}

type RESPONSE struct {
	STATUS      string `xml:"STATUS,omitempty" json:"STATUS,omitempty"`
	CUSTOMER_ID string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`

	CUSTOMERS []CUSTOMER `xml:"CUSTOMERS,omitempty" json:"CUSTOMERS,omitempty"`
	CONTACTS  []CONTACT  `xml:"CONTACTS,omitempty" json:"CONTACTS,omitempty"`
}

type FILTER struct {
	CUSTOMER_ID     string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CUSTOMER_NUMBER string `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`
	COUNTRY_CODE    string `xml:"COUNTRY_CODE,omitempty" json:"CUSTOMER_CODE,omitempty"`
	CITY            string `xml:"CITY,omitempty" json:"CITY,omitempty"`
	TERM            string `xml:"TERM,omitempty" json:"TERM,omitempty"`
	CONTACT_ID      string `xml:"CONTACT_ID,omitempty" json:"CONTACT_ID,omitempty"`
}

type ERRORS struct {
	ERROR string
}

type FBAPI struct {
	SERVICE  string   `xml:"SERVICE,omitempty" json:"SERVICE,omitempty"`
	DATA     DATA     `xml:"DATA,omitempty" json:"DATA,omitempty"`
	REQUEST  REQUEST  `xml:"REQUEST,omitempty" json:"REQUEST,omitempty"`
	RESPONSE RESPONSE `xml:"RESPONSE,omitempty" json:"RESPONSE,omitempty"`
	FILTER   FILTER   `xml:"FILTER,omitempty" json:"FILTER,omitempty"`
}

// emailadress; API-Key; false=plus | true=automatic; debugging=true|false; string xml|json
func Init(email string, apikey string, service bool, debugging bool, typ string) (*Initialization, error) {

	var init Initialization

	debug = debugging

	if email == "" || apikey == "" {
		return nil, errors.New("Invalid host or auth data to connect")
	}

	init.Email = email
	init.Apikey = apikey
	init.Typ = typ

	if !service {
		init.Serviceurl = fastbill_plus
	} else {
		init.Serviceurl = fastbill_automatic
	}

	return &init, nil

}
func (s *Initialization) FastbillRequest(xmlbody string) (*FBAPI, error) {

	var buf []byte
	var fbapi *FBAPI

	body := strings.NewReader(xmlbody)
	req, err := http.NewRequest("POST", s.Serviceurl, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.Email, s.Apikey)
	req.Header.Set("Content-Type", "application/json")

	if s.Typ == "json" {
		req.Header.Set("Content-Type", "application/json")
	} else if s.Typ == "xml" {
		req.Header.Set("Content-Type", "application/xml")
	} else {
		return nil, errors.New(s.Typ + ": Do you use JSON or XML!")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if debug != false {
		log.Println(os.Stdout, string(buf))
	}

	if s.Typ == "json" {
		e := json.Unmarshal(buf, &fbapi)
		if e == nil {
			log.Println(e)
		}

	} else if s.Typ == "xml" {
		xml.Unmarshal(buf, &fbapi)
	}

	return fbapi, nil

}

func (s *Initialization) GenerateRequest(x FBAPI) (string, error) {

	if s.Typ == "json" {

		js, err := json.Marshal(x)
		if err != nil {

			return "", err
		}

		return string(js), nil

	} else if s.Typ == "xml" {
		output, err := xml.MarshalIndent(x, "  ", "    ")
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		return "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" + string(output), nil
	} else {
		return "", errors.New(s.Typ + ": Do you use JSON or XML!")
	}

}
