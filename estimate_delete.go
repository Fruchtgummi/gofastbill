package gofastbill

import (
	"errors"
)

type EstimateDelete_Request struct {
	ESTIMATE_ID string `xml:"ESTIMATE_ID,omitempty" json:"ESTIMATE_ID,omitempty"` //	Angebots-ID
}

//DELETE Estimate
//FILTER => All fields from Struct: gofastbill.EstimateDelete_Request
//REQUIRED: ESTIMATE_ID
//RESPONSE: STATUS
func (s *Initialization) Estimate_delete(req EstimateDelete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "estimate.delete"

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
