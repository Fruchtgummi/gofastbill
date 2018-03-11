package gofastbill

import (
	"errors"
	//	"log"
)

type RecurringUpdate_Request struct {
	INVOICE_ID             string  `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`                         //Required Required	Rechnungs-ID
	DELETE_EXISTING_ITEMS  string  `xml:"DELETE_EXISTING_ITEMS,omitempty" json:"DELETE_EXISTING_ITEMS,omitempty"`   //Required	Flag zur Löschung aller bestehenden Rechnungsposten: 0 = nein | 1 = ja
	CUSTOMER_ID            string  `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       //Required	Eine bestimmte Kundennummer
	CUSTOMER_COSTCENTER_ID string  `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //ID der Kostenstelle
	CURRENCY_CODE          string  `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //Standardwährung
	TEMPLATE_ID            string  `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                       //Entwurfsnummer
	INTROTEXT              string  `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                           //Einleitungstext
	START_DATE             string  `xml:"START_DATE,omitempty" json:"START_DATE,omitempty"`                         // Required	Datum des ersten Rechnungslaufs
	FREQUENCY              string  `xml:"FREQUENCY,omitempty" json:"FREQUENCY,omitempty"`                           // Required	Wiederholrate des Rechnungslauf: Weekly = wöchentlich | 2 weeks = alle 2 Wochen | 4 weeks = alle 4 Wochen | monthly = monatlich | 2 months = alle 2 Monate | 3 months = vierteljährlich | 6 months = alle 6 Monate| yearly = jährlich| 2 years = alle 2 Jahre
	OCCURENCES             string  `xml:"OCCURENCES,omitempty" json:"OCCURENCES,omitempty"`                         //	Anzahl der wiederkehrenden Rechnungen: 0 = unbegrenzt
	OUTPUT_TYPE            string  `xml:"OUTPUT_TYPE,omitempty" json:"OUTPUT_TYPE,omitempty"`                       // Required	Typ der automatisch zu erzeugenden Rechnung: draft = Entwurf | Outgoing = Rechnung
	EMAIL_NOTIFY           string  `xml:"EMAIL_NOTIFY,omitempty" json:"EMAIL_NOTIFY,omitempty"`                     //	Flag für den Versand an die eigene e-Mailadresse: 0 = nein | 1 = ja
	DELIVERY_DATE          string  `xml:"DELIVERY_DATE,omitempty" json:"DELIVERY_DATE,omitempty"`                   //	Lieferdatum
	CASH_DISCOUNT_PERCENT  string  `xml:"CASH_DISCOUNT_PERCENT,omitempty" json:"CASH_DISCOUNT_PERCENT,omitempty"`   //	Skonto in Prozent
	CASH_DISCOUNT_DAYS     string  `xml:"CASH_DISCOUNT_DAYS,omitempty" json:"CASH_DISCOUNT_DAYS,omitempty"`         //	Skonto-Zeitraum in Tagen
	EU_DELIVERY            string  `xml:"EU_DELIVERY,omitempty" json:"EU_DELIVERY,omitempty"`                       //Flag für Anzeige einer Innergemeinschaftliche Lieferung: 0 = nein | 1 = ja
	TEMPLATE_HASH          string  `xml:"TEMPLATE_HASH,omitempty" json:"TEMPLATE_HASH,omitempty"`                   //Eindeutige ID des Templates
	ITEMS                  []ITEMS `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //Required	Liste der Artikel

}

//Update a Recurring.
//FILTER => All fields from Struct: gofastbill.RecurringUpdate_Request
//Required: INVOICE_ID, DELETE_EXISTING_ITEMS
//RESPONSE: STATUS
func (s *Initialization) Recurring_update(req RecurringUpdate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "recurring.update"

	if req.INVOICE_ID == "" {
		return nil, errors.New(s.Typ + ": Missing INVOICE_ID is empty")
	}

	if req.DELETE_EXISTING_ITEMS == "" {
		return nil, errors.New(s.Typ + ": Missing DELETE_EXISTING_ITEMS is empty (0|1)")
	}

	r.DATA.INVOICE_ID = req.INVOICE_ID
	r.DATA.DELETE_EXISTING_ITEMS = req.DELETE_EXISTING_ITEMS
	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.CUSTOMER_COSTCENTER_ID = req.CUSTOMER_COSTCENTER_ID
	r.DATA.CURRENCY_CODE = req.CURRENCY_CODE
	r.DATA.TEMPLATE_ID = req.TEMPLATE_ID
	r.DATA.INTROTEXT = req.INTROTEXT
	r.DATA.START_DATE = req.START_DATE
	r.DATA.FREQUENCY = req.FREQUENCY
	r.DATA.OCCURENCES = req.OCCURENCES
	r.DATA.OUTPUT_TYPE = req.OUTPUT_TYPE
	r.DATA.EMAIL_NOTIFY = req.EMAIL_NOTIFY
	r.DATA.DELIVERY_DATE = req.DELIVERY_DATE
	r.DATA.CASH_DISCOUNT_PERCENT = req.CASH_DISCOUNT_PERCENT
	r.DATA.CASH_DISCOUNT_DAYS = req.CASH_DISCOUNT_DAYS
	r.DATA.EU_DELIVERY = req.EU_DELIVERY
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
