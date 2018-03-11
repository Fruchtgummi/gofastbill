package gofastbill

import (
	"errors"
)

type ProjectDelete_Request struct {
	PROJECT_ID string `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"` // Required	Projektname

}

//Delete a Project.
//FILTER => All fields from Struct: gofastbill.ProjectDelete_Request
//Required: PROJECT_ID
//RESPONSE:
//STATUS
func (s *Initialization) Project_delete(req ProjectDelete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	if req.PROJECT_ID == "" {

		return nil, errors.New(s.Typ + ": Missing CUSTOMER_ID ")
	}

	r.SERVICE = "project.delete"
	r.DATA.PROJECT_ID = req.PROJECT_ID

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
