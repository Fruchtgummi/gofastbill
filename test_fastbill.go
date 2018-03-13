package gofastbill

import (
	"fmt"
	"log"

	"github.com/fatih/structs"
)

func (s *Initialization) Test_CustomerGet() {
	var x CustomerGet_Request

	fbapi, err := s.Customer_get(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_CustomerDelete(id string) {
	var x CustomerDelete_Request
	x.CUSTOMER_ID = id

	fbapi, err := s.Customer_delete(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_CustomerGetByID(id string) {
	var x CustomerGet_Request
	x.CUSTOMER_ID = id
	x.CITY = "leipzig"
	fbapi, err := s.Customer_get(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_CustomerCreateConsumer() {
	var x CustomerCreate_Request

	x.CUSTOMER_NUMBER = "123Kundennummer"
	x.CUSTOMER_TYPE = "consumer"
	x.ACADEMIC_DEGREE = "Dr. Prof."
	x.SALUTATION = "mrs"
	x.FIRST_NAME = "Joahana"
	x.LAST_NAME = "Doe"
	x.ADDRESS = "Wald und Wiesen Weg 2b"
	x.ADDRESS_2 = "Hinterhof"
	x.SECONDARY_ADDRESS = "Postfach 9911"
	x.ZIPCODE = "01234"
	x.CITY = "Gibtsgarnichtstadt"
	x.STATE = "Sachsen"
	x.COUNTRY_CODE = "DE"
	x.PHONE = "012345678989"
	x.PHONE_2 = "012345678989"
	x.FAX = "012345678989"
	x.MOBILE = "012345678989"
	x.EMAIL = "email@eamil.de"
	x.WEBSITE = "http://anydomain.com"
	x.ACCOUNT_RECEIVABLE = "Datev Debitoren"
	x.CURRENCY_CODE = "EUR"
	x.VAT_ID = "XX123456789"
	x.DAYS_FOR_PAYMENT = "14"
	x.PAYMENT_TYPE = "1"
	x.SHOW_PAYMENT_NOTICE = "Ich bin eine Payment-Notiz"
	x.BANK_NAME = "Flora & Fauna Bank"
	x.BANK_CODE = "1234"
	x.BANK_ACCOUNT_NUMBER = "12345"
	x.BANK_ACCOUNT_OWNER = "Johana Doe"
	x.BANK_ACCOUNT_MANDATE_REFERENCE = "123456"
	x.TAGS = "Softwareentwicklung, Datenbanken"

	fbapi, err := s.Customer_create(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_CustomerCreateBusiness() {
	var x CustomerCreate_Request

	x.CUSTOMER_NUMBER = "123Kundennummer"
	x.CUSTOMER_TYPE = "business"
	x.ORGANIZATION = "Johana Doe GmbH"
	x.POSITION = "CTO"
	x.ACADEMIC_DEGREE = "Dr. Dr."
	x.SALUTATION = "mrs"
	x.FIRST_NAME = "Joahana"
	x.LAST_NAME = "Doe"
	x.ADDRESS = "Wald und Wiesen Weg 2b"
	x.ADDRESS_2 = "Hinterhof"
	x.SECONDARY_ADDRESS = "Postfach 9911"
	x.ZIPCODE = "01234"
	x.CITY = "Gibtsgarnichtstadt"
	x.STATE = "Sachsen"
	x.COUNTRY_CODE = "DE"
	x.PHONE = "012345678989"
	x.PHONE_2 = "012345678989"
	x.FAX = "012345678989"
	x.MOBILE = "012345678989"
	x.EMAIL = "email@eamil.de"
	x.WEBSITE = "http://anydomain.com"
	x.ACCOUNT_RECEIVABLE = "Datev Debitoren"
	x.CURRENCY_CODE = "EUR"
	x.VAT_ID = "XX123456789"
	x.DAYS_FOR_PAYMENT = "14"
	x.PAYMENT_TYPE = "1"
	x.SHOW_PAYMENT_NOTICE = "Ich bin eine Payment-Notiz"
	x.BANK_NAME = "Flora & Fauna Bank"
	x.BANK_CODE = "1234"
	x.BANK_ACCOUNT_NUMBER = "12345"
	x.BANK_ACCOUNT_OWNER = "Johana Doe"
	x.BANK_ACCOUNT_MANDATE_REFERENCE = "123456"
	x.TAGS = "Softwareentwicklung, Datenbanken"

	fbapi, err := s.Customer_create(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_CustomerCreateConsumerUpdate(id string) {
	var x CustomerUpdate_Request

	x.CUSTOMER_ID = id
	x.CUSTOMER_NUMBER = "123Kundennummer"
	x.CUSTOMER_TYPE = "consumer"
	x.ACADEMIC_DEGREE = "Bachelor of Arts"
	x.SALUTATION = "mrs"
	x.FIRST_NAME = "Joahana"
	x.LAST_NAME = "Doe"
	x.ADDRESS = "Wald und Wiesen Weg 2b"
	x.ADDRESS_2 = "Hinterhof"
	x.SECONDARY_ADDRESS = "Postfach 9911"
	x.ZIPCODE = "01234"
	x.CITY = "Gibtsgarnichtstadt"
	x.STATE = "Sachsen"
	x.COUNTRY_CODE = "DE"
	x.PHONE = "012345678989"
	x.PHONE_2 = "012345678989"
	x.FAX = "012345678989"
	x.MOBILE = "012345678989"
	x.EMAIL = "email@eamil.de"
	x.WEBSITE = "http://anydomain.com"
	x.ACCOUNT_RECEIVABLE = "Datev Debitoren"
	x.CURRENCY_CODE = "EUR"
	x.VAT_ID = "XX123456789"
	x.DAYS_FOR_PAYMENT = "14"
	x.PAYMENT_TYPE = "1"
	x.SHOW_PAYMENT_NOTICE = "Ich bin eine Payment-Notiz"
	x.BANK_NAME = "Flora & Fauna Bank"
	x.BANK_CODE = "1234"
	x.BANK_ACCOUNT_NUMBER = "12345"
	x.BANK_ACCOUNT_OWNER = "Johana Doe"
	x.BANK_ACCOUNT_MANDATE_REFERENCE = "123456"
	x.TAGS = "Softwareentwicklung, Datenbanken"

	fbapi, err := s.Customer_update(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_CustomerCreateBusinessUpdate(id string) {
	var x CustomerUpdate_Request

	x.CUSTOMER_ID = id
	x.CUSTOMER_NUMBER = "123Kundennummer"
	x.CUSTOMER_TYPE = "business"
	x.ORGANIZATION = "Johana Doe GmbH"
	x.POSITION = "CTO"
	x.ACADEMIC_DEGREE = "Master of Arts"
	x.SALUTATION = "mrs"
	x.FIRST_NAME = "Joahana"
	x.LAST_NAME = "Doe"
	x.ADDRESS = "Wald und Wiesen Weg 2b"
	x.ADDRESS_2 = "Hinterhof"
	x.SECONDARY_ADDRESS = "Postfach 9911"
	x.ZIPCODE = "01234"
	x.CITY = "Gibtsgarnichtstadt"
	x.STATE = "Sachsen"
	x.COUNTRY_CODE = "DE"
	x.PHONE = "012345678989"
	x.PHONE_2 = "012345678989"
	x.FAX = "012345678989"
	x.MOBILE = "012345678989"
	x.EMAIL = "email@eamil.de"
	x.WEBSITE = "http://anydomain.com"
	x.ACCOUNT_RECEIVABLE = "Datev Debitoren"
	x.CURRENCY_CODE = "EUR"
	x.VAT_ID = "XX123456789"
	x.DAYS_FOR_PAYMENT = "14"
	x.PAYMENT_TYPE = "1"
	x.SHOW_PAYMENT_NOTICE = "Ich bin eine Payment-Notiz"
	x.BANK_NAME = "Flora & Fauna Bank"
	x.BANK_CODE = "1234"
	x.BANK_ACCOUNT_NUMBER = "12345"
	x.BANK_ACCOUNT_OWNER = "Johana Doe"
	x.BANK_ACCOUNT_MANDATE_REFERENCE = "123456"
	x.TAGS = "Softwareentwicklung, Datenbanken"

	fbapi, err := s.Customer_update(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_InvoiceGet() {
	var x InvoiceGet_Request

	fbapi, err := s.Invoice_get(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_InvoiceGetByID(id string) {
	var x InvoiceGet_Request
	x.INVOICE_ID = id
	fbapi, err := s.Invoice_get(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_InvoiceCreate() {
	var x InvoiceCreate_Request

	x.CUSTOMER_ID = "5353728"
	x.CUSTOMER_COSTCENTER_ID = ""
	x.CURRENCY_CODE = "EUR"
	x.TEMPLATE_ID = "1476010"
	x.INTROTEXT = `Wir bedanken uns für die gute Zusammenarbeit und stellen Ihnen vereinbarungsgemäß folgende
Lieferungen und Leistungen in Rechnung: 
`
	x.INVOICE_TITLE = "Teilnahme am Seminar"
	x.INVOICE_DATE = "2018-03-20 00:00:00"
	x.DELIVERY_DATE = "10.10.2017"
	x.CASH_DISCOUNT_PERCENT = "2.5"
	x.CASH_DISCOUNT_DAYS = "30"
	x.EU_DELIVERY = "0"
	x.IS_GROSS = "0"
	x.TEMPLATE_HASH = "300ae2b188116d2cab7c70ca6a451c04"

	var y ITEMS

	//loop

	y.ARTICLE_NUMBER = "1234567"
	y.DESCRIPTION = "I'm a description of ITEM"
	y.QUANTITY = "4"
	y.UNIT_PRICE = "25.00"
	y.VAT_PERCENT = "19"

	x.ITEMS = append(x.ITEMS, y)

	fbapi, err := s.Invoice_create(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

func (s *Initialization) Test_InvoiceUpdate(id string) {
	var x InvoiceUpdate_Request

	x.INVOICE_ID = id
	x.DELETE_EXISTING_ITEMS = "1"

	x.CUSTOMER_ID = "5353728"
	x.CUSTOMER_COSTCENTER_ID = ""
	x.CURRENCY_CODE = "EUR"
	x.TEMPLATE_ID = "1476010"
	x.INTROTEXT = `Wir bedanken uns für die gute Zusammenarbeit und stellen Ihnen vereinbarungsgemäß folgende
Lieferungen und Leistungen in Rechnung: 
`
	x.INVOICE_TITLE = "Teilnahme am Seminar Backen"
	x.INVOICE_DATE = "2018-03-20 00:00:00"
	x.DELIVERY_DATE = "10.10.2017"
	x.CASH_DISCOUNT_PERCENT = "2.5"
	x.CASH_DISCOUNT_DAYS = "30"
	x.EU_DELIVERY = "0"
	x.IS_GROSS = "0"
	x.TEMPLATE_HASH = "300ae2b188116d2cab7c70ca6a451c04"

	var y ITEMS

	//loop

	y.ARTICLE_NUMBER = "12345678"
	y.DESCRIPTION = "I'm a description of ITEM 2"
	y.QUANTITY = "4"
	y.UNIT_PRICE = "25.00"
	y.VAT_PERCENT = "19"

	x.ITEMS = append(x.ITEMS, y)

	fbapi, err := s.Invoice_update(x)

	if err != nil {
		log.Println(err)
	}

	parse(fbapi, "response")

}

//parser
func parse(fbapi *FBAPI, field string) {

	var g *structs.Struct

	switch field {
	case "response":
		g = structs.New(fbapi.RESPONSE)
	case "request":
		g = structs.New(fbapi.REQUEST)
	case "error":
		g = structs.New(fbapi.ERRORS)
	case "status":
		g = structs.New(fbapi.RESPONSE.STATUS)
	default:
		g = structs.New(fbapi)
	}

	for _, f := range g.Fields() {
		//fmt.Printf("field name: %+v\n", f.Name())

		if f.IsExported() && !f.IsZero() {
			fmt.Printf("value   : %+v\n", f.Value())

		}
	}

}
