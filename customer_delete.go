package gofastbill

import (
	"errors"
)

type CustomerDelete_Request struct {
	CUSTOMER_ID string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
}

//Delete a custome
//FILTER => All fields from Struct: gofastbillCustomerDelete_Request
//Required: CUSTOMER_ID
//Response: STATUS
func (s *Initialization) Customer_delete(req CustomerDelete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "customer.delete"

	if req.CUSTOMER_ID == "" {
		return nil, errors.New(s.Typ + ": CUSTOMER_ID is empty")
	}

	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID

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
