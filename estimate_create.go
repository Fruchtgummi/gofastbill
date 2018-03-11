package gofastbill

import (
	"errors"
)

type EstimateCreate_Request struct {
	CUSTOMER_ID   string  `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	ITEMS         []ITEMS `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                 //
	TEMPLATE_NAME string  `xml:"TEMPLATE_NAME,omitempty" json:"TEMPLATE_NAME,omitempty"` //Vorlagenname
	TEMPLATE_HASH string  `xml:"TEMPLATE_HASH,omitempty" json:"TEMPLATE_HASH,omitempty"` //Eindeutige ID des Templates
}

//CREATE Estimate
//FILTER => All fields from Struct: gofastbill.EstimateCreate_Request
//Required:
//CUSTOMER_ID, ITEMS[]
//RESPONSE: STATUS, ESTIMATE_ID
func (s *Initialization) Estimate_create(req EstimateCreate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "estimate.create"

	if req.CUSTOMER_ID == "" {
		return nil, errors.New(s.Typ + ": CUSTOMER_ID must not be empty")
	}

	if len(req.ITEMS) <= 0 {
		return nil, errors.New(s.Typ + ": ITEMS must not be empty")
	}

	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.ITEMS = req.ITEMS
	r.DATA.TEMPLATE_NAME = req.TEMPLATE_NAME
	r.DATA.TEMPLATE_HASH = req.TEMPLATE_HASH

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
