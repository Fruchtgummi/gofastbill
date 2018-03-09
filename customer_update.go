package gofastbill

import (
	"errors"
	//	"log"
)

type CustomerUpdate_Request struct {
	CUSTOMER_ID                    string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CUSTOMER_NUMBER                string `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`                               //	Eigene Kundennummer
	CUSTOMER_TYPE                  string `xml:"CUSTOMER_TYPE,omitempty" json:"CUSTOMER_TYPE,omitempty"`                                   // 	Required	Kundentyp: business = Geschäftskunde | consumer = Privatperson
	ORGANIZATION                   string `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`                                     // 	Required	Firmenname [REQUIRED] wenn CUSTOMER_TYPE = business
	POSITION                       string `xml:"POSITION,omitempty" json:"POSITION,omitempty"`                                             // 	Position in der Firma
	ACADEMIC_DEGREE                string `xml:"ACADEMIC_DEGREE,omitempty" json:"ACADEMIC_DEGREE,omitempty"`                               //	Akademischer Grad
	SALUTATION                     string `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`                                         //	Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME                     string `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`                                         //	Vorname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	LAST_NAME                      string `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                                           // 	Required	Nachname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	ADDRESS                        string `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                                               //	Adresszeile 1
	ADDRESS_2                      string `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                                           //	Adresszeile 2
	SECONDARY_ADDRESS              string `xml:"SECONDARY_ADDRESS,omitempty" json:"SECONDARY_ADDRESS,omitempty"`                           //	Lieferadresse
	ZIPCODE                        string `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                                               //	Postleitzahl
	CITY                           string `xml:"CITY,omitempty" json:"CITY,omitempty"`                                                     //	Stadt
	STATE                          string `xml:"STATE,omitempty" json:"STATE,omitempty"`                                                   //	Bundestaat
	COUNTRY_CODE                   string `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`                                     //	Länder-Code (ISO 3166 ALPHA-2)
	PHONE                          string `xml:"PHONE,omitempty" json:"PHONE,omitempty"`                                                   //	Telefon
	PHONE_2                        string `xml:"PHONE_2,omitempty" json:"PHONE_2,omitempty"`                                               //	Telefon 2
	FAX                            string `xml:"FAX,omitempty" json:"FAX,omitempty"`                                                       //	Faxnummer
	MOBILE                         string `xml:"MOBILE,omitempty" json:"MOBILE,omitempty"`                                                 //	Mobiltelefonnummer
	EMAIL                          string `xml:"EMAIL,omitempty" json:"EMAIL,omitempty"`                                                   //	E-Mailadresse
	WEBSITE                        string `xml:"WEBSITE,omitempty" json:"WEBSITE,omitempty"`                                               //	Webseite
	ACCOUNT_RECEIVABLE             string `xml:"ACCOUNT_RECEIVABLE,omitempty" json:"ACCOUNT_RECEIVABLE,omitempty"`                         //	Datev Debitoren-Kontonummer
	CURRENCY_CODE                  string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                                   //	Standardwährung
	VAT_ID                         string `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`                                                 //	USt-IdNr.
	DAYS_FOR_PAYMENT               string `xml:"DAYS_FOR_PAYMENT,omitempty" json:"DAYS_FOR_PAYMENT,omitempty"`                             //	Tage bis zum Zahlungsziel
	PAYMENT_TYPE                   string `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`                                     //	1 = ueberW | 2 = lastS | 3 = bar | 4 = paypal | 5 = vorK | 6 = kreditK
	SHOW_PAYMENT_NOTICE            string `xml:"SHOW_PAYMENT_NOTICE,omitempty" json:"SHOW_PAYMENT_NOTICE,omitempty"`                       //	Zahlungshinweis anzeigen
	BANK_NAME                      string `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`                                           //	Bankname
	BANK_CODE                      string `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`                                           //	Bankleitzahl
	BANK_ACCOUNT_NUMBER            string `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`                       //	Kontonummer
	BANK_ACCOUNT_OWNER             string `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`                         //	Kontoinhaber
	BANK_ACCOUNT_MANDATE_REFERENCE string `xml:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty" json:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty"` //	Mandatsrefernznummer
	TAGS                           string `xml:"TAGS,omitempty" json:"TAGS,omitempty"`                                                     //	Tag halt
}

//Update a customer.
//FILTER => All fields from Struct: gofastbill.CustomerUpdate_Request
//Required:
//CUSTOME_ID

func (s *Initialization) Customer_update(req CustomerUpdate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "customer.update"

	if req.CUSTOMER_TYPE == "business" || req.CUSTOMER_TYPE == "consumer" {

		if req.CUSTOMER_TYPE == "business" {
			if req.ORGANIZATION == "" {

				return nil, errors.New(s.Typ + ": ORGANIZATION must not be empty if CUSTOMER_TYPE business")
			}
		} else if req.CUSTOMER_TYPE == "consumer" {
			if req.LAST_NAME == "" {
				return nil, errors.New(s.Typ + ": LAST_NAME must not be empty if COSTUMER_TYPE consumer")
			}
		}
	} else {
		return nil, errors.New(s.Typ + ": Wrong CUSTOMER_TYPE do you use business|consumer")
	}

	if req.CUSTOMER_ID == "" {
		return nil, errors.New(s.Typ + ": Missing CUSTOMER_ID, is required for update")
	}

	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.CUSTOMER_NUMBER = req.CUSTOMER_NUMBER
	r.DATA.CUSTOMER_TYPE = req.CUSTOMER_TYPE
	r.DATA.ORGANIZATION = req.ORGANIZATION
	r.DATA.POSITION = req.POSITION
	r.DATA.ACADEMIC_DEGREE = req.ACADEMIC_DEGREE
	r.DATA.SALUTATION = req.SALUTATION
	r.DATA.FIRST_NAME = req.FIRST_NAME
	r.DATA.LAST_NAME = req.LAST_NAME
	r.DATA.ADDRESS = req.ADDRESS
	r.DATA.ADDRESS_2 = req.ADDRESS_2
	r.DATA.SECONDARY_ADDRESS = req.SECONDARY_ADDRESS
	r.DATA.ZIPCODE = req.ZIPCODE
	r.DATA.CITY = req.CITY
	r.DATA.STATE = req.STATE
	r.DATA.COUNTRY_CODE = req.COUNTRY_CODE
	r.DATA.PHONE = req.PHONE
	r.DATA.PHONE_2 = req.PHONE_2
	r.DATA.FAX = req.FAX
	r.DATA.MOBILE = req.MOBILE
	r.DATA.EMAIL = req.EMAIL
	r.DATA.WEBSITE = req.WEBSITE
	r.DATA.ACCOUNT_RECEIVABLE = req.ACCOUNT_RECEIVABLE
	r.DATA.CURRENCY_CODE = req.CURRENCY_CODE
	r.DATA.VAT_ID = req.VAT_ID
	r.DATA.DAYS_FOR_PAYMENT = req.DAYS_FOR_PAYMENT
	r.DATA.PAYMENT_TYPE = req.PAYMENT_TYPE
	r.DATA.SHOW_PAYMENT_NOTICE = req.SHOW_PAYMENT_NOTICE
	r.DATA.BANK_NAME = req.BANK_NAME
	r.DATA.BANK_CODE = req.BANK_CODE
	r.DATA.BANK_ACCOUNT_NUMBER = req.BANK_ACCOUNT_NUMBER
	r.DATA.BANK_ACCOUNT_OWNER = req.BANK_ACCOUNT_OWNER
	r.DATA.BANK_ACCOUNT_MANDATE_REFERENCE = req.BANK_ACCOUNT_MANDATE_REFERENCE
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
