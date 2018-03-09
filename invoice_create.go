package gofastbill

import (
	"errors"
	//	"log"
)

type InvoiceCreate_Request struct {
	CUSTOMER_ID            string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Required	Eine bestimmte Kundennummer
	CUSTOMER_COSTCENTER_ID string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //ID der Kostenstelle
	CURRENCY_CODE          string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Standardwährung
	TEMPLATE_ID            string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Entwurfsnummer
	INTROTEXT              string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Einleitungstext
	INVOICE_TITLE          string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Rechnungstitel
	INVOICE_DATE           string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Rechnungsdatum
	DELIVERY_DATE          string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Lieferdatum
	CASH_DISCOUNT_PERCENT  string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Skonto in Prozent
	CASH_DISCOUNT_DAYS     string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Skonto-Zeitraum in Tagen
	EU_DELIVERY            string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Flag für Anzeige einer Innergemeinschaftliche Lieferung: 0 = nein | 1 = ja
	IS_GROSS               string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Flag ob auf Brutto umgestellt werden soll: 0 = no | 1 = yes
	TEMPLATE_HASH          string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Eindeutige ID des Templates
	ITEMS                  []ITEMS `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"` //Required	Liste der Artikel

}

//Create a new Invoice.
//FILTER => All fields from Struct: gofastbill.InvoiceCreate_Request
//Required: CUSTOMER_ID, []ITEMS
func (s *Initialization) Invoice_create(req InvoiceCreate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "invoice.create"

	if req.CUSTOMER_ID == "" || len(req.ITEMS) <= 0 {

		return nil, errors.New(s.Typ + ": Missing CUSTOMER_ID or ITEMS is empty")
	}

	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.CUSTOMER_COSTCENTER_ID = req.CUSTOMER_COSTCENTER_ID
	r.DATA.CURRENCY_CODE = req.CURRENCY_CODE
	r.DATA.TEMPLATE_ID = req.TEMPLATE_ID
	r.DATA.INTROTEXT = req.INTROTEXT
	r.DATA.INVOICE_TITLE = req.INVOICE_TITLE
	r.DATA.INVOICE_DATE = req.INVOICE_DATE
	r.DATA.DELIVERY_DATE = req.DELIVERY_DATE
	r.DATA.CASH_DISCOUNT_PERCENT = req.CASH_DISCOUNT_PERCENT
	r.DATA.CASH_DISCOUNT_DAYS = req.CASH_DISCOUNT_DAYS
	r.DATA.EU_DELIVERY = req.EU_DELIVERY
	r.DATA.IS_GROSS = req.IS_GROSS
	r.DATA.TEMPLATE_HASH = req.TEMPLATE_HASH
	r.DATA.ITEMS = req.ITEMS

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
