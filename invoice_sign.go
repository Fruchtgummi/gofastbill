package gofastbill

import (
	"errors"
)

type InvoiceSign_Request struct {
	INVOICE_ID string `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`
}

//Sign Invoice; NEED: INVOICE_ID RESPONSE: "STATUS", "REMAINING_CREDITS"
func (s *Initialization) Invoice_sign(req InvoiceSign_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "invoice.sign"

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
