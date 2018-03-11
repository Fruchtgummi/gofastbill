package gofastbill

import (
	"errors"
)

type ItemGet_Request struct {
	INVOICE_ID string `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`
}

//Delete Invoice; NEED: INVOICE_ID
//RESPONSE: all field of gofastbill.Item
func (s *Initialization) Item_get(req ItemGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "item.get"

	if req.INVOICE_ID == "" {
		return nil, errors.New(s.Typ + ": INVOICE_ID is empty")
	}

	r.DATA.INVOICE_ID = req.INVOICE_ID

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
