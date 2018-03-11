package gofastbill

type ProjectGet_Request struct {
	CUSTOMER_ID string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` //
	PROJECT_ID  string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` // ID Eine bestimmte Projekt ID
}

//GET Project
//FILTER => CUSTOMER_ID, PJOJECT_ID, TASK_ID, TIME_ID, START_DATE, END_DATE, DATE
//Querying the details of one or more customers. If no filter is set, 10 customers will be returned.
func (s *Initialization) Project_get(req ProjectGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "project.get"

	r.FILTER.CUSTOMER_ID = req.CUSTOMER_ID
	r.FILTER.PROJECT_ID = req.PROJECT_ID

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
