package gofastbill

import (
	"errors"
)

type TimeUpdate_Request struct {
	TIME_ID          string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`           //
	CUSTOMER_ID      string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`           // Required	Eine bestimmte Kundennummer
	PROJECT_ID       string `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`             // Required	Eine bestimmte Projekt ID
	DATE             string `xml:"DATE,omitempty" json:"DATE,omitempty"`                         //	Datum
	TASK_ID          string `xml:"TASK_ID,omitempty" json:"TASK_ID,omitempty"`                   //	ID einer bestimmten Aufgabe
	START_TIME       string `xml:"START_TIME,omitempty" json:"START_TIME,omitempty"`             // Required	Startzeit
	END_TIME         string `xml:"END_TIME,omitempty" json:"END_TIME,omitempty"`                 //	Endzeit
	MINUTES          string `xml:"MINUTES,omitempty" json:"MINUTES,omitempty"`                   //	Minuten
	BILLABLE_MINUTES string `xml:"BILLABLE_MINUTES,omitempty" json:"BILLABLE_MINUTES,omitempty"` //	Abrechenbare Minuten
	COMMENT          string `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`                   //	Bemerkung
}

//UPDATE a Time.
//FILTER => All fields from Struct: gofastbill.TimeUpdate_Request
//Required: TIME_ID
//RESPONSE:
//COMMENT, TIME_ID
func (s *Initialization) Time_update(req TimeUpdate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "time.update"

	if req.TIME_ID == "" {

		return nil, errors.New(s.Typ + ": Missing TIME_ID ")
	}

	r.DATA.TIME_ID = req.TIME_ID
	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.PROJECT_ID = req.PROJECT_ID
	r.DATA.DATE = req.DATE
	r.DATA.TASK_ID = req.TASK_ID
	r.DATA.START_TIME = req.START_TIME
	r.DATA.END_TIME = req.END_TIME
	r.DATA.MINUTES = req.MINUTES
	r.DATA.BILLABLE_MINUTES = req.BILLABLE_MINUTES
	r.DATA.COMMENT = req.COMMENT

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
