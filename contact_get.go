package gofastbill

type ContactGet_Request struct {
	CUSTOMER_ID     string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CUSTOMER_NUMBER string `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`
	TERM            string `xml:"TERM,omitempty" json:"TERM,omitempty"`
	CONTACT_ID      string `xml:"CONTACT_ID,omitempty" json:"CONTACT_ID,omitempty"`
}

//GET Customer;
func (s *Initialization) Contact_Get(req ContactGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "contact.get"

	r.FILTER.CONTACT_ID = req.CONTACT_ID
	r.FILTER.CUSTOMER_ID = req.CUSTOMER_ID
	r.FILTER.CUSTOMER_NUMBER = req.CUSTOMER_NUMBER
	r.FILTER.TERM = req.TERM

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
