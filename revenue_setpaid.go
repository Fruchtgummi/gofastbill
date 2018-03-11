package gofastbill

import (
	"errors"
)

type RevenueSetpaid_Request struct {
	INVOICE_ID string `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`
	PAID_DATE  string `xml:"PAID_DATE,omitempty" json:"PAID_DATE,omitempty"` //Datum der Zahlunf
}

//Setpaid Revenue; NEED: INVOICE_ID, PAID_DATE; RESPONSE: "STATUS", "INVOICE_NUMBER"
func (s *Initialization) Revenue_setpaid(req RevenueSetpaid_Request) (*FBAPI, error) {

	var fastbillbody string
	var r FBAPI

	r.SERVICE = "revenue.setpaid"

	if req.INVOICE_ID == "" {
		return nil, errors.New(s.Typ + ": INVOICE_ID is empty")
	}

	if req.PAID_DATE == "" {
		return nil, errors.New(s.Typ + ": PAID_DATE is empty")
	}

	r.DATA.INVOICE_ID = req.INVOICE_ID
	r.DATA.PAID_DATE = req.PAID_DATE

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
