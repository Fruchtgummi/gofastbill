package gofastbill

import (
	"errors"
)

type EstimateCreateinvoice_Request struct {
	ESTIMATE_ID string `xml:"ESTIMATE_ID,omitempty" json:"ESTIMATE_ID,omitempty"` //	Angebots-ID
}

//CREATINVOICE Estimate
//FILTER => All fields from Struct: gofastbill.EstimateCreateinvoice_Request
//REQUIRED: ESTIMATE_ID
//RESPONSE: STATUS
func (s *Initialization) Estimate_createinvoicet(req EstimateCreateinvoice_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "estimate.createinvoice"

	if req.ESTIMATE_ID == "" {
		return nil, errors.New(s.Typ + ": ESTIMATE_ID must not be empty")
	}

	r.DATA.ESTIMATE_ID = req.ESTIMATE_ID

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
