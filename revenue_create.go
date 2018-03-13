package gofastbill

import (
	"errors"
	//	"log"
)

type RevenueCreate_Request struct {
	INVOICE_DATE   string  `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`     // Required	Rechnungsdatum
	DUE_DATE       string  `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`             //	Fälligkeitsdatum
	CUSTOMER_ID    string  `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`       // Required	Eine bestimmte Kundennummer
	INVOICE_NUMBER string  `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"` //	Rechnungsnummer
	COMMENT        string  `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`               //
	SUB_TOTAL      float64 `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`           // Required	Nettobetrag
	VAT_TOTAL      float64 `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`           //	Vorsteuerbetrag
	CURRENCY_CODE  string  `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`   //	Standardwährung

}

//Create a new Revenue.
//FILTER => All fields from Struct: gofastbill.RevenueCreate_Request
//Required: INVOICE_DATE, CUSTOMER_ID, SUB_TOTAL
//RESPONSE: STATUS, INVOICE_ID
func (s *Initialization) Revenue_create(req RevenueCreate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "revenue.create"

	if req.INVOICE_DATE == "" {
		return nil, errors.New(s.Typ + ": Missing INVOICE_DATE is empty")
	}

	if req.CUSTOMER_ID == "" {
		return nil, errors.New(s.Typ + ": Missing CUSTOMER_ID is empty")
	}

	if req.SUB_TOTAL == 0 {
		return nil, errors.New(s.Typ + ": Missing SUB_TOTAL is empty")
	}

	r.DATA.INVOICE_DATE = req.INVOICE_DATE
	r.DATA.DUE_DATE = req.DUE_DATE
	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.INVOICE_NUMBER = req.INVOICE_NUMBER
	r.DATA.COMMENT = req.COMMENT
	r.DATA.SUB_TOTAL = req.SUB_TOTAL
	r.DATA.VAT_TOTAL = req.VAT_TOTAL
	r.DATA.CURRENCY_CODE = req.CURRENCY_CODE

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
