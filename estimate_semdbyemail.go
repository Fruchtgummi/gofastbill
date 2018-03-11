package gofastbill

import (
	"errors"
)

type EstimateSendbyemail_Request struct {
	ESTIMATE_ID          string    `xml:"ESTIMATE_ID,omitempty" json:"ESTIMATE_ID,omitempty"`                   //
	RECIPIENT            RECIPIENT `xml:"RECIPIENT,omitempty" json:"RECIPIENT,omitempty"`                       //Required	Empfänger
	SUBJECT              string    `xml:"SUBJECT,omitempty" json:"SUBJECT,omitempty"`                           //Betreff
	MESSAGE              string    `xml:"MESSAGE,omitempty" json:"MESSAGE,omitempty"`                           //e-Mail Text
	RECEIPT_CONFIRMATION string    `xml:"RECEIPT_CONFIRMATION,omitempty" json:"RECEIPT_CONFIRMATION,omitempty"` //

}

//Sendbyemail Estimate; NEED: INVOICE_ID, RECIPIENT{TO,CC,BCC}, SUBJECT, MESSAGE, RECEIPT_CONFIRMATION;
//RESPONSE: "STATUS", "INVOICE_NUMBER"
func (s *Initialization) Estimate_sendbyemail(req EstimateSendbyemail_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "Estimate.sendbyemail"

	if req.ESTIMATE_ID == "" {
		return nil, errors.New(s.Typ + ": ESTIMATE_ID is empty")
	}

	if req.RECIPIENT.TO == "" {
		return nil, errors.New(s.Typ + ": You need RECIPIENT at least TO")
	}

	r.DATA.ESTIMATE_ID = req.ESTIMATE_ID
	r.DATA.SUBJECT = req.SUBJECT
	r.DATA.MESSAGE = req.MESSAGE
	r.DATA.RECEIPT_CONFIRMATION = req.RECEIPT_CONFIRMATION
	r.DATA.RECIPIENT = req.RECIPIENT

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
