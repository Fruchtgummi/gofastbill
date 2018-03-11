package gofastbill

import (
	"errors"
)

type ProjectUpdate_Request struct {
	PROJECT_ID             string `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`                         // Required	Projektname
	PROJECT_NAME           string `xml:"PROJECT_NAME,omitempty" json:"PROJECT_NAME,omitempty"`                     // Required	Projektname
	CUSTOMER_ID            string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       // Required	Eine bestimmte Kundennummer
	CUSTOMER_COSTCENTER_ID string `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //	ID der Kostenstelle
	HOUR_PRICE             string `xml:"HOUR_PRICE,omitempty" json:"HOUR_PRICE,omitempty"`                         //	Stundensatz
	CURRENCY_CODE          string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //	Standardwährung
	VAT_PERCENT            string `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`                       //	MwSt in Prozent
	START_DATE             string `xml:"START_DATE,omitempty" json:"START_DATE,omitempty"`                         //	Datum des ersten Rechnungslaufs
	END_DATE               string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"`                             //	Enddatum
}

//Update a Project.
//FILTER => All fields from Struct: gofastbill.ProjectUpdate_Request
//Required: CUSTOMER_ID, PRJOJECT_NAME
//RESPONSE:
//STATUS
func (s *Initialization) Project_update(req ProjectUpdate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	if req.PROJECT_ID == "" {

		return nil, errors.New(s.Typ + ": Missing CUSTOMER_ID ")
	}

	r.SERVICE = "project.update"
	r.DATA.PROJECT_ID = req.PROJECT_ID
	r.DATA.PROJECT_NAME = req.PROJECT_NAME
	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.CUSTOMER_COSTCENTER_ID = req.CUSTOMER_COSTCENTER_ID
	r.DATA.HOUR_PRICE = req.HOUR_PRICE
	r.DATA.CURRENCY_CODE = req.CURRENCY_CODE
	r.DATA.VAT_PERCENT = req.VAT_PERCENT
	r.DATA.START_DATE = req.START_DATE
	r.DATA.END_DATE = req.END_DATE

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
