package gofastbill

type RecurringGet_Request struct {
	INVOICE_ID string `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"` //Rechnungs-ID
}

//GET Recurring
//FILTER => INVOICE_ID
//Querying the details of one or more invoices. If no filter is set, 10 invoices will be returned. When LIMIT is set, up to 100 invoices are possible
func (s *Initialization) Recurring_get(req RecurringGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "recurring.get"

	r.FILTER.INVOICE_ID = req.INVOICE_ID

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
