package gofastbill

import (
	"errors"
)

type ExpenseCreate_Request struct {
	INVOICE_DATE   string `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`     // Required	Rechnungsdatum
	DUE_DATE       string `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`             //	Fälligkeitsdatum
	PROJECT_ID     string `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`         //	Eine bestimmte Projekt ID
	CUSTOMER_ID    string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`       //	Eine bestimmte Kundennummer
	ORGANIZATION   string `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`     // Required	Firmenname
	INVOICE_NUMBER string `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"` //	Rechnungsnummer
	COMMENT        string `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`               //	Bemerkung
	SUB_TOTAL      string `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`           // Required	Nettobetrag
	VAT_TOTAL      string `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`           //	Vorsteuerbetrag
}

//CREATE Estimate
//FILTER => All fields from Struct: gofastbill.ExpenseCreate_Request
//Required:
//CUSTOMER_ID, ITEMS[]
//RESPONSE: STATUS, ESTIMATE_ID
func (s *Initialization) Expense_create(req ExpenseCreate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "expense.create"

	if req.INVOICE_DATE == "" {
		return nil, errors.New(s.Typ + ": INVOICE_DATE must not be empty")
	}

	if req.ORGANIZATION == "" {
		return nil, errors.New(s.Typ + ": ORGANIZATION must not be empty")
	}

	if req.SUB_TOTAL == "" {
		return nil, errors.New(s.Typ + ": SUB_TOTAL must not be empty")
	}

	r.DATA.INVOICE_DATE = req.INVOICE_DATE
	r.DATA.DUE_DATE = req.DUE_DATE
	r.DATA.PROJECT_ID = req.PROJECT_ID
	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.ORGANIZATION = req.ORGANIZATION
	r.DATA.INVOICE_NUMBER = req.INVOICE_NUMBER
	r.DATA.COMMENT = req.COMMENT
	r.DATA.SUB_TOTAL = req.SUB_TOTAL
	r.DATA.VAT_TOTAL = req.VAT_TOTAL

	//TODO: Beleg als Datei Multipartform

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

//Aufbau des POST-Requests im Multipart-Format:

//curl -v -X POST \
//–u {E-Mail-Adresse}:{API-Key} \
//-H 'Content-Type: multipart/form-data, boundary=AaB03x' \
//-d '{xml body}' \ https://my.fastbill.com/api/1.0/api.php

//POST /api/1.0/api.php HTTP/1.0
//Host: my.fastbill.com
//Content-type: multipart/form-data, boundary=AaB03x
//Content-Length: {Größe des Inhalts}

//--AaB03x content-disposition: form-data; name="httpbody"
//{XML- oder JSON-Daten}
//--AaB03x
//content-disposition: form-data; name="document"; filename{Dateiname}"
//Content-Type: {MIME Typ}
//Content-Transfer-Encoding: binary
//{Datei-Inhalt}
//--AaB03x--
