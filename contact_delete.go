package gofastbill

import (
	"errors"
)

type ContactDelete_Request struct {
	CUSTOMER_ID string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CONTACT_ID  string `xml:"CONTACT_ID,omitempty" json:"CONTACT_ID,omitempty"`
}

//Delete Customer; NEED: CUSTOMER_ID, CONTACT_ID; RESPONSE: "STATUS"
func (s *Initialization) Contact_delete(req ContactDelete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "contact.delete"

	if req.CUSTOMER_ID == "" {
		return nil, errors.New(s.Typ + ": CUSTOMER_ID is empty")
	}

	if req.CONTACT_ID == "" {
		return nil, errors.New(s.Typ + ": CONTACT_ID is empty")
	}

	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.CONTACT_ID = req.CONTACT_ID

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
