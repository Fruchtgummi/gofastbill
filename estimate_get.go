package gofastbill

type EstimateGet_Request struct {
	CUSTOMER_ID         string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	ESTIMATE_ID         string `xml:"ESTIMATE_ID,omitempty" json:"ESTIMATE_ID,omitempty"`                 //	Angebots-ID
	ESTIMATE_NUMBER     string `xml:"ESTIMATE_NUMBER,omitempty" json:"ESTIMATE_NUMBER,omitempty"`         //	Angebotsnummer
	START_ESTIMATE_DATE string `xml:"START_ESTIMATE_DATE,omitempty" json:"START_ESTIMATE_DATE,omitempty"` //	Angebote ab einem bestimmten Datum
	END_ESTIMATE_DATE   string `xml:"END_ESTIMATE_DATE,omitempty" json:"END_ESTIMATE_DATE,omitempty"`     //	Angebote bis zu einem bestimmten Datum
}

//GET Estimate
//FILTER => All fields from Struct: gofastbill.EstimateGet_Request
//RESPONSE:All fields from Struct: gofastbill.ESTIMATE
func (s *Initialization) Estimate_get(req EstimateGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "estimate.get"

	r.DATA.CUSTOMER_ID = req.CUSTOMER_ID
	r.DATA.ESTIMATE_ID = req.ESTIMATE_ID
	r.DATA.ESTIMATE_NUMBER = req.ESTIMATE_NUMBER
	r.DATA.START_ESTIMATE_DATE = req.START_ESTIMATE_DATE
	r.DATA.END_ESTIMATE_DATE = req.END_ESTIMATE_DATE

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
