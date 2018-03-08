package gofastbill

import (
	"errors"
	//	"log"
)

type ContactCreate_Request struct {
	CUSTOMER_ID       string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`             //	Eigene Kundennummer
	ORGANIZATION      string `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`           // 	Required	Firmenname [REQUIRED] wenn CUSTOMER_TYPE = business
	POSITION          string `xml:"POSITION,omitempty" json:"POSITION,omitempty"`                   // 	Position in der Firma
	SALUTATION        string `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`               //	Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME        string `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`               //	Vorname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	LAST_NAME         string `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                 // 	Required	Nachname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	ADDRESS           string `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                     //	Adresszeile 1
	ADDRESS_2         string `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                 //	Adresszeile 2
	SECONDARY_ADDRESS string `xml:"SECONDARY_ADDRESS,omitempty" json:"SECONDARY_ADDRESS,omitempty"` //	Lieferadresse
	ZIPCODE           string `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                     //	Postleitzahl
	CITY              string `xml:"CITY,omitempty" json:"CITY,omitempty"`                           //	Stadt
	COUNTRY_CODE      string `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`           //	Länder-Code (ISO 3166 ALPHA-2)
	PHONE             string `xml:"PHONE,omitempty" json:"PHONE,omitempty"`                         //	Telefon
	PHONE_2           string `xml:"PHONE_2,omitempty" json:"PHONE_2,omitempty"`                     //	Telefon 2
	FAX               string `xml:"FAX,omitempty" json:"FAX,omitempty"`                             //	Faxnummer
	MOBILE            string `xml:"MOBILE,omitempty" json:"MOBILE,omitempty"`                       //	Mobiltelefonnummer
	EMAIL             string `xml:"EMAIL,omitempty" json:"EMAIL,omitempty"`                         //	E-Mailadresse
	CURRENCY_CODE     string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`         //	Standardwährung
	VAT_ID            string `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`
	CREATED           string `xml:"CREATED,omitempty" json:"CREATED,omitempty"` //	USt-IdNr.
	TAGS              string `xml:"TAGS,omitempty" json:"TAGS,omitempty"`       //	Tag halt
}

// Create customer; RETURN "STATUS, "CONTACT_ID"
func (s *Initialization) Contact_Create(req ContactCreate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "contact.create"

	if req.CUSTOMER_ID == "" {

		return nil, errors.New(s.Typ + ": CUSTOMER_ID must not be empty")

	}

	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.ORGANIZATION = req.ORGANIZATION
	r.DATA.POSITION = req.POSITION
	r.DATA.SALUTATION = req.SALUTATION
	r.DATA.FIRST_NAME = req.FIRST_NAME
	r.DATA.LAST_NAME = req.LAST_NAME
	r.DATA.ADDRESS = req.ADDRESS
	r.DATA.ADDRESS_2 = req.ADDRESS_2
	r.DATA.SECONDARY_ADDRESS = req.SECONDARY_ADDRESS
	r.DATA.ZIPCODE = req.ZIPCODE
	r.DATA.CITY = req.CITY
	r.DATA.COUNTRY_CODE = req.COUNTRY_CODE
	r.DATA.PHONE = req.PHONE
	r.DATA.PHONE_2 = req.PHONE_2
	r.DATA.FAX = req.FAX
	r.DATA.MOBILE = req.MOBILE
	r.DATA.EMAIL = req.EMAIL
	r.DATA.CURRENCY_CODE = req.CURRENCY_CODE
	r.DATA.VAT_ID = req.VAT_ID
	r.DATA.CREATED = req.CREATED
	r.DATA.TAGS = req.TAGS

	fastbillbody, err := s.GenerateRequest(r)
	if err != nil {
		return nil, err
	}

	resp, err := s.FastbillRequest(fastbillbody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
