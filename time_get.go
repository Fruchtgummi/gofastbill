package gofastbill

type TimeGet_Request struct {
	CUSTOMER_ID string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` //
	PROJECT_ID  string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` // ID Eine bestimmte Projekt ID
	TASK_ID     string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` // ID einer bestimmten Aufgabe
	TIME_ID     string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` // ID eines bestimmten Zeiteintrags
	START_DATE  string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` //	Datum des ersten Rechnungslaufs
	END_DATE    string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` //	Enddatum
	DATE        string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"` //	Datum`
}

//GET Time
//FILTER => CUSTOMER_ID, PJOJECT_ID, TASK_ID, TIME_ID, START_DATE, END_DATE, DATE
//Querying the details of one or more customers. If no filter is set, 10 customers will be returned.
func (s *Initialization) Time_get(req TimeGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "time.get"

	r.FILTER.CUSTOMER_ID = req.CUSTOMER_ID
	r.FILTER.PROJECT_ID = req.PROJECT_ID
	r.FILTER.TASK_ID = req.TASK_ID
	r.FILTER.TIME_ID = req.TIME_ID
	r.FILTER.START_DATE = req.START_DATE
	r.FILTER.END_DATE = req.END_DATE
	r.FILTER.DATE = req.DATE

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
