package gofastbill

import (
	"errors"
)

type InvoiceSendbyemail_Request struct {
	INVOICE_ID           string    `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`                     //
	RECIPIENT            RECIPIENT `xml:"RECIPIENT,omitempty" json:"RECIPIENT,omitempty"`                       //Required	Empfänger
	SUBJECT              string    `xml:"SUBJECT,omitempty" json:"SUBJECT,omitempty"`                           //Betreff
	MESSAGE              string    `xml:"MESSAGE,omitempty" json:"MESSAGE,omitempty"`                           //e-Mail Text
	RECEIPT_CONFIRMATION string    `xml:"RECEIPT_CONFIRMATION,omitempty" json:"RECEIPT_CONFIRMATION,omitempty"` //

}

type RECIPIENT struct {
	TO  string `xml:"TO,omitempty" json:"TO,omitempty"`   //
	CC  string `xml:"CC,omitempty" json:"CC,omitempty"`   //
	BCC string `xml:"BCC,omitempty" json:"BCC,omitempty"` //
}

//Sendbyemail Invoice; NEED: INVOICE_ID, RECIPIENT{TO,CC,BCC}, SUBJECT, MESSAGE, RECEIPT_CONFIRMATION; RESPONSE: "STATUS", "INVOICE_NUMBER"
func (s *Initialization) Invoice_sendbyemail(req InvoiceSendbyemail_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "invoice.sendbyemail"

	if req.INVOICE_ID == "" {
		return nil, errors.New(s.Typ + ": INVOICE_ID is empty")
	}

	if req.RECIPIENT.TO == "" {
		return nil, errors.New(s.Typ + ": You need RECIPIENT at least TO")
	}

	r.DATA.INVOICE_ID = req.INVOICE_ID
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
