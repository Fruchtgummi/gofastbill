package gofastbill

import (
	"errors"
	"reflect"
)

type DocumentCreate_Request struct {
	TYPE  string `xml:"TYPE,omitempty" json:"TYPE,omitempty"`   //	Dokumententyp
	TITLE string `xml:"TITLE,omitempty" json:"TITLE,omitempty"` //	Titel
	DATE  string `xml:"DATE,omitempty" json:"DATE,omitempty"`   //	Datum
	NOTE  string `xml:"NOTE,omitempty" json:"NOTE,omitempty"`   //	Bemerkung
}

//CREATE Document
//FILTER => All fields from Struct: gofastbill.DocumentCreate_Request
//RESPONSE: STAUS, DOCUMENT_ID
//REQUIRE: TYPE of file to upload, pdf|png
//if you to do upload a file then set s.Upload true (cooming soon)
func (s *Initialization) Document_create(req DocumentCreate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "document.create"

	if req.TYPE == "" {
		return nil, errors.New(s.Typ + ": TYPE must not be empty (PDF, PNG)")
	} else if !checkType(req.TYPE) {
		return nil, errors.New(s.Typ + ": Wrong TYPE, possible are (PDF, PNG)")
	}

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

func checkType(s string) bool {
	b := [...]string{"pdf", "png"}

	return in_array(s, b)

}

func in_array(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {

				exists = true
				return
			}
		}
	}

	return
}
