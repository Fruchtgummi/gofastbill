package gofastbill

import (
	"errors"
)

type ItemDelete_Request struct {
	INVOICE_ITEM_ID string `xml:"INVOICE_ITEM_ID,omitempty" json:"INVOICE_ITEM_ID,omitempty"`
}

//Delete Invoice; NEED: INVOICE_ITEM_ID
//RESPONSE: "STATUS"
func (s *Initialization) Item_delete(req ItemDelete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "item.delete"

	if req.INVOICE_ITEM_ID == "" {
		return nil, errors.New(s.Typ + ": INVOICE_ITEM_ID is empty")
	}

	r.DATA.INVOICE_ITEM_ID = req.INVOICE_ITEM_ID

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
