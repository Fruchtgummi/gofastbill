package gofastbill

import (
	"errors"
)

type InvoiceComplete_Request struct {
	INVOICE_ID string `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`
}

//Delete Invoice; NEED: INVOICE_ID RESPONSE: "STATUS", "INVOICE_NUMBER"
func (s *Initialization) Invoice_complete(req InvoiceComplete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "invoice.complete"

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
