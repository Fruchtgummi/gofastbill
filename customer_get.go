package gofastbill

//	"log"
//"reflect"

type CustomerGet_Request struct {
	CUSTOMER_ID     string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CUSTOMER_NUMBER string `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`
	COUNTRY_CODE    string `xml:"COUNTRY_CODE,omitempty" json:"CUSTOMER_CODE,omitempty"`
	CITY            string `xml:"CITY,omitempty" json:"CITY,omitempty"`
	TERM            string `xml:"TERM,omitempty" json:"TERM,omitempty"`
}

//GET Customer
//FILTER => CITY, COUNTRY_CODE, CUSTOMER_ID, CUSTOMER_NUMBER and TERM as atributte
//TERM:
//Search in on this fields: ORGANIZATION, FIRST_NAME, LAST_NAME, ADDRESS, ADDRESS_2, ZIPCODE, EMAIL, TAGS
//Querying the details of one or more customers. If no filter is set, 10 customers will be returned. When LIMIT is set, up to 100 customers are possible.
func (s *Initialization) Customer_get(req CustomerGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "customer.get"

	r.FILTER.CITY = req.CITY
	r.FILTER.COUNTRY_CODE = req.COUNTRY_CODE
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
