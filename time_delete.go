package gofastbill

import (
	"errors"
)

type TimeDelete_Request struct {
	TIME_ID string `xml:"TIME_ID,omitempty" json:"TIME_ID,omitempty"` //	TIME-ID
}

//DELETE Time
//FILTER => All fields from Struct: gofastbill.TimeDelete_Request
//REQUIRED: TIME_ID
//RESPONSE: STATUS
func (s *Initialization) Time_delete(req TimeDelete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "time.delete"

	if req.TIME_ID == "" {
		return nil, errors.New(s.Typ + ": TIME_ID must not be empty")
	}

	r.DATA.TIME_ID = req.TIME_ID

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
