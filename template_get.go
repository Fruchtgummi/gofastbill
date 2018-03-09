package gofastbill

type TemplateGet_Request struct {
}

//GET Template
//FILTER => NO FILTERS

func (s *Initialization) Template_get(req TemplateGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "template.get"

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
