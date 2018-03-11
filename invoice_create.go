package gofastbill

import (
	"errors"
	//	"log"
)

type InvoiceCreate_Request struct {
	CUSTOMER_ID            string  `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       //Required	Eine bestimmte Kundennummer
	CUSTOMER_COSTCENTER_ID string  `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //ID der Kostenstelle
	CURRENCY_CODE          string  `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //Standardwährung
	TEMPLATE_ID            string  `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                       //Entwurfsnummer
	INTROTEXT              string  `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                           //Einleitungstext
	INVOICE_TITLE          string  `xml:"INVOICE_TITLE,omitempty" json:"INVOICE_TITLE,omitempty"`                   //Rechnungstitel
	INVOICE_DATE           string  `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`                     //Rechnungsdatum
	DELIVERY_DATE          string  `xml:"DELIVERY_DATE,omitempty" json:"DELIVERY_DATE,omitempty"`                   //Lieferdatum
	CASH_DISCOUNT_PERCENT  string  `xml:"CASH_DISCOUNT_PERCENT,omitempty" json:"CASH_DISCOUNT_PERCENT,omitempty"`   //Skonto in Prozent
	CASH_DISCOUNT_DAYS     string  `xml:"CASH_DISCOUNT_DAYS,omitempty" json:"CASH_DISCOUNT_DAYS,omitempty"`         //Skonto-Zeitraum in Tagen
	EU_DELIVERY            string  `xml:"EU_DELIVERY,omitempty" json:"EU_DELIVERY,omitempty"`                       //Flag für Anzeige einer Innergemeinschaftliche Lieferung: 0 = nein | 1 = ja
	IS_GROSS               string  `xml:"IS_GROSS,omitempty" json:"IS_GROSS,omitempty"`                             //Flag ob auf Brutto umgestellt werden soll: 0 = no | 1 = yes
	TEMPLATE_HASH          string  `xml:"TEMPLATE_HASH,omitempty" json:"TEMPLATE_HASH,omitempty"`                   //Eindeutige ID des Templates
	ITEMS                  []ITEMS `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //Required	Liste der Artikel

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
